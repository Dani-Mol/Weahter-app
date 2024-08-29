package usuarios

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type Pasajero struct{}

// Metodo de pasajero para establcer como manejar y valiar el Ticket
func (p *Pasajero) ManejarTicket(myWindow fyne.Window) {
	openDialogButton := widget.NewButton("Abrir Diálogo", func() {
		entrada := widget.NewEntry()

		dialog.ShowCustomConfirm("Ingrese Ticket", "Aceptar", "Cancelar", entrada, func(confirmacion bool) {
			if confirmacion {
				ticket := entrada.Text
				fmt.Printf("Ticket ingresado: %s\n", ticket)
				dialog.ShowInformation("Ticket válido", fmt.Sprintf("Ticket válido: %s", ticket), myWindow)
			} else {
				dialog.ShowInformation("Gracias", "Gracias por su visita", myWindow)
			}
		}, myWindow)
	})
	myWindow.SetContent(openDialogButton)
}

func (p *Pasajero) ManejarITA(myWindow fyne.Window) {
	// Implementar la lógica para manejar el ITA
}

func (p *Pasajero) ManejarCiudad(myWindow fyne.Window) {
	// Implementar la lógica para manejar la ciudad
}
