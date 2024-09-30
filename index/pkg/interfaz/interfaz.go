package interfaz

import (
	"Weahter-app/index/pkg/usuarios"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// En esta funcion se establece la interfaz con el usuario pidiendo un metodo de acceso (Ticket, ita o Ciudad)
func CrearInterfaz(myWindow fyne.Window) {
	InterfazPrincipal(myWindow, func() {
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
		})
		myWindow.SetContent(subcategoria)
	})
}

// Este metodo establece la pagina principal de la aplicacion
func InterfazPrincipal(myWindow fyne.Window, onContinue func()) {

	imagen := canvas.NewImageFromFile("/Users/axelgodinez/Documents/Modelado y Programacion/Proyecto 1 - B/Weahter-app/index/pkg/imagenes/Home.png")
	imagen.FillMode = canvas.ImageFillContain

	continuarButton := widget.NewButton("Continuar", func() {
		if onContinue != nil {
			onContinue()
		}
	})

	botonContenedor := container.NewVBox(
		continuarButton,
	)

	contenedor := container.NewBorder(
		nil,
		botonContenedor,
		nil,
		nil,
		imagen,
	)

	myWindow.SetContent(contenedor)
	size := myWindow.Canvas().Size()
	imagen.Resize(size)

}
