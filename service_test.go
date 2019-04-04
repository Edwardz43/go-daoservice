package main

import (
	"database/sql"
	"testing"
	"time"

	config "my-daoservice/config"

	faker "github.com/bxcodec/faker"

	_ "github.com/go-sql-driver/mysql"
)

func TestDriverName(t *testing.T) {
	if config.DriverName == "mysql" {
		t.Log("test PASS")
	} else {
		t.Error("test FAIL")
	}
}

func TestOpen(t *testing.T) {
	db, err := sql.Open(config.DriverName, config.DNS)
	if err != nil {
		t.Error("test FAIL")
	} else {
		t.Log("test PASS")
	}
	db.Close()
}

func TestConnection(t *testing.T) {
	db, err := sql.Open(config.DriverName, config.DNS)
	if err != nil {
		t.Error("test FAIL")
	} else {
		t.Log("test PASS")
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		t.Error("test FAIL")
	} else {
		t.Log("test PASS")
	}
}

func TestExecute(t *testing.T) {
	stmt1 := "INSERT INTO demo.users (name, email, password, created_at) VALUES(?,?,?,?);"
	Execute(stmt1, faker.Username(), faker.Email(), faker.Password(), time.Now())
}

func TestQuery(t *testing.T) {
	stmt := "SELECT * FROM users;"
	Query(stmt)
}

func BenchmarkQuery(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Query("SELECT * FROM demo.users;")
	}
}

func BenchmarkExecute(b *testing.B) {
	stmt1 := "INSERT INTO demo.users (name, email, password, created_at) VALUES(?, ?, ?, ?);"
	for i := 0; i < b.N; i++ {
		Execute(stmt1, faker.Username(), faker.Email(), faker.Password(), time.Now())
	}

	//Execute("TRUNCATE TABLE demo.users;")
}
