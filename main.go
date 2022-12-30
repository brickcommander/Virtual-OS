package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var myOS fyne.App = app.New()
var homeWindow fyne.Window = myOS.NewWindow("Virtual OS")

var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var btn4 fyne.Widget

var img fyne.CanvasObject
var deskBtn fyne.Widget

var panelContent *fyne.Container

func main() {
	myOS.Settings().SetTheme(theme.LightTheme())
	img = canvas.NewImageFromFile("E:\\VS Code\\Virtual OS\\667008.png")

	btn1 = widget.NewButtonWithIcon("Weather App", theme.InfoIcon(), func() {
		showWeatherApp()
	})

	btn2 = widget.NewButtonWithIcon("Calculator", theme.ContentAddIcon(), func() {
		showCalculator()
	})

	btn3 = widget.NewButtonWithIcon("Gallery", theme.SearchIcon(), func() {
		showGallery()
	})

	btn4 = widget.NewButtonWithIcon("Text Editor", theme.DocumentIcon(), func() {
		showTextEditor()
	})

	deskBtn = widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
		homeWindow.SetContent(container.NewBorder(panelContent, nil, nil, nil, img))
	})

	panelContent = container.NewVBox((container.NewGridWithColumns(5, deskBtn, btn1, btn2, btn3, btn4)))

	homeWindow.Resize(fyne.NewSize(800, 500))
	homeWindow.CenterOnScreen()
	homeWindow.SetContent(container.NewBorder(panelContent, nil, nil, nil, img))
	homeWindow.ShowAndRun()
}
