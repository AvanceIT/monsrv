/*
Package db provides the methods to interact with the database.
*/
package db

import (
	"database/sql"
	"fmt"
	"github.com/AvanceIT/monitor/xmltools"
	_ "github.com/lib/pq"
)

var dbparms string
var rdbms string

func init() {
	// Database settings - to change database just import the
	// appropriate driver and change these settings.
	dbparms = "user=monadm password=password" +
		" host=penfold dbname=monitor port=5433"
	rdbms = "postgres"
}

// Function ClientList returns an array of clients currently configured
// in the database.
func ClientList() []string {
	var clients []string
	db, err := sql.Open(rdbms, dbparms)
	if err != nil {
		fmt.Printf("Error connecting to the database: %v\n", err)
		return clients
	} else {
		defer db.Close()
	}

	var cq string = "select name from clients"
	rows, err := db.Query(cq)
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

// Function AddEvent inserts the event details from the client into the
// events table in the database.
func AddEvent(ev xmltools.MonResult) {
	var stmt string
	var clientid string
	var hostid string
	db, err := sql.Open(rdbms, dbparms)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return
	} else {
		defer db.Close()
	}

	// Get the host id and client id for the host that has raised
	// the alert.
	stmt = "select id,client_id from hosts where name='" + ev.HostName + 
		"';"
	rows, err := db.Query(stmt)
	fmt.Printf("%s\n", stmt)
	if err != nil {
		fmt.Printf("Query failed: %v\n", err)
		return
	}
	for rows.Next() {
		err = rows.Scan(&hostid, &clientid)
		if err != nil {
			fmt.Printf("No results from query: %v\n", err)
			return
		}
	}
	fmt.Printf("%s, %s\n", hostid, clientid)

	// Insert the event into the events table.
	stmt = "insert into events(host_id, client_id, time_logged," +
		"time_alerted, message) values(" + hostid + ", " +
		clientid + ", '" + ev.TimeRcvd + "', '" + 
		ev.TimeRptd + "', '" + ev.Detail + "');"
	fmt.Printf("\n\n\n%s\n\n\n", stmt)
	rows, err = db.Query(stmt)
	if err != nil {
		fmt.Printf("Error inserting into db: %v\n", err)
		return
	}
}
