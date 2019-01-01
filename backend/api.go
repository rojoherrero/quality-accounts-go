package server

import (
	"github.com/jmoiron/sqlx"
	"github.com/nats-io/go-nats"
	"github.com/rojoherrero/quality-accounts/backend/handler"
	"github.com/rojoherrero/quality-accounts/backend/repository"
	"github.com/rojoherrero/quality-accounts/backend/service"
)

type api struct {
	role       handler.RoleHandler
	department handler.DepartmentHandler
	user       handler.UserHandler
}

func newApi(db *sqlx.DB, nc *nats.Conn) *api {
	return &api{
		role:       initRoleHandler(db),
		department: initDepartmentHandler(db),
		user:       initUserHandler(db),
	}
}

func initRoleHandler(db *sqlx.DB) handler.RoleHandler {
	r := repository.NewRoleRepository(db)
	s := service.NewRoleService(r)
	return handler.NewRoleHandler(s)
}

func initDepartmentHandler(db *sqlx.DB) handler.DepartmentHandler {
	r := repository.NewDepartmentRepository(db)
	s := service.NewDepartmentService(r)
	return handler.NewDepartmentHandler(s)
}

func initUserHandler(db *sqlx.DB) handler.UserHandler {
	r := repository.NewUserRepository(db)
	s := service.NewUserService(r)
	return handler.NewUserHandler(s)
}
