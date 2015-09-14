package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/emicklei/forest"
)

var fsmeta = forest.NewClient("http://localhost:8080", new(http.Client))

func Test_fsmetaExists(t *testing.T) {
	r := fsmeta.GET(t, forest.Path("/"))
	forest.ExpectStatus(t, r, 200)
}

func TestGetDeviceList_NotFound(t *testing.T) {
	var custID = "54321"
	r := fsmeta.GET(t, forest.Path("/{custID}/device", custID))
	forest.ExpectStatus(t, r, 404)
}

func TestGetDeviceList(t *testing.T) {
	var custID = "1234"
	r := fsmeta.GET(t, forest.Path("/{custID}/device", custID))
	forest.ExpectStatus(t, r, 200)
}

func TestGetDevice_One(t *testing.T) {
	var custID = "1234"
	var deviceID = "1"
	r := fsmeta.GET(t, forest.Path("/{customerID}/device/{deviceID}", custID, deviceID))
	forest.ExpectStatus(t, r, 200)
}

func TestGetDevice_Two(t *testing.T) {
	var custID = "1234"
	var deviceID = "2"
	r := fsmeta.GET(t, forest.Path("/{customerID}/device/{deviceID}", custID, deviceID))
	forest.ExpectStatus(t, r, 200)
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
