package app

import (
	"github.com/jackc/pgx"
	"github.com/nats-io/go-nats"
	"github.com/rojoherrero/quality-accounts/app/handler"
	"github.com/rojoherrero/quality-accounts/app/repository"
	"github.com/rojoherrero/quality-accounts/app/service"
)

type api struct {
	roleHandler handler.RoleHandler
}

func newApi(db *pgx.ConnPool, nc *nats.Conn) *api {
	return &api{roleHandler: initRoleHandler(db)}
}

func initRoleHandler(db *pgx.ConnPool) handler.RoleHandler {
	r := repository.NewRoleRepository(db)
	s := service.NewRoleService(r)
	return handler.NewRoleHandler(s)
}
