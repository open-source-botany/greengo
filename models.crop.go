package main

import (
	"errors"
	"strconv"
)

type crop struct {
	ID          int    `json:"id"`
	Environment string `json:"environment"`
	Name        string `json:"name"`
	Notes       string `json:"notes"`
	Qty         int    `json:"qty"`
	StartDate   string `json:"start_date"`
	HarvestDate string `json:"harvest_date"`
}

var cropList = []crop{
	crop{ID: 1, Name: "Kale", Environment: "outdoors", Qty: 3, StartDate: "05-01-2018"},
	crop{ID: 2, Name: "Cilantro", Environment: "outdoors", Qty: 2, StartDate: "05-01-2018"},
}

// Return a list of all the crops
func getAllCrops() []crop {
	return cropList
}

func getCropByID(id int) (*crop, error) {
	for _, c := range cropList {
		if c.ID == id {
			return &c, nil
		}
	}
	return nil, errors.New("Crop not found")
}

func createNewCrop(name, start_date, environment, qty, notes string) (*crop, error) {
	if amt, err := strconv.Atoi(qty); err == nil {
		c := crop{ID: len(cropList) + 1, Name: name, StartDate: start_date, Environment: environment, Qty: amt, Notes: notes}

		cropList = append(cropList, c)

		return &c, nil
	}
	return nil, nil
}
