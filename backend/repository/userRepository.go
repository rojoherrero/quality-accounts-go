package repository

//go:generate mockgen -source=$GOFILE -destination=../mock/mock_$GOFILE -package=mock

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rojoherrero/quality-accounts/backend/model"
	"strings"
)

const (
	insertUserQuery    = "insert into accounts.users(full_name, user_name, password) values ($1, $2, $3) returning id"
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
		Delete(ctx context.Context, id int64) error
		GetLogInData(ctx context.Context, username string) error
	}

	userRepository struct{ db *sqlx.DB }
)

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Save(ctx context.Context, user model.UserCreationDto) error {
	tx, _ := r.db.Beginx()
	var userId int64
	if e := tx.Get(&userId, insertUserQuery, user.FullName, user.UserName, user.Password); e != nil {
		tx.Rollback()
		return e
	}
	departmentsRolesUpdateQuery, departmentsRolesUpdateArgs := createUserDepartmentsRolesInsertQuery(user.DepartmentRoles, userId)
	if _, e := tx.ExecContext(ctx, departmentsRolesUpdateQuery, departmentsRolesUpdateArgs...); e != nil {
		tx.Rollback()
		return e
	}

	return tx.Commit()
}

func (r *userRepository) Update(ctx context.Context, user model.UserCreationDto) error {
	tx, _ := r.db.Begin()
	if _, e := tx.ExecContext(ctx, updateUserQuery, user.FullName, user.UserName, user.Password, user.ID); e != nil {
		tx.Rollback()
		return e
	}
	if len(user.DepartmentRoles.DepartmentCode) != 0{
		if _, e := tx.ExecContext(ctx, deleteUserDepartmentsRolesQuery, user.ID); e != nil {
			tx.Rollback()
			return e
		}
		departmentsRolesUpdateQuery, departmentsRolesUpdateArgs := createUserDepartmentsRolesInsertQuery(user.DepartmentRoles, user.ID)
		if _, e := tx.ExecContext(ctx, departmentsRolesUpdateQuery, departmentsRolesUpdateArgs...); e != nil {
			return e
		}
	}

	tx.Commit()
	return nil
}

func createUserDepartmentsRolesInsertQuery(departmentRoles model.DepartmentRoles, userId int64) (string, []interface{}) {
	numberOfColumns := 3
	valueStrings := make([]string, 0)
	valueArgs := make([]interface{}, 0)
	for i, roleCode := range departmentRoles.RolesCodes {
		dollarNumberSeed := i * numberOfColumns
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d)", dollarNumberSeed+1, dollarNumberSeed+2, dollarNumberSeed+3))
		valueArgs = append(valueArgs, userId)
		valueArgs = append(valueArgs, roleCode)
		valueArgs = append(valueArgs, departmentRoles.DepartmentCode)
	}
	return fmt.Sprintf(insertUserDepartmentsRolesBaseQuery, strings.Join(valueStrings, ",")), valueArgs
}

func (r *userRepository) Paginate(ctx context.Context, start, end int64) (model.PropertyMapSlice, error) {
	var users model.PropertyMapSlice
	if e := r.db.SelectContext(ctx, &users, paginateUsersQuery, start, end-start); e != nil {
		return users, e
	}
	return users, nil
}

func (r *userRepository) Delete(ctx context.Context, id int64) error {
	if _, e := r.db.ExecContext(ctx, deleteUserQuery, id); e != nil {
		return e
	}
	return nil
}

func (r *userRepository) GetLogInData(ctx context.Context, username string) error {
	return nil
}
