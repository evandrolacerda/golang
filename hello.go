package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const Driver = "teste"
const User = "teste"
const Host = "localhost"
const Port = "5432"
const Password = "teste"
const DbName = "teste"

var db *sql.DB
var err error

func main() {
	var DataSourceString string = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)

	fmt.Printf("Acessando o Banco de dados %s", DbName)

	db, err = sql.Open(Driver, DataSourceString)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("Conectado")
	}

	rows, err := db.Query("SELECT id, name, email FROM users")

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	var count int = 0

	fileHandler, err := os.Create("/tmp/teste.txt")

	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(fileHandler)

	defer fileHandler.Close()

	for rows.Next() {
		var id int
		var name string
		var email string

		err = rows.Scan(&id, &name, &email)
		fmt.Fprintln(writer, id, name, email)
		fmt.Println(id, name, email)

	}

	fmt.Printf("Há %d ocorrências do nome Evandro \n", count)

}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
