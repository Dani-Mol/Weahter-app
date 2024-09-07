package usuarios

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type Usuario struct{}

// Metodo de usuario para establcer como manejar y valiar el Ticket conforme la interfaz grafica
func (u Usuario) ManejarTicket(myWindow fyne.Window, onCancel func()) {
	openDialogButton := widget.NewButton("Abrir Diálogo", func() {
		entrada := widget.NewEntry()

		dialog.ShowCustomConfirm("Ingrese ITA", "Aceptar", "Cancelar", entrada, func(confirmacion bool) {
			if confirmacion {
				ita := entrada.Text
				fmt.Printf("ITA ingresado: %s\n", ita)
				dialog.ShowInformation("Ticket válido", fmt.Sprintf("Ita válido: %s", ita), myWindow)
				if onCancel != nil {
					onCancel()
				}
			} else {
				dialog.ShowInformation("Gracias", "Gracias por su visita", myWindow)
				if onCancel != nil {
					onCancel()
				}
			}
		}, myWindow)
	})
	myWindow.SetContent(openDialogButton)
}

// // Metodo de usuario para establcer como manejar y valiar el Ita conforme la interfaz grafica
func (u Usuario) ManejarITA(myWindow fyne.Window, onCancel func()) {
	openDialogButton := widget.NewButton("Abrir Diálogo", func() {
		entrada := widget.NewEntry()

		dialog.ShowCustomConfirm("Ingrese ITA", "Aceptar", "Cancelar", entrada, func(confirmacion bool) {
			if confirmacion {
				ita := entrada.Text
				fmt.Printf("ITA ingresado: %s\n", ita)
				dialog.ShowInformation("Ticket válido", fmt.Sprintf("Ita válido: %s", ita), myWindow)
				if onCancel != nil {
					onCancel()
				}
			} else {
				dialog.ShowInformation("Gracias", "Gracias por su visita", myWindow)
				if onCancel != nil {
					onCancel()
				}
			}
		}, myWindow)
	})
	myWindow.SetContent(openDialogButton)
}

// Metodo de usuario para establcer como manejar y valiar la Ciudad conforme la interfaz grafica
func (U Usuario) ManejarCiudad(myWindow fyne.Window, onCancel func()) {
	openDialogButton := widget.NewButton("Abrir Diálogo", func() {
		entrada := widget.NewEntry()

		dialog.ShowCustomConfirm("Ingrese Ciudad", "Aceptar", "Cancelar", entrada, func(confirmacion bool) {
			if confirmacion {
				ciudad := entrada.Text
				fmt.Printf("Ciudad ingresado: %s\n", ciudad)
				dialog.ShowInformation("Ticket válido", fmt.Sprintf("Ciudad válido: %s", ciudad), myWindow)
				if onCancel != nil {
					onCancel()
				}
			} else {
				dialog.ShowInformation("Gracias", "Gracias por su visita", myWindow)
				if onCancel != nil {
					onCancel()
				}
			}
		}, myWindow)
	})
	myWindow.SetContent(openDialogButton)
}
func ValidarEntrada(cadena string, cadena2 string) bool {
	return strings.Contains(cadena, cadena2)
}
