package server

import (
	"github.com/jmoiron/sqlx"
	"github.com/nats-io/go-nats"
	"github.com/rojoherrero/quality-accounts/server/handler"
	"github.com/rojoherrero/quality-accounts/server/repository"
	"github.com/rojoherrero/quality-accounts/server/service"
	"github.com/rs/zerolog"
)

type api struct {
	role       handler.RoleHandler
	department handler.DepartmentHandler
	user       handler.UserHandler
}

func newApi(db *sqlx.DB, nc *nats.Conn, logger zerolog.Logger) *api {
	return &api{
		role:       initRoleHandler(db, logger),
		department: initDepartmentHandler(db, logger),
		user:       initUserHandler(db, logger),
	}
}

func initRoleHandler(db *sqlx.DB, logger zerolog.Logger) handler.RoleHandler {
	r := repository.NewRoleRepository(db, logger)
	s := service.NewRoleService(r, logger)
	return handler.NewRoleHandler(s, logger)
}

func initDepartmentHandler(db *sqlx.DB, logger zerolog.Logger) handler.DepartmentHandler {
	r := repository.NewDepartmentRepository(db, logger)
	s := service.NewDepartmentService(r, logger)
	return handler.NewDepartmentHandler(s, logger)
}

func initUserHandler(db *sqlx.DB, logger zerolog.Logger) handler.UserHandler {
	r := repository.NewUserRepository(db, logger)
	s := service.NewUserService(r, logger)
	return handler.NewUserHandler(s, logger)
}
