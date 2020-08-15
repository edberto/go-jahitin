package setup

import (
	"go-jahitin/apipackages"
	"go-jahitin/helper/auth"
	"go-jahitin/helper/config"
	"go-jahitin/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupAuthMiddleware(cfg config.IConfig, toolkit *apipackages.Toolkit) gin.HandlerFunc {
	accessKey := os.Getenv("KEY_ACCESS")
	if accessKey == "" {
		accessKey = cfg.GetString("key.access")
	}

	j := auth.NewAuth(accessKey)

	return middleware.SetAuthCtx(j, cfg, toolkit)
}
