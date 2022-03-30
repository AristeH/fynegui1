package main

import "fyne.io/fyne/v2/widget"

type entryForm struct {
	Name    string        // Уникальное имя
	Title   string        // Текст метки
	Value   string        // Значение поля
	Tip     string        // Тип поля
	Buttons string        // кнопки 0-чтение, 1 - просмотр, 2 - редактирование
	Format  string        //формат вывода
	Widget  *widget.Entry //widget
}

