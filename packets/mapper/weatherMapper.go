package mapper

var weather = map[uint8]string{
	0: "clear",
	1: "light cloud",
	2: "overcast",
	3: "light rain",
	4: "heavy rain",
	5: "storm",
}

func GetMappedWeather(weatherID uint8) string {
	return weather[weatherID]
}
