package main

/*
	код писать в этом файле
	наверняка у вас будут какие-то структуры с методами, глобальные переменные ( тут можно ), функции
*/

type World struct {
	Rooms             []Room
	Items             []Item
	ExistingFurniture []Furniture
}

type Room struct {
	Paths      []Room      //Комнаты в которые можно попасть
	RFurniture []Furniture //Мебель находящаяся в комнате
	Doors      []Door      //Двери в комнате
}
func (r Room) ShowPaths() []Room  {
	return r.Paths
}

type Door struct {
	State        string //Состояние двери (открыто/закрыто)
	ItemToUnlock Item   //Предмет требующиеся для открытия
}

type Furniture struct {
	Items []Item //Предмета находящиеся на/в мебели
}

type Item struct {
	Name             string      //Название предмета
	FurnitureToApply []Furniture //Мебель к которой можно применить объект
	StorageItem      string      //Предмет требующийся для помещения объекта в инвентарь
}

type Command struct {
	Title string //Название действия типо: "осмотреться"
}

func (Command)() {

}

type Player struct {
	State     string //Состояние игрока
	Inventory Item   //Поле для рюкзака
	Location  Room   //Комната в которой сейчас находится игрок
	Commands
}

func main() {
	/*
		в этой функции можно ничего не писать,
		но тогда у вас не будет работать через go run main.go
		очень круто будет сделать построчный ввод команд тут, хотя это и не требуется по заданию
	*/
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

	return "not implemented"
}
