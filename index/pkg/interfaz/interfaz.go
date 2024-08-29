package interfaz

import (
	"Weahter-app/index/pkg/usuarios"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// En esta funcion se establece la interfaz con los roles de pasajero y servidor
func CrearInterfaz(myWindow fyne.Window) {
	roles := []string{"Pasajero", "Servidor"}

	selectWidget := widget.NewSelect(roles, func(valor string) {
		var usuario usuarios.Entrada
		switch valor {
		case "Pasajero":
			usuario = &usuarios.Pasajero{}
		case "Servidor":
			usuario = &usuarios.Servidor{}
		}
		if usuario != nil { //Se manda a seleccionar las categorias de Ticket, Ita y Ciudad y se manda a llamar
			subcategoriaroles := []string{"Ticket", "Ita", "Ciudad"}
			subcategoria := widget.NewSelect(subcategoriaroles, func(ingreso string) {
				switch ingreso {
				case "Ticket":
					usuario.ManejarTicket(myWindow)
				case "Ita":
					usuario.ManejarITA(myWindow)
				case "Ciudad":
					usuario.ManejarCiudad(myWindow)
				}
				dialog.ShowInformation("Selección realizada", fmt.Sprintf("Seleccionaste: %s", ingreso), myWindow)
			})
			myWindow.SetContent(subcategoria)
		}
	})

	myWindow.SetContent(selectWidget)
}
