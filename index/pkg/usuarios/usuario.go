package usuarios

import (
	"Weahter-app/index/pkg/clima"
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Usuario struct{}

// Metodo de usuario para establcer como manejar y valiar el Ita conforme la interfaz grafica
func (u Usuario) ManejarITA(myWindow fyne.Window, onCancel func()) {
	var mostrarIta func()

	mostrarIta = func() {
		entrada := widget.NewEntry()
		dialogo := dialog.NewCustomConfirm("Ingrese Ita", "Aceptar", "Cancelar", entrada, func(confirmacion bool) {
			if confirmacion {
				ita := entrada.Text
				climaciudad, err := clima.ObtenerClimaIta(ita)
				if err != nil {
					mostrarIta() // Muestra el diálogo nuevamente para reingresar la Ita
				} else {
					// Ita válido
					crearInterfazClima(myWindow, climaciudad)
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
	abrirDialogo := widget.NewButton("Abrir Diálogo de Ita", func() {
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
					mostrarCiudad()
				} else {
					// Ciudad válido
					crearInterfazClima(myWindow, climaciudad)
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
					mostrarTicket()
				} else {
					// Ticket válido
					crearInterfazTicketClima(myWindow, climaticket)
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
func crearInterfazTicketClima(myWindow fyne.Window, ticketClima *clima.TicketClima) {
	titulos := canvas.NewText("Detalles del Ticket", theme.Color(theme.ColorNameForeground))
	titulos.TextStyle = fyne.TextStyle{Bold: true}
	titulos.TextSize = 24
	ticketCard := widget.NewCard(
		"Ticket",
		ticketClima.Ticket,
		widget.NewLabel("Informacion del ticket y destinos"),
	)

	//Clima de Origen
	origenCard := widget.NewCard(
		"Clima en "+ticketClima.Origen,
		"",
		container.NewVBox(
			widget.NewLabel(fmt.Sprintf("Coordenadas Origen: [%.4f, %.4f]", ticketClima.CoordenadasOrigen[0], ticketClima.CoordenadasOrigen[1])),
			widget.NewLabel("Clima: "+ticketClima.ClimaOrigen.Climate),
			widget.NewLabel(fmt.Sprintf("Temperatura: %d - %d°C", ticketClima.ClimaOrigen.TempMin, ticketClima.ClimaOrigen.TempMax)),
			widget.NewLabel(fmt.Sprintf("Humedad: %d%%", ticketClima.ClimaOrigen.Humidity)),
			widget.NewLabel(fmt.Sprintf("Hora: %d:00", ticketClima.ClimaOrigen.Hour)),
		),
	)
	//Clima de Destino
	destinoCard := widget.NewCard(
		"Clima en "+ticketClima.Destino,
		"",
		container.NewVBox(
			widget.NewLabel(fmt.Sprintf("Coordenadas Destino: [%.4f, %.4f]", ticketClima.CoordenadasDestino[0], ticketClima.CoordenadasDestino[1])),
			widget.NewLabel("Clima: "+ticketClima.ClimaDestino.Climate),
			widget.NewLabel(fmt.Sprintf("Temperatura: %d - %d°C", ticketClima.ClimaDestino.TempMin, ticketClima.ClimaDestino.TempMax)),
			widget.NewLabel(fmt.Sprintf("Humedad: %d%%", ticketClima.ClimaDestino.Humidity)),
			widget.NewLabel(fmt.Sprintf("Hora: %d:00", ticketClima.ClimaDestino.Hour)),
		),
	)

	// Contenedor principal con espaciado adecuado entre elementos
	content := container.NewVBox(
		titulos,
		ticketCard,
		widget.NewSeparator(),
		origenCard,
		widget.NewSeparator(),
		destinoCard,
	)
	myWindow.SetContent(content)

}
func crearInterfazClima(myWindow fyne.Window, clima *clima.CiudadClima) {
	titulos := canvas.NewText("Detalles del Clima", theme.Color(theme.ColorNameForeground))
	titulos.TextStyle = fyne.TextStyle{Bold: true}
	titulos.TextSize = 24
	ciudadCard := widget.NewCard(
		"Ciudad y/o Ita",
		clima.Ciudad,
		widget.NewLabel("Informacion de la Ciudad"),
	)

	//Clima
	climaCard := widget.NewCard(
		"Clima en "+clima.Ciudad,
		"",
		container.NewVBox(
			widget.NewLabel("Clima: "+clima.Clima.Climate),
			widget.NewLabel(fmt.Sprintf("Temperatura: %d - %d°C", clima.Clima.TempMin, clima.Clima.TempMax)),
			widget.NewLabel(fmt.Sprintf("Humedad: %d%%", clima.Clima.Humidity)),
			widget.NewLabel(fmt.Sprintf("Hora: %d:00", clima.Clima.Hour)),
		),
	)

	// Contenedor principal con espaciado adecuado entre elementos
	content := container.NewVBox(
		titulos,
		ciudadCard,
		widget.NewSeparator(),
		climaCard,
		widget.NewSeparator(),
	)
	myWindow.SetContent(content)

}

/*
Establecer el retorno de la estructura de clima
Hacer dos metodos para retornar el clima uno para ticket y otra de ciuidad y ita
Revisar que no se regresen resultados en nil en el segundo valor de la variable para segurar que todo esta bien


Avisarle a Daniel lo del cache si funciona la llamada a la api, y preguntar si esta bien
lo de las funciones de clima

Avisarle a la Laura que revise las funciones adecuadas
*/
