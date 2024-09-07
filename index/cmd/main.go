package main

import (
	"Weather-app/index/usuarios" // Ajusta esto según la estructura de tu proyecto

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Esta es la función principal donde se ejecuta la aplicación
func main() {
	// Crea una nueva aplicación Fyne
	a := app.New()

	interfaz.CrearInterfaz(myWindow)
	myWindow.Resize(fyne.NewSize(400, 400))
	myWindow.ShowAndRun()
	// Crea una nueva ventana
	w := a.NewWindow("App de Tickets")

	// Aquí agregamos un botón para manejar los tickets
	button := widget.NewButton("Manejar Ticket", func() {
		// Llamamos a la función del paquete usuarios para manejar el ticket
		usuario := usuarios.Usuario{}
		usuario.ManejarTicket(w, func() {
			// Aquí puedes definir lo que ocurre cuando se cancela
		})
	})

	// Configuramos el contenido de la ventana con el botón
	w.SetContent(container.NewVBox(button))

	// Mostramos la ventana
	w.ShowAndRun()
}
