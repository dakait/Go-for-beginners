package main 

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	Id 		int64 	`db:"id"`
	Name 	string 	`db:"name"`
	Email	string 	`db:"email"`
}

type Profile struct {
	Id 			int64 	`db:"id"`
	PersonId    int64 	`db:"person_id"`
	Face 		string 	`db:"face"`
	Hair 		string 	`db:"hair"`
	Person  	`db:"person"`
}

func main() {
	//NOTE: database and tables were created already as the below code is not for creating 
	//any database or tables
	DB, err := sqlx.Connect("mysql", "root:hackinitiator@/dusk")
	if err == nil {
		fmt.Println("success!!")
	} 
	//queries to insert the data
	// result, _ := DB.Exec("INSERT INTO person (name, email) VALUES ('yondu', 'yondu@gotg.com')") 
	// person_id, _ := result.LastInsertId()
	// DB.Exec("INSERT INTO profile (person_id, face, hair) VALUES (?, 'oval', 'no hair')", person_id)
	
	//query to display data using left join
	var q []Profile
	DB.Select(&q, `select person.id "person.id", person.name "person.name", person.email "person.email", profile.* from profile left join person on person.id = profile.person_id`)
	fmt.Println(q)
}