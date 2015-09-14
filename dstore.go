package main

import "fmt"

var currentID int64

var devices Devices

// Give us some seed data
func init() {
	DstoreCreateDevice(
		Device{
			ID:         1,
			Name:       "boot",
			CustomerID: 1234,
			BlockSize:  4096,
			SizeGB:     10,
		})
	DstoreCreateDevice(
		Device{
			ID:         2,
			Name:       "data",
			CustomerID: 1234,
			BlockSize:  4096,
			SizeGB:     10,
		})
}

// DstoreFindDevices queries for all a customers devices
func DstoreFindDevices(cust int64) Devices {
	var results Devices

	for _, d := range devices {
		if d.CustomerID == cust {
			results = append(results, d)
		}
	}
	// return empty Device if not found
	return results
}

// DstoreFindDevice queries for a specific device
func DstoreFindDevice(cust int64, id int64) Device {
	for _, d := range devices {
		if d.ID == id && d.CustomerID == cust {
			return d
		}
	}
	// return empty Device if not found
	return Device{}
}

// DstoreCreateDevice will create a new device
func DstoreCreateDevice(d Device) Device {
	currentID++
	d.ID = currentID
	devices = append(devices, d)
	return d
}

// DstoreDestroyDevice will delete a device
func DstoreDestroyDevice(id int64) error {
	for i, d := range devices {
		if d.ID == id {
			devices = append(devices[:i], devices[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Device with id of %d to delete", id)
}
