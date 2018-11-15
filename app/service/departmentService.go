package service

import (
	"github.com/rojoherrero/quality-accounts/app/model/entity"
	"github.com/rojoherrero/quality-accounts/app/model/request"
	"github.com/rojoherrero/quality-accounts/app/repository"
)

//go:generate mockgen -source=$GOFILE -destination=../mock/mock_$GOFILE -package=mock

type (
	DepartmentService interface {
		Save(role entity.Department) error
		Update(data request.DepartmentUpdate) error
		Paginate(start, end int) ([]entity.Department, error)
		Delete(code string) error
	}

	departmentService struct {
		repo repository.DepartmentRepository
	}
)

func NewDepartmentService(repo repository.DepartmentRepository) DepartmentService {
	return &departmentService{repo: repo}
}

func (s *departmentService) Save(role entity.Department) error {}
func (s *departmentService) Update(data request.DepartmentUpdate) error {}
func (s *departmentService) Paginate(start, end int) ([]entity.Department, error) {}
func (s *departmentService) Delete(code string) error {}
