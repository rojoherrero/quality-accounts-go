package repository

//go:generate mockgen -source=$GOFILE -destination=../mock/mock_$GOFILE -package=mock

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"

	"github.com/rojoherrero/quality-accounts/backend/model"
)

const (
	insertRoleQueryBase = "insert into accounts.roles(code, name) values %s"
	updateRoleQuery = "update accounts.roles set code = $1, name = $2 where code = $3"
	paginateRolesQuery = `SELECT r.code as "role_code", r."name" as "role_name" 
						  FROM accounts.roles r 
						  ORDER BY r."name" ASC 
						  OFFSET $1 LIMIT $2`
	deleteRoleQuery = "delete from accounts.roles where code = $1"
)

type (
	RoleRepository interface {
		Save(ctx context.Context, roles []model.Role) error
		Update(ctx context.Context, data model.Role, oldCode string) error
		Paginate(ctx context.Context, start, end int) ([]model.Role, error)
		Delete(ctx context.Context, code string) error
	}

	roleRepository struct {
		db *sqlx.DB
	}
)

func NewRoleRepository(db *sqlx.DB) RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) Save(ctx context.Context, roles []model.Role) error {
	numberOfColumns := 3
	rolesLength := len(roles)
	valueStrings := make([]string, 0, rolesLength)
	valueArgs := make([]interface{}, 0, rolesLength * numberOfColumns)
	i := 0
	for _, role := range roles {
		dollarNumberSeed := i * numberOfColumns
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d)", dollarNumberSeed + 1, dollarNumberSeed + 2))
		valueArgs = append(valueArgs, role.Code)
		valueArgs = append(valueArgs, role.Name)
		i++
	}
	rawQuery := fmt.Sprintf(insertRoleQueryBase , strings.Join(valueStrings, ","))
	stmt, _ := r.db.Prepare(rawQuery)
	if _, e := stmt.Exec(valueArgs); e != nil {
		return e
	}
	return nil
}

func (r *roleRepository) Update(ctx context.Context, data model.Role, oldCode string) error {
	stmt, _ := r.db.Preparex(updateRoleQuery)
	if _, e := stmt.ExecContext(ctx, data.Code, data.Name, oldCode); e != nil {
		return e
	}
	return nil
}

func (r *roleRepository) Paginate(ctx context.Context, start, end int) ([]model.Role, error) {
	stmt, _ := r.db.Preparex(paginateRolesQuery)
	var roles []model.Role
	rows, e := stmt.Queryx(start, end-start)
	if e != nil {
		return roles, e
	}
	defer rows.Close()
	for rows.Next() {
		var role model.Role
		if e := rows.StructScan(&role); e != nil {
			return roles, e
		}
		roles = append(roles, role)
	}
	return roles, e
}

func (r *roleRepository) Delete(ctx context.Context, code string) error {
	stmt, _ := r.db.Preparex(deleteRoleQuery)
	if _, e := stmt.ExecContext(ctx, code); e != nil {
		return e
	}
	return nil
}
