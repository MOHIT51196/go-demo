package main

import (
	"fmt"
	"strings"
)

/*
Panic : Similar to error in other languages which would halt
the current subroutine and return the control flow to the caller.

Recover : Similar to catch which can handle the halt by catching
panic object. Generally used in a deferred function. Recover() returns
nil if no panic happens before it.
*/

type DB struct{}

func (db DB) CreateConnection(dsn string) {
	dbName := dsn[strings.LastIndex(dsn, "/")+1:]
	if strings.Contains(dsn, "3306") {
		panic(fmt.Sprintf("Unable to connected to DB %s", dbName))
	}
	fmt.Println("Connected to the DB", dbName)
}

// Can also use named variables to return values
// but 'return' statement is required
func InitDB(dsn string) (db DB, status string) {
	db = DB{}
	// handle panic, hence need to be placed before panic call
	// defer calls are invoked at the end of function
	defer func() {
		if panicMsg := recover(); panicMsg != nil {
			fmt.Println(panicMsg)
			status = "Failed"
		}
	}()
	db.CreateConnection(dsn)
	fmt.Println("Initialisation successfull")
	status = "Success"
	return
}
