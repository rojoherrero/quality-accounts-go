package service

//go:generate mockgen -source=$GOFILE -destination=../mock/mock_$GOFILE -package=mock

import (
	"github.com/rojoherrero/quality-accounts/app/model"
	"github.com/rojoherrero/quality-accounts/app/repository"
	common "github.com/rojoherrero/quality-common"
)

type (
	RoleService interface {
		Save(role model.RoleDepartment) error
		Update(data model.RoleDepartment, code string) error
		Paginate(start, end int) (model.RolesDepartments, error)
		Delete(code string) error
	}

	roleService struct {
		service repository.RoleRepository
		logger  common.Logger
	}
)

func NewRoleService(service repository.RoleRepository, logger common.Logger) RoleService {
	return &roleService{
		service: service,
		logger:  logger,
	}
}

func (s *roleService) Save(role model.RoleDepartment) error {
	return s.service.Save(role)
}

func (s *roleService) Update(data model.RoleDepartment, code string) error {
	return s.service.Update(data, code)
}

func (s *roleService) Paginate(start, end int) (model.RolesDepartments, error) {
	return s.service.Paginate(start, end)
}

func (s *roleService) Delete(code string) error {
	return s.service.Delete(code)
}
