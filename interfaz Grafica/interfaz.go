package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Crear una nueva aplicación
	myApp := app.New()

	// Crear una nueva ventana
	myWindow := myApp.NewWindow("Whater App")

	opciones := []string{"Ticket", "Ciudad", "ITA"}

	selectWidget := widget.NewSelect(opciones, func(valor string) {

	})

	// Mostrar la ventana
	myWindow.ShowAndRun()
}

func seleccionarOpciones(valor string, myWindow fyne.Window) {
	switch valor {
	case "ITA":
		manejarITA(myWindow)
	case "Ticket":
		manejarTicket(myWindow)
	case "Ciudad":
		manejarCiudad()
	}
}

func manejarTicket(myWindow fyne.Window) {
	var flujo = true
	for flujo {
		entrada := widget.NewEntry()
		dialog.ShowCustomConfirm("Ingrese Ticket", "Aceptar", "Cancelar", entrada, func(confirmacion bool) {
			if confirmacion {
				//Aqui se manda a llamar al metodo que da los datos del csv
				//Se pretende hacer un metodo que ingresado la entrada del usuario regrese si el Ita es correcto
				//Si es correcto se manda a llamar un metodo  que despegue las recomendaciones
				encontrado := false //Aqui enetrada va a ser el resultado del metodo anteriormente
				if !encontrado {
					fmt.Println("Ita no encontrado. Ingresa una nueva opcion por favor")
					flujo = true
				} else {
					flujo = false
				}
			} else {
				fmt.Println("Gracias por su visita")
				flujo = false
			}
		}, myWindow)
	}
}

func manejarITA(myWindow fyne.Window) {
	var flujo = true
	for flujo {
		entrada := widget.NewEntry()
		dialog.ShowCustomConfirm("Ingrese ITA", "Aceptar", "Cancelar", entrada, func(confirmacion bool) {
			if confirmacion {
				//Aqui se manda a llamar al metodo que da los datos del csv
				//Se pretende hacer un metodo que ingresado la entrada del usuario regrese si el Ita es correcto
				//Si es correcto se manda a llamar un metodo  que despegue las recomendaciones
				encontrado := false //Aqui enetrada va a ser el resultado del metodo anteriormente
				if !encontrado {
					fmt.Println("Ita no encontrado. Ingresa una nueva opcion por favor")
					flujo = true
				} else {
					flujo = false
				}
			} else {
				fmt.Println("Gracias por su visita")
				flujo = false
			}
		}, myWindow)
	}
}

func manejarCiudad(myWindow fyne.Window) {
	var flujo = true
	for flujo {
		entrada := widget.NewEntry()
		dialog.ShowCustomConfirm("Ingrese Ciudad", "Aceptar", "Cancelar", entrada, func(confirmacion bool) {
			if confirmacion {
				//Aqui se manda a llamar al metodo que da los datos del csv
				//Se pretende hacer un metodo que ingresado la entrada del usuario regrese si el Ita es correcto
				//Si es correcto se manda a llamar un metodo  que despegue las recomendaciones
				encontrado := false //Aqui enetrada va a ser el resultado del metodo anteriormente
				if !encontrado {
					fmt.Println("Ita no encontrado. Ingresa una nueva opcion por favor")
					flujo = true
				} else {
					flujo = false
				}
			} else {
				fmt.Println("Gracias por su visita")
				flujo = false
			}
		}, myWindow)
	}
}
