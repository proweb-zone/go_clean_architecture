package repository

import (
	"clean/architector/internal/data/postgresql"
	"clean/architector/internal/domain/entitie"
	"database/sql"
	"encoding/json"
	"fmt"
)

type ListenerRepo struct {
	Conn postgresql.IConnectionDb
}

type IlistenerRepo interface {
	GetListenerList()
	AddListenerDb(newListener entitie.ListenerEntitie)
}

func InitListenerRepo() *ListenerRepo {
	var connDb postgresql.IConnectionDb = postgresql.BuildConnPg()
	return &ListenerRepo{Conn: connDb}
}

func (l *ListenerRepo) GetListenerList() {
fmt.Println("получаем список слушателей")

}

func (l *ListenerRepo) AddListenerDb(newListener entitie.ListenerEntitie){
fmt.Println("listenerRepo - AddListenerDb")
var db *sql.DB = l.Conn.ConnDb()

	jsonBytes, jsonErrorObj := json.Marshal(newListener.Settings)

	if jsonErrorObj != nil {
		fmt.Println(jsonErrorObj)
		return
	}

result, err := db.Exec("INSERT INTO listeners (name, settings) VALUES ('"+newListener.Name+"', $1)", jsonBytes);

if err != nil {
	fmt.Println(err)
	return
}

fmt.Println(result)

}
