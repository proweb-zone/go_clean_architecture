package postgresql

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Connection struct {
	Conn *sql.DB
}

type IConnectionDb interface {
	ConnDb() *sql.DB
}

func BuildConnPg() *Connection {
	connStr := "host=localhost port=5433 user=kafka_user password=kafka_user dbname=kafka_db sslmode=disable"

	db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    }

	return &Connection{Conn: db}
}

func (c * Connection) ConnDb() *sql.DB {
return c.Conn
}
