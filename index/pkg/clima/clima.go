package clima

import (
	"Weahter-app/index/pkg/cache"
	"Weahter-app/index/pkg/dataset"
	"fmt"
)

type TicketClima struct {
	Ticket             string
	Origen             string
	Destino            string
	CoordenadasOrigen  [2]float64
	CoordenadasDestino [2]float64
	ClimaOrigen        Clima
	ClimaDestino       Clima
}

type CiudadClima struct {
	Ciudad string
	Clima  Clima
}

type Clima struct {
	Climate  string
	TempMin  int
	TempMax  int
	Humidity int
	Hour     int
}

func ObtenerClimaTicket(ticket string) (*TicketClima, error) {

	//Inicializacion de cache
	cacheprincipal := cache.GetCacheSingleton(false)

	data := dataset.NewDatasetManager()
	//Obtener los codigos Ita del Ticket
	iata, confirmacion := data.GetIATAs(ticket)
	if !confirmacion {
		return nil, fmt.Errorf("el ticket no existe")
	}

	//Obtener las coordenadas de los Ita de origen y destino
	cordsOrigen, origenExiste := data.GetCoords(iata[0])
	cordsDestino, destinoExiste := data.GetCoords(iata[1])
	if !origenExiste || !destinoExiste {
		return nil, fmt.Errorf("no se encontraron coordenadas para el ticket")
	}
	//Obtener el clima de Origen y de destino de la ciudad
	climaOrigen := cacheprincipal.GetWeather(cordsOrigen[0], cordsOrigen[1])
	climaDestino := cacheprincipal.GetWeather(cordsDestino[0], cordsDestino[1])

	info := &TicketClima{
		Ticket:             ticket,
		Origen:             iata[0],
		Destino:            iata[1],
		CoordenadasOrigen:  [2]float64{cordsOrigen[0], cordsOrigen[1]},
		CoordenadasDestino: [2]float64{cordsDestino[0], cordsDestino[1]},
	}

	// Asignar clima del origen
	if len(climaOrigen) > 0 {
		info.ClimaOrigen = Clima{
			Climate:  climaOrigen[0].Climate,
			TempMin:  climaOrigen[0].TempMin,
			TempMax:  climaOrigen[0].TempMax,
			Humidity: climaOrigen[0].Humidity,
			Hour:     climaOrigen[0].Hour,
		}
	}

	// Asignar clima del destino
	if len(climaDestino) > 0 {
		info.ClimaDestino = Clima{
			Climate:  climaDestino[0].Climate,
			TempMin:  climaDestino[0].TempMin,
			TempMax:  climaDestino[0].TempMax,
			Humidity: climaDestino[0].Humidity,
			Hour:     climaDestino[0].Hour,
		}
	}

	return info, nil
}
func ObtenerClimaCiudad(ciudad string) (*CiudadClima, error) {

	cacheprincipal := cache.GetCacheSingleton(false)

	data := dataset.NewDatasetManager()
	//Obtener los codigos Ita del Ticket
	iata := data.GetIATA(ciudad)

	//Obtener las coordenadas de los Ita de origen y destino
	cords, exite := data.GetCoords(iata)
	if !exite {
		return nil, fmt.Errorf("no se encontraron coordenadas para la ciudad")
	}

	clima := cacheprincipal.GetWeather(cords[0], cords[1])

	info := &CiudadClima{
		Ciudad: ciudad,
	}

	if len(clima) > 0 {
		info.Clima = Clima{
			Climate:  clima[0].Climate,
			TempMin:  clima[0].TempMin,
			TempMax:  clima[0].TempMax,
			Humidity: clima[0].Humidity,
			Hour:     clima[0].Hour,
		}
	}

	return info, nil
}

func ObtenerClimaIta(ita string) (*CiudadClima, error) {

	//Cache
	cacheprincipal := cache.GetCacheSingleton(false)

	data := dataset.NewDatasetManager()
	//Obtener los codigos Ita del Ticket

	//Obtener las coordenadas de los Ita de origen y destino
	exite := data.IsValidIATA(ita)
	if !exite {
		return nil, fmt.Errorf("no se encontraron coordenadas para la ita ")
	}
	cords, existe := data.GetCoords(ita)
	if !existe {
		return nil, fmt.Errorf("no se encontraron coordenadas para la ciudad")
	}

	clima := cacheprincipal.GetWeather(cords[0], cords[1])

	info := &CiudadClima{
		Ciudad: ita,
	}

	if len(clima) > 0 {
		info.Clima = Clima{
			Climate:  clima[0].Climate,
			TempMin:  clima[0].TempMin,
			TempMax:  clima[0].TempMax,
			Humidity: clima[0].Humidity,
			Hour:     clima[0].Hour,
		}
	}

	return info, nil
}
