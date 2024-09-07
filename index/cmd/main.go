package main

import (
	"Weahter-app/index/pkg/interfaz"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

// Funcion main unicamente crea la interfaz
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Mi Aplicación")

	interfaz.CrearInterfaz(myWindow)
	myWindow.Resize(fyne.NewSize(400, 400))
	myWindow.ShowAndRun()

}
