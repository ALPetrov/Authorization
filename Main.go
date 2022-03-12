package main

import (
	"./request"
	"./funcAdd"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type tableUser struct {
	name     string
	lastName string
	login    string
	password string
	deleted  string
}

func main() {

	request.Goodbye()

	funcAdd.Hello()



	db, err := sql.Open("mysql", "root:Lozin3992ka@/testbd")

	if err != nil {
		panic(err)
	}
	defer db.Close()



	//request.Insert()
	
	//result, err := db.Exec("insert into testbd.user (name, lastName,
	//login, password, deleted)
	//values ('Aleks', 'Piatrou', 'PetAL', '12345', 'No')")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(result.RowsAffected())

	rows, err := db.Query("select * from testbd.user")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	users := []tableUser{}

	for rows.Next() {
		p := tableUser{}
		err := rows.Scan(&p.name, &p.lastName, &p.login, &p.password, &p.deleted)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, p)
	}
	for _, p := range users {

		fmt.Println(p.name, p.lastName, p.login, p.password, p.deleted)
	}

}
