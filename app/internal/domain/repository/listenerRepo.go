package repository

import (
	"database/sql"
	"fmt"
)

type ListenerRepo struct {
	Conn *sql.DB
}

type IlistenerRepo interface {
	GetListenerList()
}

func InitListenerRepo(conn *sql.DB) *ListenerRepo {
	return &ListenerRepo{Conn: conn}
}

func (l *ListenerRepo) GetListenerList() {
// TODO
fmt.Println(l.Conn)
}

func ChangeStatusListener(status bool){
// TODO
}
