package usuarios

import "fyne.io/fyne/v2"

// Se establece una interfaz con los metodos de tanto usuario como servidor
type Entrada interface {
	ManejarTicket(myWindow fyne.Window)
	ManejarITA(myWindow fyne.Window)
	ManejarCiudad(myWindow fyne.Window)
}
