package repository

import (
	"clean/architector/internal/data/postgresql"
	"fmt"
)

type ListenerRepo struct {
	Conn postgresql.IConnectionDb
}

type IlistenerRepo interface {
	GetListenerList()
}

func InitListenerRepo() *ListenerRepo {
	var connDb postgresql.IConnectionDb = postgresql.BuildConnPg()
	return &ListenerRepo{Conn: connDb}
}

func (l *ListenerRepo) GetListenerList() {
fmt.Println("получаем список слушателей")

}

func ChangeStatusListener(status bool){
// TODO
}
