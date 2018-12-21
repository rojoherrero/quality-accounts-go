package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/rojoherrero/quality-accounts/app/model"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

const (
	insertUserQuery    = "insert into accounts.users(full_name, user_name, password) values ($1, $2, $3);"
	updateUserQuery    = "update accounts.roles set code = $1, name = $2 where code = $3"
	paginateUsersQuery = `SELECT r.code as "role_code", r."name" as "role_name" 
						  FROM accounts.roles r 
						  ORDER BY r."name" ASC 
						  OFFSET $1 LIMIT $2`
	deleteUserQuery = "delete from accounts.roles where code = $1"

	userRolesDepartmentTable = "accounts.user_roles_departments"
	userIdColumn = "user_id"
	roleCodeColumn = "role_code"
	departmentCodeColumn = "department_code"
)

type (
	UserRepository interface {
		Save(ctx context.Context, user model.UserCreationDto) error
		Update(ctx context.Context, user model.User, roleId, departmentId string) error
		Paginate(ctx context.Context, start, end int64) ([]model.User, error)
		Delete(ctx context.Context, code string) error
		GetOne(ctx context.Context, username string) (model.User, error)
		GetLogInData(ctx context.Context, username string) error
	}

	userRepository struct{ db *sqlx.DB }
)

func NewAccountRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Save(ctx context.Context, user model.UserCreationDto) error {
	hash, e := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if e != nil {
		return e
	}
	user.Password = string(hash)
	tx, _ := r.db.Begin()
	userStmt, _ := tx.Prepare(insertUserQuery)
	result, _ := userStmt.ExecContext(ctx, user.UserName, user.Password, user.FullName)
	userId, _ := result.LastInsertId()

	roleDepartmentStmt, _ := tx.Prepare(pq.CopyIn(userRolesDepartmentTable, userIdColumn, roleCodeColumn, departmentCodeColumn))

	for departmentCode, roles := range user.RolesDepartments {
		for _, roleCode := range roles {
			if _, e := roleDepartmentStmt.Exec(userId, roleCode, departmentCode); e != nil{
				return e
			}
		}
	}
	if _, e = roleDepartmentStmt.Exec(); e != nil {
		return e
	}

	return tx.Commit()
}

func (r *userRepository) createRoleDepartmentInsertQuery(sortedData []*userIdRoleCodeDepartmentCode) ([]interface{}, string) {
	numberOfFields := 3
	sortedDataLength := len(sortedData)
	valueStrings := make([]string, 0, sortedDataLength)
	valueArgs := make([]interface{}, 0, sortedDataLength * numberOfFields)
	i := 0

	for _, item := range sortedData {
		dollarNumberSeed := i * numberOfFields
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d)", dollarNumberSeed + 1, dollarNumberSeed + 2, dollarNumberSeed + 3))
		valueArgs = append(valueArgs, item.userID)
		valueArgs = append(valueArgs, item.roleCode)
		valueArgs = append(valueArgs, item.departmentCode)
		i++
	}
	rawQuery := fmt.Sprintf("insert into accounts.user_roles_departments(user_id, role_code,departmente_code) values %s", strings.Join(valueStrings, ","))
	return valueArgs, rawQuery
}

type userIdRoleCodeDepartmentCode struct {
	userID         int64
	roleCode       string
	departmentCode string
}

type sortedData []*userIdRoleCodeDepartmentCode

func (r *userRepository) transformDepartmentsRoles(departmentRoles map[string][]string, userId int64) sortedData {
	var maxLength = 0
	for _, roles := range departmentRoles {
		maxLength += len(roles)
	}
	data := make(sortedData, 0, maxLength)
	for departmentCode, roles := range departmentRoles {
		for _, role := range roles {
			userDepartmentRole := new(userIdRoleCodeDepartmentCode)
			userDepartmentRole.userID = userId
			userDepartmentRole.roleCode = role
			userDepartmentRole.departmentCode = departmentCode
			data = append(data, userDepartmentRole)
		}
	}
	return data
}

func (r *userRepository) Update(ctx context.Context, user model.User, roleId, departmentId string) error {
	return nil
}

func (r *userRepository) Paginate(ctx context.Context, start, end int64) ([]model.User, error) {
	return nil, nil
}

func (r *userRepository) Delete(ctx context.Context, code string) error {
	return nil
}

func (r *userRepository) GetOne(ctx context.Context, username string) (model.User, error) {
	return nil, nil
}

func (r *userRepository) GetLogInData(ctx context.Context, username string) error {
	return nil
}
