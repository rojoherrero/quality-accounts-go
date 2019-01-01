package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/nats-io/go-nats"
	"github.com/rojoherrero/quality-accounts/backend"
)

const appPrefix = "accounts"

func main() {
	//configService, e := cfg.InitConfigService(appPrefix)
	//if e != nil {
	//	panic(e)
	//}
	//dsn, e := configService.GetPostgresDSN()

	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=quality-go sslmode=disable"

	db := sqlx.MustConnect("postgres", dsn)
	db.Ping()

	nc, _ := nats.Connect(nats.DefaultURL)

	app := server.InitServer(db, nc)

	app.Run("8080")
}
