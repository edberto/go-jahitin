package setup

import (
	"go-jahitin/boot/middleware"
	"go-jahitin/helper/auth"
	"go-jahitin/helper/config"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupAuthMiddleware(cfg config.IConfig) gin.HandlerFunc {
	accessKey := os.Getenv("KEY_ACCESS")
	if accessKey == "" {
		accessKey = cfg.GetString("key.access")
	}

	j := auth.NewAuth(accessKey)

	return middleware.SetAuthCtx(j, cfg)
}
