package main

import (
	"fmt"
	"github.com/system-server2025/global/config"
	"github.com/system-server2025/service"
)

func main() {
	app := service.InitApp()
	cfg := config.GetConfig()
	app.Echo.Logger.Fatal(app.Echo.Start(fmt.Sprintf(":%d", cfg.Server.Port)))
}
