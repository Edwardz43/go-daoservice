package main

import (
	"fmt"
	"time"
)

func main() {

	// insert a demo user
	insertStmt := "INSERT INTO demo.users (name, email, password, created_at) VALUES(?,?,?,?);"

	Execute(insertStmt, "demoUser", "faker@test.com", "passwd", time.Now())

	// select user by user name
	queryStmt := "SELECT * FROM demo.users WHERE name = ?;"

	data := Query(queryStmt, "demoUser")

	// convert interface to map
	user, ok := data[0].(map[string]interface{})
	if !ok {
		fmt.Println("Some error")
	}
	for k, v := range user {
		fmt.Printf("%s: %s\n", k, v)
	}

	// delete demo user
	deleteStmt := "DELETE FROM demo.users WHERE name = ?;"

	Execute(deleteStmt, "demoUser")
}
