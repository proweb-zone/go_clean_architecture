package repository

import (
	"clean/architector/internal/data/postgresql"
	"clean/architector/internal/domain/entitie"
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

type IConsumerRepo interface {
	GetConsumersList() []entitie.ConsumerEntitie
	AddConsumerDb(newConsumer entitie.ConsumerEntitie) (bool, error)
}

type ConsumerRepo struct {
	Conn postgresql.IConnectionDb
}

func InitConsumerRepo() *ConsumerRepo{
	var connDb postgresql.IConnectionDb = postgresql.BuildConnPg()
	return &ConsumerRepo{Conn: connDb}
}

func(c *ConsumerRepo) GetConsumersList() []entitie.ConsumerEntitie {
	fmt.Println("создаем consumer")
	var db *sql.DB = c.Conn.ConnDb()
	rows, errQuery := db.Query("SELECT * FROM consumers where status=1")

	if errQuery != nil {
		panic(errQuery)
	}

	defer rows.Close()

	consumersList := make([]entitie.ConsumerEntitie, 0)
	for rows.Next() {
		consumersItem := new(entitie.ConsumerEntitie)
		err := rows.Scan(&consumersItem.Id, &consumersItem.Name, &consumersItem.Host, &consumersItem.Port, &consumersItem.Deelay, &consumersItem.Status)
		if err != nil {
			log.Fatal(err)
		}

		consumersList = append(consumersList, *consumersItem)
	}

	return consumersList
}


func(c *ConsumerRepo) AddConsumerDb(newConsumer entitie.ConsumerEntitie) (bool, error) {
	fmt.Print("consumerRepo - AddConsumerDb")
	var db *sql.DB = c.Conn.ConnDb()
	var buildFieldQuery string = "name, host, port"
	var buildValueQuery string = "'" + newConsumer.Name + "', '" + newConsumer.Host + "', '" + newConsumer.Port + "'"

	if newConsumer.Deelay > 0 {
		buildFieldQuery += ", deelay"
		buildValueQuery += ", " + strconv.Itoa(newConsumer.Deelay) + ""
	}

	if newConsumer.Status > 0 {
		buildFieldQuery += ", status"
		buildValueQuery += ", " + strconv.Itoa(newConsumer.Status) + ""
	}

	response, err := db.Exec("INSERT INTO consumers (" + buildFieldQuery + ") VALUES (" + buildValueQuery + ");")

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
