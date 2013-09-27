/*
Package db provides the methods to interact with the database.
*/
package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var dbparms string

func init() {
	dbparms = "user=monadm password=password" +
		" host=penfold dbname=monitor port=5433"
}

// Function ClientList returns an array of clients currently configured
// in the database.
func ClientList() []string {
	var clients []string
	pgdb, err := sql.Open("postgres", dbparms)
	if err != nil {
		fmt.Printf("Error connecting to the database: %v\n", err)
		return clients
	} else {
		defer pgdb.Close()
	}

	var cq string = "select name from clients"
	rows, err := pgdb.Query(cq)
	if err != nil {
		fmt.Printf("Query failed: %v\n", err)
		return clients
	}
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			fmt.Printf("No results from query: %v\n", err)
			return clients
		}
		clients = append(clients, name)
	}
	return clients
}
