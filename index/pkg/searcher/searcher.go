package searcher

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/agnivade/levenshtein"
	"golang.org/x/text/unicode/norm"
)

// Ruta del dataset definido en el caché
const datasetPath = "dataset/dataset.csv"

// Función que genera un mapa con los tickets y sus IATA asociados
func ObtenerTickets() (map[string][]string, error) {
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

// Función que obtiene los datos del caché (simulación)
func ObtenerCache() map[string][][]string {
	return map[string][][]string{
		"ABC": {{"Soleado", "30°C", "15°C", "40%"}},
		"XYZ": {{"Lluvioso", "22°C", "18°C", "60%"}},
	}
}

// Función que toma datos del caché y los formatea
func TomarDelCache(tickets map[string][]string, cache map[string][][]string, origen, destino string) ([]string, error) {
	var datos []string
	_ = time.Now().Hour()

	// Si el origen y destino son iguales o no hay coincidencias en caché
	if origen == destino || cache[origen] == nil || cache[destino] == nil {
		return []string{"NULL", "NULL", "NULL", "NULL", "NULL", "NULL"}, nil
	}

	// Agrega datos del origen
	origenData := cache[origen][0]
	datos = append(datos, origen, tickets[origen][0], tickets[origen][1], origenData[0], origenData[1], origenData[3])

	// Agrega datos del destino
	destinoData := cache[destino][0]
	datos = append(datos, destino, tickets[destino][0], tickets[destino][1], destinoData[0], destinoData[1], destinoData[3])

	return datos, nil
}

// Función que busca en el caché utilizando Levenshtein
func BuscarEnCache(texto string) [][]string {
	tickets, err := ObtenerTickets()
	if err != nil {
		fmt.Println("Error obteniendo tickets:", err)
		return nil
	}

	cache := ObtenerCache()
	var resultados [][]string

	// Normaliza el texto de entrada
	textoNormalizado := norm.NFC.String(strings.ToLower(texto))

	// Busca por ticket
	for ticket, info := range tickets {
		if strings.Contains(strings.ToLower(ticket), textoNormalizado) {
			resultado, _ := TomarDelCache(tickets, cache, info[0], info[1])
			resultados = append(resultados, resultado)
		}
	}

	// Si no encuentra por ticket, busca por IATA
	if len(resultados) == 0 {
		for _, info := range tickets {
			origen := strings.ToLower(info[0])
			destino := strings.ToLower(info[1])

			if strings.Contains(origen, textoNormalizado) || strings.Contains(destino, textoNormalizado) {
				resultado, _ := TomarDelCache(tickets, cache, info[0], info[1])
				resultados = append(resultados, resultado)
			}
		}
	}

	// Si no encuentra coincidencias exactas, usa Levenshtein
	if len(resultados) == 0 {
		var mejorCoincidencia string
		minDistancia := -1

		for _, info := range tickets {
			origen := strings.ToLower(info[0])
			distancia := levenshtein.ComputeDistance(textoNormalizado, origen)

			if minDistancia == -1 || distancia < minDistancia {
				mejorCoincidencia = info[0]
				minDistancia = distancia
			}
		}

		if mejorCoincidencia != "" {
			resultado, _ := TomarDelCache(tickets, cache, mejorCoincidencia, tickets[mejorCoincidencia][1])
			resultados = append(resultados, resultado)
		}
	}

	return resultados
}
