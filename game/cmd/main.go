package main

import (
	"fmt"
	. "github.com/Japsty/99_hw/game/models"
)

/*
	код писать в этом файле
	наверняка у вас будут какие-то структуры с методами, глобальные переменные ( тут можно ), функции
*/

func main() {
	/*
		в этой функции можно ничего не писать,
		но тогда у вас не будет работать через go run main.go
		очень круто будет сделать построчный ввод команд тут, хотя это и не требуется по заданию
	*/
	//item1 := Item{
	//	Title:            "лампа",
	//	FurnitureToApply: nil,
	//	StorageItem:      "",
	//}
	//item2 := Item{
	//	Title:            "ведро",
	//	FurnitureToApply: nil,
	//	StorageItem:      "",
	//}
	//furniture := Furniture{
	//	Title: "стол",
	//	Items: []Item{item1, item2},
	//}
	//furniture2 := Furniture{
	//	Title: "кровать",
	//	Items: []Item{item2, item1},
	//}
	kitchen := Room{
		Title:     "кухня",
		Paths:     nil,
		Furniture: nil,
		Doors:     nil,
	}
	sleepingroom := Room{
		Title:     "спальня",
		Paths:     nil,
		Furniture: nil,
		Doors:     nil,
	}
	livingroom := Room{
		Title:     "гостинная",
		Paths:     []Room{kitchen, sleepingroom},
		Furniture: nil,
		Doors:     nil,
	}
	p := Player{
		Inventory: nil,
		Location:  livingroom,
		Quests:    nil,
	}
	fmt.Println(p.Lookup())
}
