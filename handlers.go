package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Index is the base route
func Index(w http.ResponseWriter, r *http.Request) {
	results := "OK"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(results); err != nil {
		panic(err)
	}
}

// DeviceIndex is a route for all devices
func DeviceIndex(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var custdevices Devices
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	customerID, err := strconv.ParseInt(vars["customerID"], 10, 64)
	if err != nil {
		// Need to return a 412 here
		fmt.Println(err)
	}

	custdevices = DstoreFindDevices(customerID)
	if custdevices == nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(custdevices); err != nil {
			panic(err)
		}
	}
}

// DeviceShow is a route a specific device
func DeviceShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var custdevice Device
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	customerID, err := strconv.ParseInt(vars["customerID"], 10, 64)
	if err != nil {
		// Need to return a 412 here
		fmt.Println(err)
	}
	deviceID, err := strconv.ParseInt(vars["deviceID"], 10, 64)
	if err != nil {
		// Need to return a 412 here
		fmt.Println(err)
	}

	custdevice = DstoreFindDevice(customerID, deviceID)
	if custdevice == (Device{}) {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(custdevice); err != nil {
			panic(err)
		}
	}
}

// DeviceCreate adds a new device
func DeviceCreate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var device Device

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	customerID, err := strconv.ParseInt(vars["customerID"], 10, 64)
	if err != nil {
		// Need to return a 412 here
		fmt.Println(err)
	}

	deviceID, err := strconv.ParseInt(vars["deviceID"], 10, 64)
	if err != nil {
		// Need to return a 412 here
		fmt.Println(err)
	}

	if DstoreFindDevice(customerID, deviceID) != (Device{}) {
		w.WriteHeader(http.StatusConflict)
		return
	}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &device); err != nil {
		w.WriteHeader(http.StatusPreconditionFailed) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
			panic(err)
		}
	}

	device.ID = deviceID
	device.CustomerID = customerID

	newdeviceID := DstoreCreateDevice(device)
	if newdeviceID != deviceID {
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(device); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode("OK"); err != nil {
			panic(err)
		}
	}
}

// DeviceDelete deletes a new device
func DeviceDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	customerID, err := strconv.ParseInt(vars["customerID"], 10, 64)
	if err != nil {
		// Need to return a 412 here
		fmt.Println(err)
	}

	deviceID, err := strconv.ParseInt(vars["deviceID"], 10, 64)
	if err != nil {
		// Need to return a 412 here
		fmt.Println(err)
	}

	if DstoreFindDevice(customerID, deviceID) == (Device{}) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	//  Do Delete stuff here
	if DstoreDestroyDevice(customerID, deviceID) {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	return
}

// DeviceUpdate deletes a new device
func DeviceUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	customerID, err := strconv.ParseInt(vars["customerID"], 10, 64)
	if err != nil {
		// Need to return a 412 here
		fmt.Println(err)
	}

	deviceID, err := strconv.ParseInt(vars["deviceID"], 10, 64)
	if err != nil {
		// Need to return a 412 here
		fmt.Println(err)
	}

	if DstoreFindDevice(customerID, deviceID) == (Device{}) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//  Do Delete stuff here
	w.WriteHeader(http.StatusOK)
	return
}
