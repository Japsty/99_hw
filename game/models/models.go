package models

import (
	"strings"
)

type World struct {
	Rooms  []Room
	Player Player
	Step   int
}

type Player struct {
	//State     string //Состояние игрока
	Backpack *Backpack //Поле для рюкзака
	Location Room      //Комната в которой сейчас находится игрок
	Quests   []string  //Задания игрока
}

func NewPlayer(backpack *Backpack, location Room, quests []string) Player {
	player := Player{
		Backpack: backpack,
		Location: location,
		Quests:   quests,
	}
	return player
}

// Команда "осмотреться"
func (p Player) Lookup() string {
	var builder strings.Builder
	counter := 0
	if len(p.Location.Furnitures) == 0 {
		builder.WriteString("пустая комната. ")
		return builder.String()
	} else {
		for key, items := range p.Location.Furnitures {
			counter++
			builder.WriteString("на ")
			builder.WriteString(key)
			builder.WriteString(": ")
			for _, item := range items {
				builder.WriteString(item.Title)
				if counter == len(p.Location.Furnitures) {
					builder.WriteString(". ")
				} else {
					builder.WriteString(", ")
				}
			}
		}
	}
	builder.WriteString(p.Location.ToString())
	return builder.String()
}

// Команда "идти (комната)"
func (p *Player) GoTo(RoomTitle string) string {
	var builder strings.Builder
	for key, room := range p.Location.Paths {
		if key == RoomTitle && room.Locked == false {
			p.Location = room
			builder.WriteString(p.Location.Description)
			builder.WriteString(p.Location.ToString())
			return builder.String()
		} else if key == RoomTitle && room.Locked == true {
			builder.WriteString("дверь закрыта")
			return builder.String()
		}
	}
	builder.WriteString("нет пути в ")
	builder.WriteString(RoomTitle)
	return builder.String()
}

// Команда "взять (предмет)"
func (p *Player) TakeIt(ItemName string) string {
	var builder strings.Builder
	if p.Backpack.State == false {
		builder.WriteString("некуда класть")
		return builder.String()
	}
	for _, items := range p.Location.Furnitures {
		for _, item := range items {
			if item.Title == ItemName {
				builder.WriteString("предмет добавлен в инвентарь: ")
				builder.WriteString(item.Title)
				return builder.String()
			}
		}
	}
	builder.WriteString("нет такого")
	return builder.String()
}

// Команда "надеть (предмет)"
func (p *Player) PutOn(ItemName string) string {
	var builder strings.Builder
	for _, items := range p.Location.Furnitures {
		for _, item := range items {
			if item.Title == ItemName {
				if item.Title == "рюкзак" {
					p.Backpack.State = true
					p.Backpack.Title = item.Title
				}
				builder.WriteString("вы надели: ")
				builder.WriteString(item.Title)
				return builder.String()
			}
		}
	}
	builder.WriteString("нет такого")
	return builder.String()
}

// Команда "применить (предмет) (мебель)"
func (p *Player) Apply(ItemName string, FurnitureName string) string {
	var builder strings.Builder
	findItem := false
	findFurniture := false
	for key, items := range p.Backpack.Items {
		if key == ItemName {
			findItem = true
		}
	}
	for key, _ := range p.Location.Furnitures {
		if key == FurnitureName {
			findFurniture = true
		}
	}
	if findItem == true && findFurniture != true {
		builder.WriteString("не к чему применить")
	} else if findItem != true {
		builder.WriteString("нет предмета в инвентаре - ")
		builder.WriteString(ItemName)
	}
	return builder.String()
}

// Рюкзак как айтем и рюкзак на игроке - разные сущности
type Backpack struct {
	State bool            //состояние рюкзака
	Title string          //название рюкзака
	Items map[string]Item //Перечень предметов в рюкзаке
}

type Room struct {
	Title       string            //Название комнаты
	Description string            //Обстановка в комнате
	Locked      bool              //Закрыта ли комната
	Paths       map[string]Room   // Пути куда можно пройти
	Furnitures  map[string][]Item //Мапа мебели
}

func NewRoom(Title string, Description string, Locked bool, Paths map[string]Room, Furnitures map[string][]Item) Room {
	room := Room{
		Title:       Title,
		Description: Description,
		Locked:      Locked,
		Paths:       Paths,
		Furnitures:  Furnitures,
	}
	return room
}

// Выводит можно пройти - (перечень комнат в которые есть проход из данной комнаты)
func (r Room) ToString() string {
	var builder strings.Builder
	builder.WriteString("можно пройти - ")
	counter := 0
	for key := range r.Paths {
		counter += 1
		builder.WriteString(key)
		if counter == len(r.Paths) {
			return builder.String()
		}
		builder.WriteString(", ")
	}
	return builder.String()
}

type Item struct {
	Title   string //Название предмета
	UseWith string //К чему можно применить предмет
}

func NewItem(Title string, UseWith string) Item {
	item := Item{
		Title:   Title,
		UseWith: UseWith,
	}
	return item
}
