/*
El paquete principal implementa un mecanismo de almacenamiento en caché para datos meteorológicos, utilizando un conjunto de datos CSV de códigos IATA y una API meteorológica.
Version 1.0
*/
package cache

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

// Clave API y rutas de archivos para la memoria caché y el conjunto de datos.
const (
	API_KEY    = "5e31a7313683592fc55490ec53637486"
	JSON_CACHE = "data/cache.json"
	DATA_SET   = "data/dataset1.csv"
)

// Mutexes globales para evitar el acceso simultáneo a la memoria caché y al temporizador.
var (
	timerActive sync.Mutex
	cacheActive sync.Mutex
)

// Variable para almacenar la unica instancia de cache
var (
	cacheInstance *CacheData
	once          sync.Once
)

// El clima representa la estructura de datos meteorológicos recuperada de la API.
type Weather struct {
	Climate  string
	TempMin  int
	TempMax  int
	Humidity int
	Hour     int
}

// CacheData contiene la estructura de los datos meteorológicos almacenados en caché.
type CacheData struct {
	Flag    string                     `json:"flag"`
	Records map[string][][]interface{} `json:"records"`
}

// Inicializa el cache una sola vez sino existe una instancia
func GetCacheSingleton(reset bool) *CacheData {
	once.Do(func() {
		// Inicializa la instancia del caché si no ha sido creada
		cacheInstance = GetCache(reset)
	})
	return cacheInstance
}

// getCache lee el caché de un archivo JSON o inicializa uno nuevo si es necesario.
// Si reset es verdadero, borra el caché y comienza de nuevo.
func GetCache(reset bool) *CacheData {
	var cache *CacheData
	if reset {
		// Inicializa un nuevo caché si el restablecimiento es verdadero
		cache = &CacheData{Flag: time.Now().Format("2006-01-02 15:04:05"), Records: make(map[string][][]interface{})}
		SaveCache(cache)
	} else {
		// Intenta leer desde el archivo de caché existente
		data, err := ioutil.ReadFile(JSON_CACHE)
		if err != nil {
			// Si falta el archivo, inicializar un nuevo caché
			cache = &CacheData{Flag: time.Now().Format("2006-01-02 15:04:05"), Records: make(map[string][][]interface{})}
			SaveCache(cache)
		} else {
			// Desagrupar datos JSON en la estructura de caché
			err := json.Unmarshal(data, &cache)
			if err != nil {
				// Si los datos están dañados, reinicie la caché
				cache = &CacheData{Flag: time.Now().Format("2006-01-02 15:04:05"), Records: make(map[string][][]interface{})}
				SaveCache(cache)
			} else {
				// Comprueba si el caché ha expirado
				CheckCache(cache)
			}
		}
	}
	return cache
}

// checkCache verifica si el caché tiene más de 3 horas de antigüedad.
// Si el caché ha expirado, activa una actualización.
func CheckCache(cache *CacheData) {
	cacheTime, _ := time.Parse("2006-01-02 15:04:05", cache.Flag)
	if time.Since(cacheTime) > 3*time.Hour {
		UpdateCache()
	}
}

// saveCache escribe los datos de caché en el archivo JSON.
func SaveCache(cache *CacheData) {
	data, _ := json.Marshal(cache)
	_ = ioutil.WriteFile(JSON_CACHE, data, 0644)
}

// iataRegistration lee los códigos IATA y sus coordenadas (latitud, longitud) de un conjunto de datos CSV.
// Devuelve un mapa con los códigos IATA como claves y sus coordenadas como valores.
func IataRegistration() map[string][2]float64 {
	iatas := make(map[string][2]float64)
	file, err := os.Open(DATA_SET)
	if err != nil {
		fmt.Println("Error reading dataset:", err)
		return iatas
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	for _, row := range records[1:] {
		originLat, _ := strconv.ParseFloat(row[2], 64)
		originLon, _ := strconv.ParseFloat(row[3], 64)
		destLat, _ := strconv.ParseFloat(row[4], 64)
		destLon, _ := strconv.ParseFloat(row[5], 64)

		// Registrar códigos IATA de origen y destino y sus coordenadas
		if _, exists := iatas[row[0]]; !exists {
			iatas[row[0]] = [2]float64{originLat, originLon}
		}
		if _, exists := iatas[row[1]]; !exists {
			iatas[row[1]] = [2]float64{destLat, destLon}
		}
	}
	return iatas
}

// getWeather realiza una solicitud a una API meteorológica utilizando latitud y longitud y devuelve una matriz de estructuras meteorológicas.
// Simula la solicitud de API.
func (cache CacheData) GetWeather(lat, lon float64) []Weather {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?lat=%f&lon=%f&appid=%s", lat, lon, API_KEY)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error calling API:", err)
		return nil
	}
	defer resp.Body.Close()

	var weatherData []Weather
	// Simulating data retrieval for now.
	return weatherData
}

// updateCache borra el caché actual y lo vuelve a llenar recuperando datos meteorológicos para códigos IATA.
// Hace una pausa entre solicitudes para evitar exceder los límites de velocidad de la API.
func UpdateCache() {
	cacheActive.Lock()
	defer cacheActive.Unlock()

	iatas := IataRegistration()
	cache := GetCacheSingleton(true)

	for iata, coords := range iatas {
		weatherData := cache.GetWeather(coords[0], coords[1])
		cache.Records[iata] = [][]interface{}{}

		// Append weather data to cache
		for _, weather := range weatherData {
			cache.Records[iata] = append(cache.Records[iata], []interface{}{weather.Climate, weather.TempMin, weather.TempMax, weather.Humidity, weather.Hour})
		}
		time.Sleep(1 * time.Second) // Throttle requests to avoid exceeding API rate limits
	}

	SaveCache(cache)
	StartTimer()
}

// startTimer inicia un temporizador en segundo plano que activa actualizaciones de caché cada 3 horas.
func StartTimer() {
	go func() {
		timerActive.Lock()
		defer timerActive.Unlock()
		for {
			time.Sleep(3 * time.Hour)
			UpdateCache()
		}
	}()
}

// runCache inicializa el proceso de caché cuando se inicia el programa.
// Si la caché está vacía, activa una actualización. De lo contrario, inicia el temporizador de actualización periódica.
func RunCache() {
	timerActive.Lock()
	defer timerActive.Unlock()

	cache := GetCacheSingleton(false)
	if len(cache.Records) == 0 {
		UpdateCache()
	} else {
		StartTimer()
	}
}

// main es el punto de entrada del programa. Inicia el proceso de caché.
func main() {
	RunCache()
}
