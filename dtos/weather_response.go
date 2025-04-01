package dtos

type WeatherResponse struct {
	HttpStatus int     `json:"httpStatus"`
	Celsius    float32 `json:"temp_C"`
	Fahrenheit float32 `json:"temp_F"`
	Kelvin     float32 `json:"temp_K"`
}
