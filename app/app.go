package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
	"github.com/nats-io/go-nats"
	"net/http"
)

type App struct {
	router *mux.Router
	api    *api
}

func InitApp(db *pgx.ConnPool, nc *nats.Conn) *App {
	a := &App{api: newApi(db, nc), router: mux.NewRouter()}
	a.createRoutes()
	return a

}

func (a *App) createRoutes() {
	a.rolesRoutes()
}

func (a *App) rolesRoutes() {
	a.router.HandleFunc("/role", a.api.roleHandler.Save).Methods(http.MethodPost)
	a.router.HandleFunc("/role/update", a.api.roleHandler.Update).Methods(http.MethodPut)
	a.router.HandleFunc("/role/{start}/{finish}", a.api.roleHandler.Paginate).Methods(http.MethodGet)
	a.router.HandleFunc("/role", a.api.roleHandler.Delete).Methods(http.MethodDelete)

}

func (a *App) Run(port string) {
	listeningPort := fmt.Sprintf(":%s", port)
	http.ListenAndServe(listeningPort, a.router)
}
