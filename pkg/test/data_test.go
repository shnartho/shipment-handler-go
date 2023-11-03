package data_test

import (
	"testing"

	"www.github.com/shnartho/shipment-handler-go/pkg/data"
)

func TestAddToSlice(t *testing.T) {
	d := data.NewData()
	err := data.AddToSlice(d, 750)

	if err != nil {
		t.Errorf("AddToSlice returned an error: %v", err)
	}

	// Adding the same value should return an error
	err = data.AddToSlice(d, 750)
	if err == nil {
		t.Errorf("AddToSlice should return an error when adding an existing value")
	}
}

func TestRemoveFromSlice(t *testing.T) {
	d := data.NewData()
	data.AddToSlice(d, 250)
	err := data.RemoveFromSlice(d, 250)

	if err != nil {
		t.Errorf("RemoveFromSlice returned an error: %v", err)
	}

	// Removing a non-existent value should return an error
	err = data.RemoveFromSlice(d, 111)
	if err == nil {
		t.Errorf("RemoveFromSlice should return an error when removing a non-existent value")
	}
}

func TestUpdateSlice(t *testing.T) {
	d := data.NewData()
	err := data.UpdateSlice(d, 250, 251)

	if err != nil {
		t.Errorf("UpdateSlice returned an error: %v", err)
	}

	// Updating a non-existent value should return an error
	err = data.UpdateSlice(d, 300, 600)
	if err == nil {
		t.Errorf(" non-existent value")
	}
}

func TestPacksNeeded(t *testing.T) {
	d := data.NewData()
	result := data.PacksNeeded(d, 251)
	expected := "1*500"

	if result != expected {
		t.Errorf("got: %s, expected: %s", result, expected)
	}
}
