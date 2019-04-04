package main

import (
	"database/sql"
	"my-daoservice/config"

	_ "github.com/go-sql-driver/mysql"
)

// Execute executes prepare statment without any return rows (e.g. create/update/delete).
// The return value will be true if execution success.
func Execute(stmt string, params ...interface{}) bool {
	db, err := sql.Open(config.DriverName, config.DNS)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmtIns, err := db.Prepare(stmt)
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(params...)
	if err != nil {
		panic(err.Error())
	}

	return true
}

// Query executes a query that returns row(select).
// The return data will be the type map[string]interface{}.
func Query(stmt string, params ...interface{}) []interface{} {
	db, err := sql.Open(config.DriverName, config.DNS)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmtOut, err := db.Prepare(stmt)
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query(params...)
	if err != nil {
		panic(err.Error())
	}

	//Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
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
		}
		data = append(data, d)
	}

	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

	return data
}
