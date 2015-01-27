package main
import "net/http"

import (
	"strings"
	"strconv"
	"encoding/json"
	"io/ioutil"
)

type weatherData struct
{
	Name string `json:"name"` 
	Main struct
	{
		Kelvin float64 `json:"temp"`
		Humidity int `json:"humidity"`
	} `json:"main"`

	Wind struct
	{
		Speed float64 `json:"speed"`
		Deg float64 `json:"deg"`
	} `json:"wind"`
} 

func main() {
	http.HandleFunc("/", weather)

	http.ListenAndServe(":8080", nil)
}

func weather(w http.ResponseWriter, r *http.Request) {
	var wdata weatherData
	city := strings.SplitN(r.URL.Path, "/", 3)
	if len(city[2]) == 0 {
		w.Write([]byte("Error, CITY,STATE not inputted.\n"))
		return
	}
	wdata, err := query(city[2])
	if err != nil {
		w.Write([]byte("Error! Probably no internet connection.\n"))
		return
	}
	if len(wdata.Name) == 0 {
		w.Write([]byte("Error, City not found.\n"))
		return
	}
	
	//city name
	w.Write([]byte("City:"))
	w.Write([]byte(wdata.Name))

	//temperature
	var temp float64
	temp = (wdata.Main.Kelvin - 273.15) * 1.8 + 32
	w.Write([]byte(", Temperature:"))
	w.Write([]byte(FloatToString(temp)))

	//humidity
	w.Write([]byte("F, Humidity:"))
	w.Write([]byte(strconv.Itoa(wdata.Main.Humidity)))

	//wind speed
	w.Write([]byte(", Wind Speed:"))
	w.Write([]byte(FloatToString(wdata.Wind.Speed)))

	//wind dir
	w.Write([]byte(", Wind Direction:"))
	w.Write([]byte(FloatToString(wdata.Wind.Deg)))

	w.Write([]byte("\n"));
} 

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 3, 64)
}

func query(city string) (weatherData, error) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q="+ city)
	if err != nil {
		return weatherData{}, err
	}
	//var response []byte
	response, err := ioutil.ReadAll(resp.Body)

	//var data weatherData
	data := weatherData{}
	err = json.Unmarshal(response, &data)

	return data, err
} 
