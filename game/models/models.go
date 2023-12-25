package models

import (
	"strings"
)

type Text interface {
	ToString() string
}

type World struct {
	Rooms  []Room
	Player Player
}

type Player struct {
	//State     string //Состояние игрока
	Backpack *Backpack //Поле для рюкзака
	Location Room      //Комната в которой сейчас находится игрок
	Quests   []string  //Задания игрока
}

func (p Player) Lookup() string {
	var builder strings.Builder
	if p.Location.Furniture == nil {
		builder.WriteString("пустая комната. ")
	}
	builder.WriteString(p.Location.ToString())
	return builder.String()
}

type Backpack struct {
	basis Item
	Items []Item //Перечень предметов в рюкзаке
}

func (p *Player) PutOn(item Item) string {
	if item.Title == "рюкзак" {
		p.Backpack = new(Backpack)
		return "вы надели: рюкзак"
	}
	return "ну я хуй знает, рюкзак это или че"
}

func (p *Player) TakeItem(item Item) string {
	if p.Backpack == nil {
		return "некуда класть"
	}
	var builder strings.Builder
	builder.WriteString("предмет добавлен в инвентарь: ")
	builder.WriteString(item.Title)
	if p.Location.Furniture != nil {

	}
	p.Backpack.Items = append(p.Backpack.Items, item)
	return builder.String()
}

type Room struct {
	Title       string      //Название комнаты
	Description string      //Обстановка в комнате
	Furniture   []Furniture //Мебель находящаяся в комнате
	Doors       []Door      //Двери в комнате
}

func (r *Room) ToString() string {
	var builder strings.Builder
	builder.WriteString("можно пройти - ")
	for idx, rooms := range r.Doors {
		builder.WriteString(r.Title)
		if idx < len(r.Doors)-1 {
			builder.WriteString(", ")
		} else {
			builder.WriteString(" ")
		}
	}
	return builder.String()
}

type Door struct {
	IsClosed     bool //Состояние двери (открыто/закрыто)
	ItemToUnlock Item //Предмет требующиеся для открытия
	RoomBeside   Room //Комната находящаяся за дверью
}

func (d *Door) Open(item Item) string {
	var builder strings.Builder
	if d.ItemToUnlock == item {
		builder.WriteString("дверь открыта")
		return builder.String()
	}
	builder.WriteString("нельзя применить")
	return builder.String()
}

type Furniture struct {
	Title string //Название Мебели
	Items []Item //Предмета находящиеся на/в мебели
}

func (f Furniture) FindFurniture(furniture []Furniture) []Furniture {
	fSlice := []Furniture{}
	for i := 0; i < len(furniture); i++ {
		fSlice = append(fSlice, furniture[i])
	}
	return fSlice
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
	Title   string //Название предмета
	UseWith string //К чему можно применить предмет
	Suffix  string //Окончание названия предмета при обращении
}

func (i Item) ToString() string {
	return i.Title
}
