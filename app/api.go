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
		role:       initRoleHandler(db),
		department: initDepartmentHandler(db),
	}
}

func initRoleHandler(db *pgx.ConnPool) handler.RoleHandler {
	r := repository.NewRoleRepository(db)
	s := service.NewRoleService(r)
	return handler.NewRoleHandler(s)
}

func initDepartmentHandler(db *pgx.ConnPool) handler.DepartmentHandler {
	r := repository.NewDepartmentRepository(db)
	s := service.NewDepartmentService(r)
	return handler.NewDepartmentHandler(s)
}
