package main

import (
	"fmt"
	"net/http"
	"renebizelli/go/weather/configs"
	viacep "renebizelli/go/weather/externals/ViaCEP"
	weatherAPI "renebizelli/go/weather/externals/WeatherAPI"
	"renebizelli/go/weather/internals/webserver"
	"time"
)

func main() {

	configs := configs.LoadConfig("./")

	mux := http.NewServeMux()

	timeout := time.Duration(configs.SERVICES_TIMEOUT) * time.Second

	cep := viacep.NewCEPService(mux, configs.VIACEP_URL, timeout)

	weather := weatherAPI.NewWeatherService(mux, configs.WEATHERAPI_URL, configs.WEATHERAPI_KEY, timeout)

	handler := webserver.NewHandler(mux, cep, weather, timeout)
	handler.RegisterRoutes()

	fmt.Printf("Web server running on port %v\n", configs.WEBSERVER_PORT)

	http.ListenAndServe(fmt.Sprintf(":%v", configs.WEBSERVER_PORT), mux)

}
