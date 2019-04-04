package main

import (
	"fmt"
	dao "my-daoservice/service"
	"time"
)

func main() {

	// insert a demo user
	insertStmt := "INSERT INTO demo.users (name, email, password, created_at) VALUES(?,?,?,?);"

	ok := dao.Execute(insertStmt, "demoUser", "faker@test.com", "passwd", time.Now())
	if !ok {
		panic("Some error")
	}

	// select user by user name
	queryStmt := "SELECT * FROM demo.users WHERE name = ?;"

	data := dao.Query(queryStmt, "demoUser")

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

	dao.Execute(deleteStmt, "demoUser")
}
