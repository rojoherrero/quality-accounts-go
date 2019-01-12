package service

//go:generate mockgen -source=$GOFILE -destination=../mock/mock_$GOFILE -package=mock

import (
	"context"
	"github.com/rs/zerolog"

	"github.com/rojoherrero/quality-accounts/server/model"
	"github.com/rojoherrero/quality-accounts/server/repository"
)

type (
	RoleService interface {
		Save(ctx context.Context, role []model.Role) error
		Update(ctx context.Context, data model.Role) error
		Paginate(ctx context.Context, start, end int) ([]model.Role, error)
		Delete(ctx context.Context, code string) error
	}

	roleService struct {
		service repository.RoleRepository
		logger  zerolog.Logger
	}
)

func NewRoleService(service repository.RoleRepository, logger zerolog.Logger) RoleService {
	return &roleService{service: service, logger: logger}
}

func (s *roleService) Save(ctx context.Context, role []model.Role) error {
	return s.service.Save(ctx, role)
}

func (s *roleService) Update(ctx context.Context, data model.Role) error {
	return s.service.Update(ctx, data)
}

func (s *roleService) Paginate(ctx context.Context, start, end int) ([]model.Role, error) {
	return s.service.Paginate(ctx, start, end)
}

func (s *roleService) Delete(ctx context.Context, code string) error {
	return s.service.Delete(ctx, code)
}
