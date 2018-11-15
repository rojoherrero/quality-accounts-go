package main

import (
	. "github.com/rojoherrero/quality-accounts/app"
	"github.com/rojoherrero/quality-common"
)

const appPrefix = "accounts"

func main() {
	configService := common.InitConfigService(appPrefix)
	db, nc := configService.GetDataSources()

	app := InitApp(db, nc)

	app.Run("8080")

}
