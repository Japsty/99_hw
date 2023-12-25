package internal

func HandleCommand(command string) string {
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
