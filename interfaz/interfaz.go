package interfaz

import (
	"Weahter-app/usuarios"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// Se crea un interfaz de seleccion primero de Pasajero o Servidores
func CrearInterfaz(myWindow fyne.Window) {
	roles := []string{"Pasajero", "Servidores"}

	selectWidget := widget.NewSelect(roles, func(valor string) {
		var usuario usuarios.Entrada
		switch valor {
		case "Pasajero":
			usuario = &usuarios.Pasajero{}
		case "Servidores":
			usuario = &usuarios.Servidor{}
		}
		if usuario != nil {
			subcategoriaroles := []string{"Ticket", "Ita", "Ciudad"}
			subcategoria := widget.NewSelect(subcategoriaroles, func(ingreso string) {
				println("Seleccionaste:", ingreso)
				switch ingreso {
				case "Ticket":
					usuario.ManejarTicket(myWindow)
				case "Ita":
					usuario.ManejarITA(myWindow)
				case "Ciudad":
					usuario.ManejarCiudad(myWindow)
				}
			})
			myWindow.SetContent(subcategoria)
		}
	})

	// Mostrar la ventana
	myWindow.SetContent(selectWidget)
	myWindow.ShowAndRun()
}
