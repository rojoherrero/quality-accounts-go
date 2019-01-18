package main

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	nats "github.com/nats-io/go-nats"
	"github.com/rojoherrero/quality-accounts/server"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

func main() {
	appConfig := readConfigFile()
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	postgres := appConfig.connectToPostgres()
	nc := appConfig.connectToNATS()
	server.New(postgres, nc, logger).Run(appConfig.getServerPort())
}

type config struct {
	*viper.Viper
}

func readConfigFile() *config {
	cfg := viper.New()
	cfg.SetConfigName("application")
	cfg.AddConfigPath(".")
	if e := cfg.ReadInConfig(); e != nil {
		panic(e)
	}
	return &config{cfg}
}

func (c *config) connectToPostgres() *sqlx.DB {
	data := c.Sub("datasources.postgres")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		data.GetString("host"), data.GetString("port"), data.GetString("user"),
		data.GetString("password"), data.GetString("dbname"))
	return sqlx.MustConnect("postgres", dsn)
}

func (c *config) connectToNATS() *nats.Conn {
	data := c.Sub("datasources.nats")
	dsn := fmt.Sprintf("nats://%s:%s", data.GetString("host"), data.GetString("port"))
	nc, e := nats.Connect(dsn)
	if e != nil {
		panic(e)
	}
	return nc
}

func (c *config) getServerPort() string {
	return fmt.Sprintf(":%s", c.GetString("server.port"))
}
