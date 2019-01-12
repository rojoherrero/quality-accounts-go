package repository

//go:generate mockgen -source=$GOFILE -destination=../mock/mock_$GOFILE -package=mock

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"strings"

	"github.com/rojoherrero/quality-accounts/server/model"
)

const (
	insertRoleQueryBase = "insert into accounts.roles(code, name) values %s"
	updateRoleQuery     = "update accounts.roles set name = $1 where code = $2"
	paginateRolesQuery  = `SELECT r.code as "role_code", r."name" as "role_name" 
						  FROM accounts.roles r 
						  ORDER BY r."name" ASC 
						  OFFSET $1 LIMIT $2`
	deleteRoleQuery = "delete from accounts.roles where code = $1"
)

type (
	RoleRepository interface {
		Save(ctx context.Context, roles []model.Role) error
		Update(ctx context.Context, data model.Role) error
		Paginate(ctx context.Context, start, end int) ([]model.Role, error)
		Delete(ctx context.Context, code string) error
	}

	roleRepository struct {
		db     *sqlx.DB
		logger zerolog.Logger
	}
)

func NewRoleRepository(db *sqlx.DB, logger zerolog.Logger) RoleRepository {
	return &roleRepository{db: db, logger: logger}
}

func (r *roleRepository) Save(ctx context.Context, roles []model.Role) error {
	numberOfColumns := 2
	rolesLength := len(roles)
	valueStrings := make([]string, 0, rolesLength)
	valueArgs := make([]interface{}, 0, rolesLength*numberOfColumns)
	for i, role := range roles {
		dollarNumberSeed := i * numberOfColumns
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d)", dollarNumberSeed+1, dollarNumberSeed+2))
		valueArgs = append(valueArgs, role.Code)
		valueArgs = append(valueArgs, role.Name)
	}
	rawQuery := fmt.Sprintf(insertRoleQueryBase, strings.Join(valueStrings, ","))
	stmt, _ := r.db.Prepare(rawQuery)
	if _, e := stmt.ExecContext(ctx, valueArgs...); e != nil {
		return e
	}
	return nil
}

func (r *roleRepository) Update(ctx context.Context, data model.Role) error {
	stmt, _ := r.db.Preparex(updateRoleQuery)
	if _, e := stmt.ExecContext(ctx, data.Name, data.Code); e != nil {
		return e
	}
	return nil
}

func (r *roleRepository) Paginate(ctx context.Context, start, end int) ([]model.Role, error) {
	var roles []model.Role
	if e := r.db.SelectContext(ctx, &roles, paginateRolesQuery, start, end-start); e != nil {
		return roles, e
	}
	return roles, nil
}

func (r *roleRepository) Delete(ctx context.Context, code string) error {
	stmt, _ := r.db.Preparex(deleteRoleQuery)
	if _, e := stmt.ExecContext(ctx, code); e != nil {
		return e
	}
	return nil
}
