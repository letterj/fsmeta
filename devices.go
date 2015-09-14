package main

// Device is the structure of an fs device
type Device struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	CustomerID int64  `json:"customerid"`
	BlockSize  int32  `json:"blocksize"`
	SizeGB     int64  `json:"sizegb"`
}

// Devices is an map of devices
type Devices []Device
