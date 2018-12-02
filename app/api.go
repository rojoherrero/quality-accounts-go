package app

import (
	"github.com/jackc/pgx"
	"github.com/nats-io/go-nats"
	"github.com/rojoherrero/quality-accounts/app/handler"
	"github.com/rojoherrero/quality-accounts/app/repository"
	"github.com/rojoherrero/quality-accounts/app/service"
	"github.com/rojoherrero/quality-common"
)

type api struct {
	role       handler.RoleHandler
	department handler.DepartmentHandler
}

func newApi(db *pgx.ConnPool, nc *nats.Conn, logger common.Logger) *api {
	return &api{
		role:       initRoleHandler(db, logger),
		department: initDepartmentHandler(db, logger),
	}
}

func initRoleHandler(db *pgx.ConnPool, logger common.Logger) handler.RoleHandler {
	r := repository.NewRoleRepository(db, logger)
	s := service.NewRoleService(r, logger)
	return handler.NewRoleHandler(s, logger)
}

func initDepartmentHandler(db *pgx.ConnPool, logger common.Logger) handler.DepartmentHandler {
	r := repository.NewDepartmentRepository(db, logger)
	s := service.NewDepartmentService(r, logger)
	return handler.NewDepartmentHandler(s, logger)
}
