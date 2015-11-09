package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type dbInfo struct {
	port     int
	location string
	name     string
	user     string
	password string
	sslmode  string
}

var devices Devices

// checkErr is used to simplify coding
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//DbConnect is used to create a databsae connection
func DbConnect(connectionString string) *sql.DB {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Printf("Connection error: %s\n", err.Error())
		return nil
	}
	return db
}

// DstoreFindDevices queries for all a customers devices
func DstoreFindDevices(cust int64) Devices {
	var results Devices

	sqlstmt := "SELECT id, name, customerid, blocksize, sizegb FROM fsdevice WHERE customerid = $1;"
	rows, err := db.Query(sqlstmt, cust)
	checkErr(err)

	for rows.Next() {
		var d Device
		err = rows.Scan(&d.ID, &d.Name, &d.CustomerID, &d.BlockSize, &d.SizeGB)
		checkErr(err)
		results = append(results, d)
	}
	// return empty Device if not found
	return results
}

// DstoreFindDevice queries for a specific device
func DstoreFindDevice(cust int64, id int64) Device {
	var d Device

	sqlstmt := "SELECT id, name, customerid, blocksize, sizegb FROM fsdevice WHERE customerid = $1 and id = $2;"
	rows, err := db.Query(sqlstmt, cust, id)
	checkErr(err)
	for rows.Next() {
		err = rows.Scan(&d.ID, &d.Name, &d.CustomerID, &d.BlockSize, &d.SizeGB)
		checkErr(err)
		return d
	}
	// return empty Device if not found
	return Device{}
}

// DstoreCreateDevice will create a new device
func DstoreCreateDevice(d Device) int64 {
	var newid int64

	sqlstmt := "INSERT INTO fsdevice (id, name, customerid, blocksize, sizegb) VALUES($1, $2, $3, $4, $5) RETURNING id;"
	rows, err := db.Query(sqlstmt, d.ID, d.Name, d.CustomerID, d.BlockSize, d.SizeGB)
	checkErr(err)
	for rows.Next() {
		err = rows.Scan(&newid)
		checkErr(err)
	}
	return newid
}

// DstoreDestroyDevice will delete a device
func DstoreDestroyDevice(custid int64, deviceid int64) bool {

	sqlstmt := "DELETE FROM fsdevice WHERE customerid = $1 and id = $2;"
	rows, err := db.Exec(sqlstmt, custid, deviceid)
	checkErr(err)
	affected, err := rows.RowsAffected()
	checkErr(err)
	if affected == 1 {
		return true
	}
	fmt.Printf("There was a problem deleteing device %d for customer %d.  %d rows were returned instead of 1.", deviceid, custid, affected)
	return false
}
