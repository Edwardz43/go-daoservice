package main

import (
	"fmt"

	"github.com/Edwardz43/go-daoservice/service"

	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// insert a demo user
	insertStmt := "INSERT INTO demo.users (name, email, password, created_at) VALUES(?,?,?,?);"

	ok := service.Execute(insertStmt, "demoUser", "faker@test.com", "passwd", time.Now())
	if !ok {
		panic("Some error")
	}

	// select user by user name
	queryStmt := "SELECT * FROM demo.users WHERE name = ?;"

	data := service.Query(queryStmt, "demoUser")

	// convert interface to map
	user, ok := data[0].(map[string]interface{})
	if !ok {
		panic("Some error")
	}
	for k, v := range user {
		fmt.Printf("%s: %s\n", k, v)
	}

	// delete demo user
	deleteStmt := "DELETE FROM demo.users WHERE name = ?;"

	service.Execute(deleteStmt, "demoUser")
}
