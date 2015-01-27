## go-weather
A simple attempt at programming with Go using the openweathermap API. This is a daemon running in the background that listens to port 8080. When sending the correctly structured HTTP request, weather information about the CITY,STATE you requested about is returned. 

Weather information retrieved include: Temperature, humidity, wind speed, and wind direction.

## Building the binary
run "go build"

## Running
1. Make sure you have a solid internet connection so that the weather API can be called successfully.
2. run "./weather" to start the webserver on port 8080 of localhost.
3. In a terminal, run "curl localhost:8080/weather/CITY,STATE"
	
	Example: curl localhost:8080/weather/piscataway,nj

## How data is searched
Do a request against localhost:8080/weather/CITY,STATE where CITY is a city and STATE is the state in which CITY is in. The CITY,STATE is extracted from the url and is used as the argument to search for in the weather API. The API does most of the work to retrieve the data, and I simply return it.
