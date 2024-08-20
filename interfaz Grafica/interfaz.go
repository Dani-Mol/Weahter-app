package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Crear una nueva aplicación
	myApp := app.New()

	// Crear una nueva ventana
	myWindow := myApp.NewWindow("Hola Fyne")

	// Crear un botón con un texto y una función que se ejecuta al ser presionado
	helloButton := widget.NewButton("Presióname", func() {
		myWindow.SetContent(widget.NewLabel("¡Hola Mundo!"))
	})

	// Establecer el contenido de la ventana como el botón
	myWindow.SetContent(container.NewVBox(
		helloButton,
	))

	// Mostrar la ventana
	myWindow.ShowAndRun()
}
