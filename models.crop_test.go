package main

import "testing"

// Test the function that fetches all crops
func TestGetAllCrops(t *testing.T) {
	clist := getAllCrops()

	if len(clist) != len(cropList) {
		t.Fail()
	}

	// Check that each member is identical
	for i, v := range clist {
		if v.Environment != cropList[i].Environment ||
			v.ID != cropList[i].ID ||
			v.Qty != cropList[i].Qty ||
			v.StartDate != cropList[i].StartDate ||
			v.Name != cropList[i].Name {

			t.Fail()
			break
		}
	}
	if cropList[0].ID != 1 {
		t.Fail()
	}
}
