package web

import (
	"clean/architector/internal/data/postgresql"
	"clean/architector/internal/domain/repository"
	"clean/architector/internal/domain/usecase"
	"database/sql"
)

func StartListenerHandler(){
var db *sql.DB = postgresql.ConnPg()

var iListenerRepo repository.IlistenerRepo = repository.InitListenerRepo(db)
usecase.StartListeners(iListenerRepo)
}

func AddListener(){
 // TODO
}

func DeleteListener(){
// TODO
}
