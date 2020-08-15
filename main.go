package main

import (
	"fmt"
	"go-jahitin/boot"
	"go-jahitin/helper/config"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	configPath = "config.yaml"
)

func main() {
	cfg := config.NewConfig(configPath)

	host, port := os.Getenv("APP_HOST"), os.Getenv("APP_PORT")
	if host == "" || port == "" {
		host, port = cfg.GetString("app.host"), cfg.GetString("app.port")
	}

	env := os.Getenv("GIN_MODE")
	if env == "" {
		env = cfg.GetString("app.mode")
	}

	if env == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()
	boot.Setup(r, cfg)
	boot.InitializeRoutes(r, cfg)

	r.Run(fmt.Sprintf("%s:%s", host, port))
}
