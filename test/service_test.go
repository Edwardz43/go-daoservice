package test

import (
	"database/sql"
	"flag"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"testing"
	"time"

	dao "my-daoservice/service"

	config "my-daoservice/config"

	faker "github.com/bxcodec/faker"

	_ "github.com/go-sql-driver/mysql"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

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
	dao.Execute(stmt1, faker.Username(), faker.Email(), faker.Password(), time.Now())
}

func TestQuery(t *testing.T) {
	stmt := "SELECT * FROM users;"
	dao.Query(stmt)
}

func BenchmarkQuery(b *testing.B) {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	for i := 0; i < b.N; i++ {
		dao.Query("SELECT * FROM demo.users;")
	}

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close()
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}

func BenchmarkExecute(b *testing.B) {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	stmt1 := "INSERT INTO demo.users (name, email, password, created_at) VALUES(?, ?, ?, ?);"
	for i := 0; i < b.N; i++ {
		dao.Execute(stmt1, faker.Username(), faker.Email(), faker.Password(), time.Now())
	}

	dao.Execute("TRUNCATE TABLE demo.users;")

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close()
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}
