package setup

import (
	"go-jahitin/boot/middleware"
	"go-jahitin/helper/config"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func SetupPostgreSQL(cfg config.IConfig) gin.HandlerFunc {
	postgresqlURI := os.Getenv("POSTGRESQL_URI")
	if postgresqlURI == "" {
		postgresqlURI = cfg.GetString("database.postgresql.uri")
	}

	db, err := middleware.ConnectPostgresql(postgresqlURI)
	if err != nil {
		log.Print(errors.Wrap(err, "Failed to connect to PostgreSQL database"))
	}

	if err := db.Ping(); err != nil {
		log.Print(errors.Wrap(err, "Failed to ping PostgreSQL database"))
	} else {
		log.Print("Postgresql database successfully connected")
	}

	handlerFunc := middleware.SetPostgresCtx(db)
	return handlerFunc
}
