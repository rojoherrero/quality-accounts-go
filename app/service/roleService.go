package service

//go:generate mockgen -source=$GOFILE -destination=../mock/mock_$GOFILE -package=mock

import (
	"github.com/rojoherrero/quality-accounts/app/model/entity"
	"github.com/rojoherrero/quality-accounts/app/model/request"
	"github.com/rojoherrero/quality-accounts/app/repository"
)

type (
	RoleService interface {
		Save(role entity.Role) error
		Update(data request.RoleUpdate) error
		Paginate(start, end int) ([]entity.Role, error)
		Delete(code string) error
	}

	roleService struct {
		service repository.RoleRepository
	}
)

func NewRoleService(service repository.RoleRepository) RoleService {
	return &roleService{service: service}
}

func (s *roleService) Save(role entity.Role) error {
	return s.service.Save(role)
}
func (s *roleService) Update(data request.RoleUpdate) error {
	return s.service.Update(data)
}
func (s *roleService) Paginate(start, end int) ([]entity.Role, error) {
	return s.service.Paginate(start, end)
}
func (s *roleService) Delete(code string) error {
	return s.service.Delete(code)
}
