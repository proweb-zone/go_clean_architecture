package postgresql

import (
	"clean/architector/internal/app"
	"database/sql"
)

func ConnPg() *sql.DB {
	cfg := app.MustLoad()
	connStr := "user="+cfg.Postgresql.POSTGRES_USER+" password="+cfg.Postgresql.POSTGRES_PASSWORD+" dbname="+cfg.Postgresql.POSTGRES_DB+" sslmode=disable"

	db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    }

	return db
}
