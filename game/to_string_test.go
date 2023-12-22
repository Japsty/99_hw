package main

import "testing"

func TestToStringWithEmptyTitle(t *testing.T) {
	item1 := Item{
		Title:            "Chair",
		FurnitureToApply: nil,
		StorageItem:      "",
	}
	item2 := Item{
		Title:            "Lamp",
		FurnitureToApply: nil,
		StorageItem:      "",
	}
	furniture := Furniture{
		Title: "",
		Items: []Item{
			item1,
			item2,
		},
	}
	expected := ""
	result := furniture.ToString()
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
