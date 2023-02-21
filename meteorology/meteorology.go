package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData

func (temperatureUnit TemperatureUnit) String() string {
	units := []string{"°C", "°F"}
	return units[temperatureUnit]
}

func (temperature Temperature) String() string {
	return fmt.Sprintf("%v %v", temperature.degree, temperature.unit)
}

func (speedUnit SpeedUnit) String() string {
	units := []string{"km/h", "mph"}
	return units[speedUnit]
}

func (speed Speed) String() string {
	return fmt.Sprintf("%v %v", speed.magnitude, speed.unit)
}

func (meteorologyData MeteorologyData) String() string {
	return fmt.Sprintf("%v: %v, Wind %v at %v, %v%% Humidity", meteorologyData.location, meteorologyData.temperature,
		meteorologyData.windDirection, meteorologyData.windSpeed, meteorologyData.humidity)
}
