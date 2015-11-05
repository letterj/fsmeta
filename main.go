package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB // global variable to share it between main and the HTTP handler

func main() {

	var err error

	//dbsrc := dbInfo{5432, "localhost", "fsdisk", "jay1279", "jay1279", "disable"}
	//dbsource := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
	//	dbsrc.user, dbsrc.password, dbsrc.location, dbsrc.name)
	dbsrc := dbInfo{5432, "localhost", "circle_test", "ubuntu", "", "disable"}
	dbsource := fmt.Sprintf("postgres://%s/%s?sslmode=disable",
		dbsrc.location, dbsrc.name)

	fmt.Println(dbsource)

	db = DbConnect(dbsource)
	db.SetMaxIdleConns(5)

	err = db.Ping() // This DOES open a connection if necessary. This makes sure the database is accessible
	if err != nil {
		log.Fatalf("Error on opening database connection: %s", err.Error())
	}

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
