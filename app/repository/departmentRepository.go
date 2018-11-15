package repository

//go:generate mockgen -source=$GOFILE -destination=../mock/mock_$GOFILE -package=mock

import (
	"github.com/jackc/pgx"
	"github.com/rojoherrero/quality-accounts/app/model/entity"
)

type (
	DepartmentRepository interface {
		Save(dept entity.Department) error
		Update(data entity.Department) error
		Paginate(start, end int) ([]entity.Department, error)
		Delete(code string) error
	}

	departmentRepository struct {
		db *pgx.ConnPool
	}
)

func NewDepartmentRepository(db *pgx.ConnPool) DepartmentRepository {
	return &departmentRepository{db: db}
}

func (r *departmentRepository) Save(dept entity.Department) error {

}

func (r *departmentRepository) Update(data entity.Department) error {

}

func (r *departmentRepository) Paginate(start, end int) ([]entity.Department, error) {

}

func (r *departmentRepository) Delete(code string) error {

}
