package usuarios

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type Entrada interface {
	ManejarTicket(ventana fyne.Window)
	ManejarITA(ventana fyne.Window)
	ManejarCiudad(ventana fyne.Window)
}
type Pasajero struct{}

type Servidor struct{}

func (p *Pasajero) ManejarTicket(myWindow fyne.Window) {
	var flujo = true
	for flujo {
		entrada := widget.NewEntry()
		dialog.ShowCustomConfirm("Ingrese Ticket", "Aceptar", "Cancelar", entrada, func(confirmacion bool) {
			if confirmacion{
				ticket:=entrada.Text
				if{ // Verificacion de que el ticket es valido conforme los datos
					fmt.Println("Ticket valido:", ticket)
					flujo=false
				}else{
					dialog.ShowInformation("Error", "Ticket no valido. Intentelo de nuevo", myWindow)
				}
			}else{
				fmt.Println("Gracias por su visita")
				flujo=false
			}

		}, myWindow)
	}
}

func (p *Pasajero) ManejarITA(myWindow fyne.Window) {
	var flujo = true
	for flujo {
		entrada := widget.NewEntry()
		dialog.ShowCustomConfirm("Ingresa ITA", "Aceptar", "Cancelar", entrada, func(confirmacion bool) {
			if confirmacion{
				ticket:=entrada.Text
				if{ // Verificacion de que el ticket es valido conforme los datos
					fmt.Println("Ticket valido:", ticket)
					flujo=false
				}else{
					dialog.ShowInformation("Error", "Ticket no valido. Intentelo de nuevo", myWindow)
				}
			}else{
				fmt.Println("Gracias por su visita")
				flujo=false
			}

		}, myWindow)
	}
}

func (p *Pasajero) ManejarCiudad(myWindow fyne.Window) {
	var flujo = true
	for flujo {
		entrada := widget.NewEntry()
		dialog.ShowCustomConfirm("Ingresa Ciudad", "Aceptar", "Cancelar", entrada, func(confirmacion bool) {
			if confirmacion{
				ticket:=entrada.Text
				if{ // Verificacion de que el ticket es valido conforme los datos
					fmt.Println("Ticket valido:", ticket)
					flujo=false
				}else{
					dialog.ShowInformation("Error", "Ticket no valido. Intentelo de nuevo", myWindow)
				}
			}else{
				fmt.Println("Gracias por su visita")
				flujo=false
			}

		}, myWindow)
	}

}
func (s *Servidor) ManejarTicket(myWindow fyne.Window) {
	var flujo = true
	for flujo {
		entrada := widget.NewEntry()
		dialog.ShowCustomConfirm("Ingrese Ticket", "Aceptar", "Cancelar", entrada, func(confirmacion bool) {
			if confirmacion{
				ticket:=entrada.Text
				if{ // Verificacion de que el ticket es valido conforme los datos
					fmt.Println("Ticket valido:", ticket)
					flujo=false
				}else{
					dialog.ShowInformation("Error", "Ticket no valido. Intentelo de nuevo", myWindow)
				}
			}else{
				fmt.Println("Gracias por su visita")
				flujo=false
			}

		}, myWindow)
	}

}

func (p *Servidor) ManejarITA(myWindow fyne.Window) {
	var flujo = true
	for flujo {
		entrada := widget.NewEntry()
		dialog.ShowCustomConfirm("Ingrese Ticket", "Aceptar", "Cancelar", entrada, func(confirmacion bool) {
			if confirmacion{
				ticket:=entrada.Text
				if{ // Verificacion de que el ticket es valido conforme los datos
					fmt.Println("Ticket valido:", ticket)
					flujo=false
				}else{
					dialog.ShowInformation("Error", "Ticket no valido. Intentelo de nuevo", myWindow)
				}
			}else{
				fmt.Println("Gracias por su visita")
				flujo=false
			}

		}, myWindow)
	}

}

func (p *Servidor) ManejarCiudad(myWindow fyne.Window) {
	var flujo = true
	for flujo {
		entrada := widget.NewEntry()
		dialog.ShowCustomConfirm("Ingrese Ticket", "Aceptar", "Cancelar", entrada, func(confirmacion bool) {
			if confirmacion{
				ticket:=entrada.Text
				if{ // Verificacion de que el ticket es valido conforme los datos
					fmt.Println("Ticket valido:", ticket)
					flujo=false
				}else{
					dialog.ShowInformation("Error", "Ticket no valido. Intentelo de nuevo", myWindow)
				}
			}else{
				fmt.Println("Gracias por su visita")
				flujo=false
			}

		}, myWindow)
	}

}
