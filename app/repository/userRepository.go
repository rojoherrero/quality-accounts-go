package repository

//go:generate mockgen -source=$GOFILE -destination=../mock/mock_$GOFILE -package=mock

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rojoherrero/quality-accounts/app/model"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

const (
	insertUserQuery    = "insert into accounts.users(full_name, user_name, password) values ($1, $2, $3)"
	updateUserQuery    = `update accounts.users 
							 set full_name = $1,
							     user_name = $2,
							     password = $3,
							     updated = now() 
							 where id = $4`
	paginateUsersQuery = `select to_jsonb(row_to_json(t)) as users
						  from (
   						      select
							      u.id,
     						      u.full_name,
								  u.user_name,
     						      u.password,
     						      array_to_json(array_agg(row_to_json(r))) as user_roles,
     						      array_to_json(array_agg(row_to_json(d))) as user_departments
   						      from accounts.users u
          					  inner join accounts.user_roles_departments urd on u.id = urd.user_id
          					  inner join accounts.roles r on urd.role_code = r.code
          					  inner join accounts.departments d on d.code = urd.departmente_code
							  group by u.id
   							  order by u.user_name asc
            				  OFFSET $1
   							  LIMIT $2
						  ) t`
	deleteUserDepartmentsRolesQuery     = "delete from accounts.user_roles_departments urd where urd.user_id = $1"
	insertUserDepartmentsRolesBaseQuery = "insert into accounts.user_roles_departments(user_id, role_code, departmente_code) values %s"

	deleteUserQuery = "delete from accounts.users where id = $1 or user_name = $2"
)

type (
	UserRepository interface {
		Save(ctx context.Context, user model.UserCreationDto) error
		Update(ctx context.Context, user model.UserCreationDto) error
		Paginate(ctx context.Context, start, end int64) (model.PropertyMapSlice, error)
		Delete(ctx context.Context, code string) error
		GetLogInData(ctx context.Context, username string) error
	}

	userRepository struct{ db *sqlx.DB }
)

func NewAccountRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Save(ctx context.Context, user model.UserCreationDto) error {
	hash, e := r.hashPassword(user.Password)
	if e != nil {
		return e
	}
	user.Password = string(hash)
	tx, _ := r.db.Begin()
	userStmt, _ := tx.Prepare(insertUserQuery)
	result, _ := userStmt.ExecContext(ctx, user.UserName, user.Password, user.FullName)
	userId, _ := result.LastInsertId()

	departmentsRolesUpdateQuery, departmentsRolesUpdateArgs := createUserDepartmentsRolesInsertQuery(user.DepartmentRoles, userId)
	departmentsRolesUpdateStmt, _ := tx.Prepare(departmentsRolesUpdateQuery)
	if _, e = departmentsRolesUpdateStmt.ExecContext(ctx, departmentsRolesUpdateArgs); e != nil {
		return e
	}

	return tx.Commit()
}

func (r *userRepository) Update(ctx context.Context, user model.UserCreationDto) error {
	hash, e := r.hashPassword(user.Password)
	if e != nil {
		return e
	}
	user.Password = string(hash)
	tx, _ := r.db.Begin()
	userUpdateStmt, _ := tx.Prepare(updateUserQuery)
	if _, e = userUpdateStmt.ExecContext(ctx, user.ID, user.UserName, user.Password, user.FullName); e != nil {
		return e
	}

	if _, e = tx.ExecContext(ctx, deleteUserDepartmentsRolesQuery, user.ID); e != nil {
		return e
	}
	departmentsRolesUpdateQuery, departmentsRolesUpdateArgs := createUserDepartmentsRolesInsertQuery(user.DepartmentRoles, user.ID)
	departmentsRolesUpdateStmt, _ := tx.Prepare(departmentsRolesUpdateQuery)
	if _, e = departmentsRolesUpdateStmt.ExecContext(ctx, departmentsRolesUpdateArgs); e != nil {
		return e
	}

	return nil
}

func createUserDepartmentsRolesInsertQuery(rolesDepartments map[string][]string, userId int64) (string, []interface{}) {
	numberOfColumns := 3
	valueStrings := make([]string, 0)
	valueArgs := make([]interface{}, 0)
	i := 0
	for departmentCode, roles := range rolesDepartments {
		for _, roleCode := range roles {
			dollarNumberSeed := i * numberOfColumns
			valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d)", dollarNumberSeed+1, dollarNumberSeed+2, dollarNumberSeed+3))
			valueArgs = append(valueArgs, userId)
			valueArgs = append(valueArgs, roleCode)
			valueArgs = append(valueArgs, departmentCode)
			i++
		}
	}
	return fmt.Sprintf(insertUserDepartmentsRolesBaseQuery, strings.Join(valueStrings, ",")), valueArgs
}


func (r *userRepository) hashPassword(rawPassword string) (string, error) {
	hash, e := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	return string(hash), e
}

func (r *userRepository) Paginate(ctx context.Context, start, end int64) (model.PropertyMapSlice, error) {
	stmt, _ := r.db.PrepareContext(ctx, paginateUsersQuery)
	rows, e := stmt.Query(start, end-start)
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	var users model.PropertyMapSlice
	for rows.Next() {
		var user model.PropertyMap
		if e := rows.Scan(&user); e != nil {
			return users, e
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *userRepository) Delete(ctx context.Context, code string) error {
	if _, e := r.db.ExecContext(ctx, deleteUserQuery, code); e != nil {
		return e
	}
	return nil
}

func (r *userRepository) GetLogInData(ctx context.Context, username string) error {
	return nil
}
