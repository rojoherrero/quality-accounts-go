package repository

//go:generate mockgen -source=$GOFILE -destination=../mock/mock_$GOFILE -package=mock

import (
	"github.com/jackc/pgx"
	"github.com/rojoherrero/quality-accounts/app/model/entity"
	"github.com/rojoherrero/quality-accounts/app/model/request"
)

type (
	RoleRepository interface {
		Save(role entity.Role) error
		Update(data request.RoleUpdate) error
		Paginate(start, end int) ([]entity.Role, error)
		Delete(code string) error
	}

	roleRepository struct {
		db *pgx.ConnPool
	}
)

func NewRoleRepository(db *pgx.ConnPool) RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) Save(role entity.Role) error {
	_, e := r.db.Exec("insert into accounts.roles(code, description) value ($1, $2)", role.Code, role.Description)
	return e
}

func (r *roleRepository) Update(data request.RoleUpdate) error {
	var e error
	if data.NewCode == "" {
		_, e = r.db.Exec("update accounts.roles set description = $1 where code = $2", data.NewDescription, data.OldCode)
	} else {
		_, e = r.db.Exec("update accounts.roles set code = $1, description = $2 where code = $3", data.NewCode, data.NewDescription, data.OldCode)
	}

	return e
}

func (r *roleRepository) Paginate(start, end int) ([]entity.Role, error) {
	var roles []entity.Role
	rows, e := r.db.Query("select r.code, r.description from accounts.roles r order by r.code asc limit $1 offset $2", end-start, start)
	defer rows.Close()
	if e != nil {
		return roles, e
	}
	for rows.Next() {
		var role entity.Role
		e := rows.Scan(&role.Code, &role.Description)
		if e != nil {
			return roles, e
		}
		roles = append(roles, role)
	}
	return roles, nil
}

func (r *roleRepository) Delete(code string) error {
	_, e := r.db.Exec("delete from accounts.roles r where r.code = $1", code)
	return e
}
