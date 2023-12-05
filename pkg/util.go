package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetData(city string) {
	config, err := LoadConfig()
	if err != nil {
		panic("Problem in unmarshal of env file")
	}
	api_key := config.WeatherCli_API_KEY
	coor_URL := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%v&limit=1&appid=%v", city, api_key)
	res, err := http.Get(coor_URL)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not found")
	}
	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic("Unable to read the body of the API")
	}

	str := string(body)
	str = str[1 : len(str)-1]
	body = []byte(str)

	var coor Coor
	json.Unmarshal(body, &coor)

	// fmt.Println(coor.Lat)
	URL := fmt.Sprintf("http://api.openweathermap.org/data/2.5/forecast?lat=%v&lon=%v&appid=%v", coor.Lat, coor.Lon, api_key)
	// fmt.Println(URL)

	response, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		panic("Weather API not found")
	}

	body, err = io.ReadAll(response.Body)

	if err != nil {
		panic("Unable to read the body of the API")
	}

	json.Unmarshal(body, &data)

	// fmt.Println(string(hello))
}

func GetTempInfo(city string) {
	GetData(city)
	for i, list := range data.List {
		if i > 9 {
			break
		}
		timeValue := time.Unix(list.DataTime, 0)
		formattedTime := timeValue.Format("02-01-2023 15:04:05")
		fmt.Println(formattedTime)
		fmt.Printf("Temperature: %.2fC, Feels Like: %.2fC, Humidity: %v%% \n", list.Main.Temp-273.15, list.Main.FeelsLike-273.15, list.Main.Humidity)
	}
}

func GetWindInfo(city string) {
	GetData(city)
	for i, list := range data.List {
		if i > 9 {
			break
		}
		timeValue := time.Unix(list.DataTime, 0)
		formattedTime := timeValue.Format("02-01-2023 15:04:05")
		fmt.Println(formattedTime)
		fmt.Printf("Speed: %v m/s, Gust: %v m/s, Degree: %v \n", list.Wind.Speed, list.Wind.Gust, list.Wind.Deg)
	}
}

func GetPressureInfo(city string) {
	GetData(city)
	for i, list := range data.List {
		if i > 9 {
			break
		}
		timeValue := time.Unix(list.DataTime, 0)
		formattedTime := timeValue.Format("02-01-2023 15:04:05")
		fmt.Println(formattedTime)
		fmt.Printf("Atmospneric Pressure(Sea Level): %v hPa, Atmospheric Pressure(Ground Level): %v hPa\n", list.Main.SeaLevel, list.Main.GroundLevel)
	}
}

func GetVisibilityInfo(city string) {
	GetData(city)
	for i, list := range data.List {
		if i > 9 {
			break
		}
		timeValue := time.Unix(list.DataTime, 0)
		formattedTime := timeValue.Format("02-01-2023 15:04:05")
		fmt.Println(formattedTime)
		fmt.Printf("Visibility: %v metres\n", list.Visibility)
	}
}

func GetPrecipitationInfo(city string) {
	GetData(city)
	for i, list := range data.List {
		if i > 9 {
			break
		}
		timeValue := time.Unix(list.DataTime, 0)
		formattedTime := timeValue.Format("02-01-2023 15:04:05")
		fmt.Println(formattedTime)
		fmt.Printf("Probability Of Precipitaion: %v %% metres\n", list.Pop)
	}
}
