package repository

import (
	"clean/architector/internal/data/postgresql"
	"clean/architector/internal/domain/entitie"
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

type ListenerRepo struct {
	Conn postgresql.IConnectionDb
}

type IlistenerRepo interface {
	GetListenerList() []*entitie.ListenerEntitie
	AddListenerDb(newListener entitie.ListenerEntitie) (bool, error)
}

func InitListenerRepo() *ListenerRepo {
	var connDb postgresql.IConnectionDb = postgresql.BuildConnPg()
	return &ListenerRepo{Conn: connDb}
}

func (l *ListenerRepo) GetListenerList() []*entitie.ListenerEntitie {
	fmt.Println("получаем список слушателей")
	var db *sql.DB = l.Conn.ConnDb()
	listenerListDb, err := db.Query("SELECT * FROM listeners WHERE status=1")

	defer listenerListDb.Close()

	listenerList := make([]*entitie.ListenerEntitie, 0)
	for listenerListDb.Next() {
		listenerItem := new(entitie.ListenerEntitie)

		err := listenerListDb.Scan(&listenerItem.Id, &listenerItem.Name, &listenerItem.Host, &listenerItem.Port, &listenerItem.Deelay, &listenerItem.Status)
		if err != nil {
			log.Fatal(err)
		}

		listenerList = append(listenerList, listenerItem)
	}
	if err = listenerListDb.Err(); err != nil {
		log.Fatal(err)
	}

	return listenerList
}

func (l *ListenerRepo) AddListenerDb(newListener entitie.ListenerEntitie) (bool, error) {
	fmt.Println("listenerRepo - AddListenerDb")
	var db *sql.DB = l.Conn.ConnDb()

	var buildFieldQuery string = "name, host, port"
	var buildValueQuery string = "'" + newListener.Name + "', '" + newListener.Host + "', '" + newListener.Port + "'"

	if newListener.Deelay > 0 {
		buildFieldQuery += ", deelay"
		buildValueQuery += ", " + strconv.Itoa(newListener.Deelay) + ""
	}

	if newListener.Status > 0 {
		buildFieldQuery += ", status"
		buildValueQuery += ", " + strconv.Itoa(newListener.Status) + ""
	}

	response, err := db.Exec("INSERT INTO listeners (" + buildFieldQuery + ") VALUES (" + buildValueQuery + ");")

	if err != nil {
		fmt.Print(err.Error())
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
