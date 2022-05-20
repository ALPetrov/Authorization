package main

import (
	"database/sql"
	"fmt"
	//"github.com/ALPetrov/authorization/request"
	//"github.com/ALPetrov/authorization/funcadd"
	_"github.com/go-sql-driver/mysql"
)

type tableUser struct {
	id int 
	name     string
	lastName string
	login    string
	password string
}
type query struct {
	name  string
	lastName string
	login    string
	password string
}

func main() {

	a := query{
	name: "A4",
	lastName: "P4",
	login: "AP4",
	password: "4",
	}
	
	db, err := sql.Open("mysql", "root:mysql57@/mydb")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := "insert into mydb.user (name, lastName,login, password) values (?, ?, ?, ?)"
	result, err := db.Exec(query, a.name, a.lastName, a.login, a.password)

	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId())

	rows, err := db.Query("SELECT * FROM mydb.user")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	users := []tableUser{}

	for rows.Next() {
		p := tableUser{}
		
		err := rows.Scan(&p.id, &p.name, &p.lastName, &p.login, &p.password)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, p)
	}

	for _, p := range users {

		fmt.Println(p.id, p.name, p.lastName, p.login, p.password)
	}
}
