package usuarios

import (
	"Weahter-app/index/pkg/clima"
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type Usuario struct{}

// Metodo de usuario para establcer como manejar y valiar el Ticket conforme la interfaz grafica
/*
func (u Usuario) ManejarTicket(myWindow fyne.Window, onCancel func()) {

	var mostrarTicket func()

	mostrarTicket = func() {
		entrada := widget.NewEntry()
		dialogo := dialog.NewCustomConfirm("Ingrese Ticket", "Aceptar", "Cancelar", entrada, func(confirmacion bool) {
			if confirmacion {
				ita := entrada.Text
				climaticket, error := clima.ObtenerClimaTicket(ita)
				if error != nil {
					dialog.ShowError(fmt.Errorf("ticket no válido"), myWindow)
					mostrarTicket()
					return
				}
				fmt.Printf("Ticket: %s\n", climaticket.Ticket)
			} else {
				dialog.ShowInformation("Gracias", "Gracias por su visita", myWindow)
				return
			}
			if onCancel != nil {
				onCancel()
			}
		}, myWindow)

		dialogo.Show()
	}
	mostrarTicket()


}*/

// // Metodo de usuario para establcer como manejar y valiar el Ita conforme la interfaz grafica
func (u Usuario) ManejarITA(myWindow fyne.Window, onCancel func()) {
	var mostrarIta func()

	mostrarIta = func() {
		entrada := widget.NewEntry()
		dialogo := dialog.NewCustomConfirm("Ingrese Ita", "Aceptar", "Cancelar", entrada, func(confirmacion bool) {
			if confirmacion {
				ita := entrada.Text
				climaciudad, err := clima.ObtenerClimaIta(ita)
				if err != nil {
					//dialog.ShowError(err, myWindow) // Muestra el error
					mostrarIta() // Muestra el diálogo nuevamente para reingresar el ticket
				} else {
					// Ticket válido, manejar aquí la lógica necesaria
					fmt.Printf("Ticket válido: %s\n", climaciudad.Ciudad)
					if onCancel != nil {
						onCancel()
					}
				}
			} else {
				if onCancel != nil {
					onCancel()
				}
			}
		}, myWindow)

		dialogo.Show()
		// No se establece foco automáticamente
	}

	// Botón para abrir el diálogo
	abrirDialogo := widget.NewButton("Abrir Diálogo de Ticket", func() {
		mostrarIta()
	})

	myWindow.SetContent(container.NewVBox(abrirDialogo))
}

// Metodo de usuario para establcer como manejar y valiar la Ciudad conforme la interfaz grafica
func (U Usuario) ManejarCiudad(myWindow fyne.Window, onCancel func()) {
	var mostrarCiudad func()

	mostrarCiudad = func() {
		entrada := widget.NewEntry()
		dialogo := dialog.NewCustomConfirm("Ingrese Ciudad", "Aceptar", "Cancelar", entrada, func(confirmacion bool) {
			if confirmacion {
				ita := entrada.Text
				climaciudad, err := clima.ObtenerClimaCiudad(ita)
				if err != nil {
					//dialog.ShowError(err, myWindow) // Muestra el error
					mostrarCiudad()
				} else {
					// Ticket válido
					fmt.Printf("Ciudad válido: %s\n", climaciudad.Ciudad)
					if onCancel != nil {
						onCancel()
					}
				}
			} else {
				if onCancel != nil {
					onCancel()
				}
			}
		}, myWindow)

		dialogo.Show()

	}

	// Botón para abrir el diálogo
	abrirDialogo := widget.NewButton("Abrir Diálogo de Ciudad", func() {
		mostrarCiudad()
	})

	myWindow.SetContent(container.NewVBox(abrirDialogo))

}
func ValidarEntrada(cadena string, cadena2 string) bool {
	return strings.Contains(cadena, cadena2)
}
func (u Usuario) ManejarTicket(myWindow fyne.Window, onCancel func()) {
	var mostrarTicket func()

	mostrarTicket = func() {
		entrada := widget.NewEntry()
		dialogo := dialog.NewCustomConfirm("Ingrese Ticket", "Aceptar", "Cancelar", entrada, func(confirmacion bool) {
			if confirmacion {
				ita := entrada.Text
				climaticket, err := clima.ObtenerClimaTicket(ita)
				if err != nil {
					//dialog.ShowError(err, myWindow) // Muestra el error
					mostrarTicket()
				} else {
					// Ticket válido
					fmt.Printf("Ticket válido: %s\n", climaticket.Destino)
					if onCancel != nil {
						onCancel()
					}
				}
			} else {
				if onCancel != nil {
					onCancel()
				}
			}
		}, myWindow)

		dialogo.Show()
	}

	// Botón para abrir el diálogo
	abrirDialogo := widget.NewButton("Abrir Diálogo de Ticket", func() {
		mostrarTicket()
	})

	myWindow.SetContent(container.NewVBox(abrirDialogo))
}

/*
Establecer el retiorno de la estructura de clima
Hacer dos metodos para retornar el clima uno para ticket y otra de ciuidad y ita
Revisar que no se regresen resultados en nil en el segundo valor de la variable para segurar que todo esta bien


Avisarle a Daniel lo del cache si funciona la llamada a la api, y preguntar si esta bien
lo de las funciones de clima

Avisarle a la Laura que revise las funciones adecuadas
*/
