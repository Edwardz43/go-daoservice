package main

import (
	"database/sql"
	"log"
	"my-daoservice/config"

	_ "github.com/go-sql-driver/mysql"
)

// Execute ...
func Execute(stmt string, params ...interface{}) bool {
	db, err := sql.Open(config.DriverName, config.DNS)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmtIns, err := db.Prepare(stmt) // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	// Insert square numbers for 0-24 in the database

	_, err = stmtIns.Exec(params...) // Insert tuples (i, i^2)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	return true
}

// Query queries
func Query(stmt string, params ...interface{}) []interface{} {
	db, err := sql.Open(config.DriverName, config.DNS)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmtOut, err := db.Prepare(stmt)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query(params...)
	if err != nil {
		log.Fatal(err.Error())
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	//Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	//Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	//rows.Scan wants '[]interface{}' as an argument, so we must copy the
	//references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	//log.Println(len(scanArgs))
	data := make([]interface{}, 0)
	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// Now do something with the data.
		// Here we just print each column as a string.
		d := make(map[string]interface{})
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			d[columns[i]] = value
			//fmt.Println(columns[i], ": ", value)
		}
		data = append(data, d)
		//fmt.Println("-----------------------------------")
	}

	// proper error handling instead of panic in your app
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

	return data
}
