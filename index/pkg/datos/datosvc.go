package datos

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func LeerCsv(documento string) map[string]bool {
	archivo, error := os.Open(documento)
	if error != nil {
		fmt.Println("Error al abrir el csv ", error)
	}
	defer archivo.Close()

	lectorArchivo := csv.NewReader(archivo)
	lectorArchivo.FieldsPerRecord = -1

	destinos := make(map[string]bool)
	for {
		fila, incidente := lectorArchivo.Read()
		if incidente == io.EOF {
			break
		}
		if incidente != nil {
			fmt.Println("Error al leer una linea del archivo CSV:", incidente)
		}
		destino := fila[1]
		if !destinos[destino] {
			destinos[destino] = true
		}
	}
	return destinos
}
