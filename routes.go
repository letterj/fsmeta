package main

import "net/http"

// Route structure
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is a map of Route
type Routes []Route

// the route map for the app
var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"DeviceIndex",
		"GET",
		"/{customerID}/device",
		DeviceIndex,
	},
	Route{
		"DeviceShow",
		"GET",
		"/{customerID}/device/{deviceID}",
		DeviceShow,
	},
	Route{
		"DeviceCreate",
		"POST",
		"/{customerID}/device/{deviceID}",
		DeviceCreate,
	},
	Route{
		"DeviceDelete",
		"DELETE",
		"/{customerID}/device/{deviceID}",
		DeviceDelete,
	},
	Route{
		"DeviceUpdate",
		"PUT",
		"/{customerID}/device/{deviceID}",
		DeviceUpdate,
	},
}
