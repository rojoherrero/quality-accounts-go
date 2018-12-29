package service

//go:generate mockgen -source=$GOFILE -destination=../mock/mock_$GOFILE -package=mock

import (
	"context"

	"github.com/rojoherrero/quality-accounts/backend/model"
	"github.com/rojoherrero/quality-accounts/backend/repository"
)

type (
	RoleService interface {
		Save(ctx context.Context, role []model.Role) error
		Update(ctx context.Context, data model.Role, oldCode string) error
		Paginate(ctx context.Context, start, end int) ([]model.Role, error)
		Delete(ctx context.Context, code string) error
	}

	roleService struct {
		service repository.RoleRepository
	}
)

func NewRoleService(service repository.RoleRepository) RoleService {
	return &roleService{service: service}
}

func (s *roleService) Save(ctx context.Context, role []model.Role) error {
	return s.service.Save(ctx, role)
}

func (s *roleService) Update(ctx context.Context, data model.Role, oldCode string) error {
	return s.service.Update(ctx, data, oldCode)
}

func (s *roleService) Paginate(ctx context.Context, start, end int) ([]model.Role, error) {
	return s.service.Paginate(ctx, start, end)
}

func (s *roleService) Delete(ctx context.Context, code string) error {
	return s.service.Delete(ctx, code)
}
