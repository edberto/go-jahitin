package boot

import (
	"go-jahitin/boot/setup"
	"go-jahitin/helper/config"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine, cfg config.IConfig) {
	dbHF := setup.SetupPostgreSQL(cfg)
	r.Use(dbHF)
}