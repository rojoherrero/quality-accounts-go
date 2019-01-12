package service

import (
	"context"
	"github.com/rojoherrero/quality-accounts/server/model"
	"github.com/rojoherrero/quality-accounts/server/repository"
	"github.com/rs/zerolog"
)

//go:generate mockgen -source=$GOFILE -destination=../mock/mock_$GOFILE -package=mock

type (
	DepartmentService interface {
		Save(ctx context.Context, department []model.Department) error
		Update(ctx context.Context, department model.Department) error
		Paginate(ctx context.Context, start, end int) ([]model.Department, error)
		Delete(ctx context.Context, code string) error
	}

	departmentService struct {
		repo   repository.DepartmentRepository
		logger zerolog.Logger
	}
)

func NewDepartmentService(repo repository.DepartmentRepository, logger zerolog.Logger) DepartmentService {
	return &departmentService{repo: repo, logger: logger}
}

func (s *departmentService) Save(ctx context.Context, departments []model.Department) error {
	return s.repo.Save(ctx, departments)
}

func (s *departmentService) Update(ctx context.Context, department model.Department) error {
	return s.repo.Update(ctx, department)
}

func (s *departmentService) Paginate(ctx context.Context, start, end int) ([]model.Department, error) {
	return s.repo.Paginate(ctx, start, end)
}

func (s *departmentService) Delete(ctx context.Context, code string) error {
	return s.repo.Delete(ctx, code)
}
