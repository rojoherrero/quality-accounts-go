package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
	"github.com/nats-io/go-nats"
	"github.com/rojoherrero/quality-common"
)

type App struct {
	engine *gin.Engine
	api    *api
}

func InitApp(db *pgx.ConnPool, nc *nats.Conn, logger common.Logger) *App {
	a := &App{
		api:    newApi(db, nc, logger),
		engine: gin.Default(),
	}
	a.createRoutes()
	return a

}

func (a *App) createRoutes() {
	a.rolesRoutes()
}

func (a *App) rolesRoutes() {
	a.engine.POST("/role", a.api.role.Save)
	a.engine.PUT("/role/update", a.api.role.Update)
	a.engine.GET("/role/{start}/{finish}", a.api.role.Paginate)
	a.engine.DELETE("/role", a.api.role.Delete)

}

func (a *App) Run(port string) {
	listeningPort := fmt.Sprintf(":%s", port)
	a.engine.Run(listeningPort)
}
