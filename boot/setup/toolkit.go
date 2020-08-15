package setup

import (
	"database/sql"
	"go-jahitin/apipackages"
	"go-jahitin/helper/config"
	"go-jahitin/helper/constants"
	"log"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

var (
	roleCheckerTag  = "rolecheck"
	roleCheckerFunc = func(fl validator.FieldLevel) bool {
		if _, e := constants.RoleAtoI[strings.ToLower(fl.Field().String())]; !e {
			return false
		}
		return true
	}
)

func SetupToolkit(cfg config.IConfig) *apipackages.Toolkit {
	tk := new(apipackages.Toolkit)

	db := SetupPostgresqlDB(cfg)
	tk.DB = db

	val := validator.New()
	if err := val.RegisterValidation(roleCheckerTag, roleCheckerFunc); err != nil {
		log.Printf("Failed to initiate role checker validator: %v", err)
	}
	tk.Validator = val

	return tk
}

func SetupPostgresqlDB(cfg config.IConfig) *sql.DB {
	postgresqlURI := os.Getenv("POSTGRESQL_URI")
	if postgresqlURI == "" {
		postgresqlURI = cfg.GetString("database.postgresql.uri")
	}

	db, err := sql.Open("postgres", postgresqlURI)
	if err != nil {
		log.Print(errors.Wrap(err, "Failed to connect to PostgreSQL database"))
	}

	if err := db.Ping(); err != nil {
		log.Print(errors.Wrap(err, "Failed to ping PostgreSQL database"))
	} else {
		log.Print("Postgresql database successfully connected")
	}

	return db
}
