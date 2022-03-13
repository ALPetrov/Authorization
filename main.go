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
	deleted  string
}
type query struct {
	name  string
	lastName string
	login    string
	password string
	deleted  string
}

func main() {
a := query{
	name: "A3",
	lastName: "P3",
	login: "AP3",
	password: "3",
	deleted: "No", 
}

fmt.Println(a)
	//request.Goodbye()
	//funcadd.Hello()
	
	db, err := sql.Open("mysql", "root:Lozin3992ka@/testbd")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := fmt.Sprintf("insert into testbd.table1 (name, lastName,login, password, deleted) values (%s, %s, %s, %s, %s)", a.name, a.lastName, a.login, a.password, a.deleted)
	result, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId())

	rows, err := db.Query("SELECT * FROM testbd.table1")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	users := []tableUser{}


	for rows.Next() {
		p := tableUser{}
		
		err := rows.Scan(&p.id, &p.name, &p.lastName, &p.login, &p.password, &p.deleted)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, p)
	}

	for _, p := range users {

		fmt.Println(p.id, p.name, p.lastName, p.login, p.password, p.deleted)
	}
}
