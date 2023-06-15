package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:master@tcp(127.0.0.1:3306)/go_db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Connected")

	// create, err := db.Query("CREATE TABLE if not exists TEST(ID INT PRIMARY KEY AUTO_INCREMENT, NAME TEXT, SALARY INT)")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer create.Close()
	// fmt.Println("Created table")

	// insert, err := db.Query("INSERT INTO TEST VALUES(1, 'Kavi', 240000)")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()
	// fmt.Println("Inserted")

	results, err := db.Query("SELECT * FROM TEST")
	if err != nil {
		panic(err.Error())
	}

	defer results.Close()

	// for results.Next() {
	// 	var user User;
	// 	err = results.Scan(&user.Name)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	fmt.Println(user.Name)
	// }

}

type User struct {
	Name string `json:"name"`
}