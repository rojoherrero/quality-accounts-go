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
	updateDepartmentQuery = "UPDATE accounts.roles SET code = $1, name = $2 WHERE code = $3"
	paginateDepartmentsQuery = `SELECT r.code as "role_code", r."name" as "role_name" 
								FROM accounts.roles r 
						  		ORDER BY r."name" ASC 
						  		OFFSET $1 LIMIT $2`
	deleteDepartmentQuery = "DELETE FROM accounts.departments WHERE code = $1"
)

type (
	DepartmentRepository interface {
		Save(ctx context.Context, department []model.Department) error
		Update(ctx context.Context, department model.Department, oldCode string) error
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
	numberOfColumns := 3
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
	if _, e := stmt.Exec(valueArgs); e != nil {
		return e
	}
	return nil
}

func (r *departmentRepository) Update(ctx context.Context, department model.Department, oldCode string) error {
	stmt, _ := r.db.Preparex(updateDepartmentQuery)
	if _, e := stmt.ExecContext(ctx, department.Code, department.Name, oldCode); e != nil {
		return e
	}
	return nil
}

func (r *departmentRepository) Paginate(ctx context.Context, start, end int) ([]model.Department, error) {
	stmt, _ := r.db.Preparex(paginateDepartmentsQuery)
	var departments []model.Department
	rows, e := stmt.Queryx(start, end-start)
	if e != nil {
		return departments, e
	}
	defer rows.Close()
	for rows.Next() {
		var department model.Department
		if e := rows.StructScan(&department); e != nil {
			return departments, e
		}
		departments = append(departments, department)
	}
	return departments, e
}

func (r *departmentRepository) Delete(ctx context.Context, code string) error {
	stmt, _ := r.db.Preparex(deleteDepartmentQuery)
	if _, e := stmt.ExecContext(ctx, code); e != nil {
		return e
	}
	return nil
}
