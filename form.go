package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type ButtonData struct {
	Parameters string // параметры кнопки формат Имяпараметра:Значение;... ИмяпараметраN:ЗначениеN;
	Fun        string
	Container  string         // функция выполняемая при нажатии на кнопку
	Widget     *widget.Button // отображение кнопки на экране
}

//FormData - данные формы
type FormData struct {
	ID        string                       // ID - ГУИД формы
	Entry     map[string]entryForm         // Entry  - список полей ввода формы
	Table     map[string]*TableOtoko       // Table  - список таблиц формы
	Tree      map[string]*TreeOtoko        // Table  - список таблиц формы
	Button    map[string]ButtonData        // Button - список кнопок формы
	Container map[string]fyne.CanvasObject // Container - список контейнеров формы
	form      map[string][]string          // form иерархия контейнеров формы
	W         fyne.Window
}

//findButton - найдем кнопку в списке виджетов формы
// возвращает ссылку на форму FormData, и на кнопку ButtonData
func findButton(d *widget.Button) (*FormData, *ButtonData) {
	for _, f := range appValues {
		for _, b := range f.Button {
			if b.Widget == d {
				return f, &b
			}
		}
	}
	return &FormData{}, &ButtonData{}
}

// ToolBarCreate - создание командной панели
//
func ToolBarCreate(idform string, c string, but [][]string, color color.Color) *fyne.Container {

	// Получим форму с данными
	fd := appValues[idform]
	// создадим кнопки формы
	con := container.NewHBox()
	for _, value := range but {
		d := widget.NewButtonWithIcon(value[Name], GetIcon(value[Icon]), nil)
		d.OnTapped = func() {
			var br [][]string
			f, b := findButton(d)
			output := make([]string, 12)

			output[ID] = f.ID                 // гуид подсистмы/кнопки
			output[Parameters] = b.Parameters //параметрыфункции
			output[Fun] = b.Fun               // функция для кнопки

			br = append(br, output)

			d := GetData{ID: idform, Data: br, Container: c}
			UpdateContainer(d)
			//SendMessage("GetDataContainer", d)
		}
		//ff := strings.Split(value[Name],":")
		// сохраним кнопку в FormData
		fd.Button[value[ID]] = ButtonData{Fun: value[Fun], Container: c, Parameters: value[Parameters], Widget: d}
		con.Add(d)
	}
	return container.New(layout.NewMaxLayout(), canvas.NewRectangle(color), con)
}

func GetIcon(n string) fyne.Resource {
	switch n {
	case "DocumentCreateIcon":
		return theme.DocumentCreateIcon()
	case "AccountIcon":
		return theme.AccountIcon()
	}
	return nil
}
