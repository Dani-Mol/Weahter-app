package backend

import (
	"encoding/csv"
	"os"
)

// Función para procesar tickets
func procesarTicket(ticket string) string {
	tickets, err := obtenerTickets()
	if err != nil {
		return "Error al obtener tickets"
	}

	if _, existe := tickets[ticket]; existe {
		return "Ticket válido"
	} else {
		return "Ticket inválido"
	}
}

// Función para procesar ITA
func procesarITA(ita string) string {
	if ValidarEntrada(ita, "ITA esperado") {
		return "ITA válido"
	}
	return "ITA inválido"
}

// Función para procesar ciudad
func procesarCiudad(ciudad string) string {
	if ciudad == "Ciudad esperada" { // Aquí puedes agregar la lógica real
		return "Ciudad válida"
	}
	return "Ciudad inválida"
}

// Función para validar la entrada según el tipo (ticket, ITA o ciudad)
func validarEntrada(tipo, entrada string) string {
	switch tipo {
	case "ticket":
		return procesarTicket(entrada)
	case "ita":
		return procesarITA(entrada)
	case "ciudad":
		return procesarCiudad(entrada)
	default:
		return "Tipo desconocido"
	}
}

// Ruta del dataset definido en el caché
const datasetPath = "dataset/dataset.csv"

// Función que genera un mapa con los tickets y sus IATA asociados
func obtenerTickets() (map[string][]string, error) {
	registros := make(map[string][]string)
	file, err := os.Open(datasetPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, row := range data[1:] {
		registros[row[0]] = []string{row[1], row[2]} // num_ticket, origen, destino
	}

	return registros, nil
}

// Función para validar si una cadena contiene otra (ya existente en tu código)
func ValidarEntrada(cadena string, cadena2 string) bool {
	return trings.Contains(cadena, cadena2)
}
