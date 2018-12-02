package service

import (
	"github.com/rojoherrero/quality-accounts/app/model"
	"github.com/rojoherrero/quality-accounts/app/repository"
	"github.com/rojoherrero/quality-common"
)

//go:generate mockgen -source=$GOFILE -destination=../mock/mock_$GOFILE -package=mock

type (
	DepartmentService interface {
		Save(dept model.RoleDepartment) error
		Update(data model.RoleDepartmentUpdate) error
		Paginate(start, end int) (model.RolesDepartments, error)
		Delete(code string) error
	}

	departmentService struct {
		repo   repository.DepartmentRepository
		logger common.Logger
	}
)

func NewDepartmentService(repo repository.DepartmentRepository, logger common.Logger) DepartmentService {
	return &departmentService{
		repo:   repo,
		logger: logger,
	}
}

func (s *departmentService) Save(dept model.RoleDepartment) error {
	return s.repo.Save(dept)
}

func (s *departmentService) Update(data model.RoleDepartmentUpdate) error {
	return s.repo.Update(data)
}

func (s *departmentService) Paginate(start, end int) (model.RolesDepartments, error) {
	return s.repo.Paginate(start, end)
}

func (s *departmentService) Delete(code string) error {
	return s.repo.Delete(code)
}
