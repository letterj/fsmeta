package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"testing"

	_ "github.com/lib/pq"

	"github.com/emicklei/forest"
)

var fsmeta = forest.NewClient("http://localhost:8080", new(http.Client))

func addDevice(t *testing.T, c string, d string, n string) *http.Response {
	tDevice := fmt.Sprintf(`{"name": "%s", "blocksize": 4096, "sizegb": 10}`, n)
	return fsmeta.POST(t, forest.Path("/{customerID}/device/{deviceID}", c, d).Body(tDevice))
}

// setupDB cleans out the database tables
func setupDB() {
	tbls := []string{"fsinode", "fsdevice"}
	tdbinfo := fmt.Sprintf("postgres://%s/%s?sslmode=disable",
		"localhost", "fsdisk")
	tdb, err := sql.Open("postgres", tdbinfo)
	checkErr(err)

	// Truncate tables in the tbls list
	var stmt string
	for t := range tbls {
		stmt = fmt.Sprintf("TRUNCATE TABLE %s;", tbls[t])
		_, err := tdb.Exec(stmt)
		checkErr(err)
	}
	tdb.Close()
}

//Health check test
func Test_fsmetaExists(t *testing.T) {
	r := fsmeta.GET(t, forest.Path("/"))
	forest.ExpectStatus(t, r, 200)
}

//Looking for a device list that doesn't exist
func TestGetDeviceList_NotFound(t *testing.T) {
	//Setup
	setupDB()

	//Actual Test
	var custID = "54321"
	r := fsmeta.GET(t, forest.Path("/{custID}/device", custID))
	forest.ExpectStatus(t, r, 404)
}

//Looking for a list of devices for a specific customer
func TestGetDeviceList(t *testing.T) {
	//Setup
	setupDB()
	var custID = "1234"
	r1 := addDevice(t, custID, "1", "boot")
	r2 := addDevice(t, custID, "2", "data")

	//Actual Test
	rt := fsmeta.GET(t, forest.Path("/{custID}/device", custID))
	forest.ExpectStatus(t, r1, 201)
	forest.ExpectStatus(t, r2, 201)
	forest.ExpectStatus(t, rt, 200)
}

// Looking for a specific device for a specific customer
func TestGetDevice_One(t *testing.T) {
	//Setup
	setupDB()
	var custID = "1234"
	var deviceID = "1"

	r1 := addDevice(t, custID, deviceID, "boot")

	//Actual Test
	rt := fsmeta.GET(t, forest.Path("/{customerID}/device/{deviceID}", custID, deviceID))
	forest.ExpectStatus(t, r1, 201)
	forest.ExpectStatus(t, rt, 200)
}

// Look for a specific device for a specific customer
func TestGetDevice_Two(t *testing.T) {
	//Setup
	setupDB()
	var custID = "1234"
	var deviceID = "2"
	r1 := addDevice(t, custID, "1", "boot")
	r2 := addDevice(t, custID, deviceID, "data")

	//Actual Test
	rt := fsmeta.GET(t, forest.Path("/{customerID}/device/{deviceID}", custID, deviceID))
	forest.ExpectStatus(t, r1, 201)
	forest.ExpectStatus(t, r2, 201)
	forest.ExpectStatus(t, rt, 200)
}

func TestGetDevice_NotFound(t *testing.T) {
	var custID = "1234"
	var deviceID = "999999999"
	r := fsmeta.GET(t, forest.Path("/{customerID}/device/{deviceID}", custID, deviceID))
	forest.ExpectStatus(t, r, 404)
}

func TestCreateDevice_Conflict(t *testing.T) {
	var custID = "1234"
	var deviceID = "1"
	var name = "boot"
	tDevice := fmt.Sprintf(`{"name": "%s", "blocksize": 4096, "sizegb": 10}`, name)
	r := fsmeta.POST(t, forest.Path("/{customerID}/device/{deviceID}", custID, deviceID).Body(tDevice))
	forest.ExpectStatus(t, r, 409)
}

func TestCreateDevice_Created(t *testing.T) {
	var custID = "1234"
	var deviceID = "9"
	var name = "data1"
	tDevice := fmt.Sprintf(`{"name": "%s", "blocksize": 4096, "sizegb": 10}`, name)
	r := fsmeta.POST(t, forest.Path("/{customerID}/device/{deviceID}", custID, deviceID).Body(tDevice))
	forest.ExpectStatus(t, r, 201)
}

func TestDeleteDevice_NotFound(t *testing.T) {
	var custID = "1234"
	var deviceID = "99999"
	r := fsmeta.DELETE(t, forest.Path("/{customerID}/device/{deviceID}", custID, deviceID))
	forest.ExpectStatus(t, r, 404)
}

func TestDeleteDevice_Deleted(t *testing.T) {
	var custID = "1234"
	var deviceID = "9"
	r := fsmeta.DELETE(t, forest.Path("/{customerID}/device/{deviceID}", custID, deviceID))
	forest.ExpectStatus(t, r, 204)
}
