package repository

import (
	"github.com/jackc/pgx"
	"github.com/rojoherrero/quality-accounts/app/model"
	"github.com/rojoherrero/quality-common"
)

type (
	AccountRepository interface {
		Save(dep model.Account) error
		Update(data model.RoleDepartmentUpdate) error
		Paginate(start, end int) (model.Accounts, error)
		Delete(code string) error
	}

	accountRepository struct {
		db     *pgx.ConnPool
		logger common.Logger
	}
)

func NewAccountRepository(db *pgx.ConnPool, logger common.Logger) AccountRepository {
	return &accountRepository{
		db:     db,
		logger: logger,
	}
}

func (r *accountRepository) Save(dep model.Account) error {

}

func (r *accountRepository) Update(data model.RoleDepartmentUpdate) error {

}

func (r *accountRepository) Paginate(start, end int) (model.Accounts, error) {

}

func (r *accountRepository) Delete(code string) error {

}
