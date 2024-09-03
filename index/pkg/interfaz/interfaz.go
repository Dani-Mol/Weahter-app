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
	InterfazPrincipal(myWindow)
	usuario := &usuarios.Usuario{}
	subcategoriaroles := []string{"Ticket", "Ita", "Ciudad"}
	subcategoria := widget.NewSelect(subcategoriaroles, func(ingreso string) {
		switch ingreso {
		case "Ticket":
			usuario.ManejarTicket(myWindow, func() {
				CrearInterfaz(myWindow)
			})
		case "Ita":
			usuario.ManejarITA(myWindow, func() {
				CrearInterfaz(myWindow)
			})
		case "Ciudad":
			usuario.ManejarCiudad(myWindow, func() {
				CrearInterfaz(myWindow)
			})
		}
		dialog.ShowInformation("Selección realizada", fmt.Sprintf("Seleccionaste: %s", ingreso), myWindow)
	})
	myWindow.SetContent(subcategoria)
}
func InterfazPrincipal(myWindow fyne.Window) {
	/*
		principal := widget.NewLabel("Interfaz principal")
		myWindow.SetContent(principal)
	*/

}
