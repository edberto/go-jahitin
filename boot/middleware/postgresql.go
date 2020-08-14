package middleware

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	postgresqlKey = "postgresql"
)

func ConnectPostgresql(uri string) (*sql.DB, error) {
	return sql.Open("postgres", uri)
}

func SetPostgresCtx(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(postgresqlKey, db)
		c.Next()
	}
}

func GetPostgres(c *gin.Context) *sql.DB {
	return c.MustGet(postgresqlKey).(*sql.DB)
}
