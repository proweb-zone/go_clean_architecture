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
	AddListenerDb(newListener entitie.ListenerEntitie) (bool, error)
}

func InitListenerRepo() *ListenerRepo {
	var connDb postgresql.IConnectionDb = postgresql.BuildConnPg()
	return &ListenerRepo{Conn: connDb}
}

func (l *ListenerRepo) GetListenerList() {
	fmt.Println("получаем список слушателей")

}

func (l *ListenerRepo) AddListenerDb(newListener entitie.ListenerEntitie) (bool, error) {
	fmt.Println("listenerRepo - AddListenerDb")
	var db *sql.DB = l.Conn.ConnDb()

	jsonBytes, jsonErrorObj := json.Marshal(newListener.Settings)

	if jsonErrorObj != nil {
		return false, fmt.Errorf("Error incorrect Setting object")
	}

	response, err := db.Exec("INSERT INTO listeners (name, settings) VALUES ('"+newListener.Name+"', $1)", jsonBytes)

	if err != nil {
		return false, fmt.Errorf(err.Error())
	}

	rowsAffected, errResponse := response.RowsAffected()

	if errResponse != nil {
		return false, fmt.Errorf("При записи в БД произошла ошибка")
	}

	if rowsAffected > 0 {
		return true, nil
	} else {
		return false, fmt.Errorf("Ошибка записи в БД")
	}

}
