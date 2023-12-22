package main

import (
	"fmt"
	"strings"
)

/*
	код писать в этом файле
	наверняка у вас будут какие-то структуры с методами, глобальные переменные ( тут можно ), функции
*/

type Text interface {
	ToString() string
}

type World struct {
	Rooms  []Room
	Player Player
}

type Player struct {
	//State     string //Состояние игрока
	Inventory Item     //Поле для рюкзака
	Location  Room     //Комната в которой сейчас находится игрок
	Quests    []string //Задания игрока
}

func (p Player) Lookup() string {
	var builder strings.Builder
	if p.Location.Furniture == nil {
		builder.WriteString("пустая комната. ")
	}
	builder.WriteString(p.Location.ToString())
	return builder.String()
}

type Room struct {
	Title       string      //Название комнаты
	Description string      //Обстановка в комнате
	Paths       []Room      //Комнаты в которые можно попасть
	Furniture   []Furniture //Мебель находящаяся в комнате
	Doors       []Door      //Двери в комнате
}

func (r *Room) ToString() string {
	var builder strings.Builder
	builder.WriteString("можно пройти - ")
	for idx, rooms := range r.Paths {
		builder.WriteString(rooms.Title)
		if idx < len(r.Paths)-1 {
			builder.WriteString(", ")
		} else {
			builder.WriteString(" ")
		}
	}
	return builder.String()
}

type Door struct {
	State        string //Состояние двери (открыто/закрыто)
	ItemToUnlock Item   //Предмет требующиеся для открытия
}

type Furniture struct {
	Title string //Название Мебели
	Items []Item //Предмета находящиеся на/в мебели
}

func (f *Furniture) ToString() string {
	var builder strings.Builder
	builder.WriteString("на")
	builder.WriteString(" ")
	builder.WriteString(f.Title)
	builder.WriteString("е: ")
	for idx, item := range f.Items {
		builder.WriteString(item.ToString())
		if idx < len(f.Items)-1 {
			builder.WriteString(", ")
		} else {
			builder.WriteString(".")
		}
	}
	return builder.String()
}

type Item struct {
	Title            string   //Название предмета
	FurnitureToApply []string //Мебель к которой можно применить объект
	StorageItem      string   //Предмет требующийся для помещения объекта в инвентарь
	Suffix           string   //Окончание названия предмета при обращении
}

func (i Item) ToString() string {
	return i.Title
}

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
		Inventory: Item{},
		Location:  livingroom,
		Quests:    nil,
	}
	fmt.Println(p.Lookup())
}

func initGame() {
	/*
		эта функция инициализирует игровой мир - все комнаты
		если что-то было - оно корректно перезатирается
	*/
}

func handleCommand(command string) string {
	/*
		данная функция принимает команду от "пользователя"
		и наверняка вызывает какой-то другой метод или функцию у "мира" - списка комнат
	*/
	switch command {
	case "осмотреться":

	case "идти":
		fallthrough
	case "применить":
		fallthrough
	case "взять":
		fallthrough
	default:
		return "неизвестная команда"
	}

	return "not implemented"
}
