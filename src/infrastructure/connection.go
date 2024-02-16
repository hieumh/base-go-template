package infrastructure

import (
	"database/sql"
	"os"
)

func newConnection() (*sql.DB, error) {
	return sql.Open("sql", os.Getenv("SQL_CONNECTION"))
}
