package repository

//go:generate mockgen -source=$GOFILE -destination=../mock/mock_$GOFILE -package=mock

import (
	"github.com/jackc/pgx"
	"github.com/rojoherrero/quality-accounts/app/model"
	"github.com/rojoherrero/quality-common"
)

type (
	DepartmentRepository interface {
		Save(dep model.RoleDepartment) error
		Update(data model.RoleDepartmentUpdate) error
		Paginate(start, end int) (model.RolesDepartments, error)
		Delete(code string) error
	}

	departmentRepository struct {
		db     *pgx.ConnPool
		logger common.Logger
	}
)

func NewDepartmentRepository(db *pgx.ConnPool, logger common.Logger) DepartmentRepository {
	return &departmentRepository{
		db:     db,
		logger: logger,
	}
}

func (r *departmentRepository) Save(dep model.RoleDepartment) error {
	const newDepartment = "insert into accounts.departments(code, description) value ($1, $2)"
	_, e := r.db.Exec(newDepartment, dep.Code, dep.Description)
	return e
}

func (r *departmentRepository) Update(data model.RoleDepartmentUpdate) error {
	var e error
	if data.NewCode == "" {
		_, e = r.db.Exec("update accounts.departments set description = $1 where code = $2", data.NewDescription, data.OldCode)
	} else {
		_, e = r.db.Exec("update accounts.departments set code = $1, description = $2 where code = $3", data.NewCode, data.NewDescription, data.OldCode)
	}

	return e
}

func (r *departmentRepository) Paginate(start, end int) (model.RolesDepartments, error) {
	var deps model.RolesDepartments
	rows, e := r.db.Query("select r.code, r.description from accounts.departments r order by r.code asc limit $1 offset $2", end-start, start)
	defer rows.Close()
	if e != nil {
		return deps, e
	}
	for rows.Next() {
		var dep model.RoleDepartment
		e := rows.Scan(&dep.Code, &dep.Description)
		if e != nil {
			return deps, e
		}
		dep.Type = model.Department
		deps = append(deps, dep)
	}
	return deps, nil
}

func (r *departmentRepository) Delete(code string) error {
	_, e := r.db.Exec("delete from accounts.departments r where r.code = $1", code)
	return e
}
