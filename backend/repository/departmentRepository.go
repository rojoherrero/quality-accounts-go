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
	insertDepartmentQueryBase = "INSERT INTO accounts.departments(code, name) VALUES %s"
	updateDepartmentQuery = "UPDATE accounts.departments SET name = $1 WHERE code = $2"
	paginateDepartmentsQuery = `SELECT d.code as "department_code", d."name" as "department_name"
								FROM accounts.departments d
								ORDER BY d.name ASC 
						  		OFFSET $1 LIMIT $2`
	deleteDepartmentQuery = "DELETE FROM accounts.departments WHERE code = $1"
)

type (
	DepartmentRepository interface {
		Save(ctx context.Context, department []model.Department) error
		Update(ctx context.Context, department model.Department) error
		Paginate(ctx context.Context, start, end int) ([]model.Department, error)
		Delete(ctx context.Context, code string) error
	}

	departmentRepository struct {
		db *sqlx.DB
	}
)

func NewDepartmentRepository(db *sqlx.DB) DepartmentRepository {
	return &departmentRepository{db: db}
}

func (r *departmentRepository) Save(ctx context.Context, departments []model.Department) error {
	numberOfColumns := 2
	rolesLength := len(departments)
	valueStrings := make([]string, 0, rolesLength)
	valueArgs := make([]interface{}, 0, rolesLength * numberOfColumns)
	i := 0
	for _, department := range departments {
		dollarNumberSeed := i * numberOfColumns
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d)", dollarNumberSeed + 1, dollarNumberSeed + 2))
		valueArgs = append(valueArgs, department.Code)
		valueArgs = append(valueArgs, department.Name)
		i++
	}
	rawQuery := fmt.Sprintf(insertDepartmentQueryBase, strings.Join(valueStrings, ","))
	stmt, _ := r.db.Prepare(rawQuery)
	if _, e := stmt.Exec(valueArgs...); e != nil {
		return e
	}
	return nil
}

func (r *departmentRepository) Update(ctx context.Context, department model.Department) error {
	stmt, _ := r.db.Preparex(updateDepartmentQuery)
	if _, e := stmt.ExecContext(ctx, department.Name, department.Code); e != nil {
		return e
	}
	return nil
}

func (r *departmentRepository) Paginate(ctx context.Context, start, end int) ([]model.Department, error) {
	var departments []model.Department
	if e := r.db.SelectContext(ctx, &departments, paginateDepartmentsQuery, start, end-start); e != nil {
		return departments, e
	}
	return departments, nil
}

func (r *departmentRepository) Delete(ctx context.Context, code string) error {
	stmt, _ := r.db.Preparex(deleteDepartmentQuery)
	if _, e := stmt.ExecContext(ctx, code); e != nil {
		return e
	}
	return nil
}
