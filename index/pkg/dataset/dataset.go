/*
Este archivo contiene la implementación de la clase DatasetManager, la cual se encarga de leer los datos del dataset y de la base de datos de códigos IATA.
Version 1.0
*/

package dataset

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type DatasetManager struct {
	Coords     map[string][2]float64
	Iatas      map[string][2]string
	Names      map[string]string
	NamesList  []string
	ValidNames []string
}

// Constructor para inicializar la clase DatasetManager
func NewDatasetManager() *DatasetManager {
	dm := &DatasetManager{
		Coords:    make(map[string][2]float64),
		Iatas:     make(map[string][2]string),
		Names:     make(map[string]string),
		NamesList: []string{},
	}
	dm.readTickets()
	dm.readNames()
	dm.filterValidNames()
	return dm
}

// Lectura de los tickets del dataset
func (dm *DatasetManager) readTickets() {
	file, err := os.Open("Resources/dataset2.csv")
	if err != nil {
		fmt.Println("Error al abrir dataset:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	for i, line := range records {
		if i == 0 {
			continue
		}
		ticketNumber := line[0]
		iataOrig := line[1]
		iataDest := line[2]

		dm.Iatas[ticketNumber] = [2]string{iataOrig, iataDest}

		if _, exists := dm.Coords[iataOrig]; !exists {
			latOrig, _ := strconv.ParseFloat(line[3], 64)
			longOrig, _ := strconv.ParseFloat(line[4], 64)
			dm.Coords[iataOrig] = [2]float64{latOrig, longOrig}
		}

		if _, exists := dm.Coords[iataDest]; !exists {
			latDest, _ := strconv.ParseFloat(line[5], 64)
			longDest, _ := strconv.ParseFloat(line[6], 64)
			dm.Coords[iataDest] = [2]float64{latDest, longDest}
		}
	}
}

// Lectura de la base de datos de códigos IATA
func (dm *DatasetManager) readNames() {
	file, err := os.Open("Resources/name-iata.csv")
	if err != nil {
		fmt.Println("Error al abrir name-iata:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	for i, line := range records {
		if i == 0 {
			continue
		}
		dm.Names[line[0]] = line[2]
		dm.Names[line[1]] = line[2]
		dm.Names[line[2]] = line[2]

		dm.NamesList = append(dm.NamesList, line[0], line[1], line[2])
	}
}

// Filtra los nombres válidos
func (dm *DatasetManager) filterValidNames() {
	for _, name := range dm.NamesList {
		if dm.IsValidIATA(name) || dm.IsValidIATA(dm.Names[name]) {
			dm.ValidNames = append(dm.ValidNames, name)
		}
	}
}

// Obtención de coordenadas por código IATA
func (dm *DatasetManager) GetCoords(iata string) (coords [2]float64, exists bool) {
	coords, exists = dm.Coords[iata]
	return
}

// Obtención de códigos IATA por ticket
func (dm *DatasetManager) GetIATAs(ticket string) (iatas [2]string, exists bool) {
	iatas, exists = dm.Iatas[ticket]
	return
}

// Verificación de si un código IATA es válido
func (dm *DatasetManager) IsValidIATA(iata string) bool {
	_, exists := dm.Coords[iata]
	return exists
}

// Obtención de la lista de nombres
func (dm *DatasetManager) GetNamesList() []string {
	return dm.NamesList
}

// Obtención del código IATA por nombre
func (dm *DatasetManager) GetIATA(name string) (iata string) {
	return dm.Names[name]
}

// Obtención de la lista de códigos IATA válidos
func (dm *DatasetManager) GetValidNamesList() []string {
	return dm.ValidNames
}

// Ejemplo de uso
func main() {
	datasetManager := NewDatasetManager()
	fmt.Println("Lista de nombres:", datasetManager.GetNamesList())
	coords, _ := datasetManager.GetCoords("JFK")
	fmt.Println("Coordenadas para un IATA específico:", coords)
}
