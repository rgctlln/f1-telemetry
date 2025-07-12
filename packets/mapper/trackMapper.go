package mapper

var TrackNames = map[int8]string{
	0:  "Melbourne, Australia",
	1:  "Paul Ricard, France",
	2:  "Shanghai, China",
	3:  "Sakhir, Bahrain",
	4:  "Catalunya, Spain",
	5:  "Monaco, Monaco",
	6:  "Montreal, Canada",
	7:  "Silverstone, United Kingdom",
	8:  "Hockenheim, Germany",
	9:  "Hungaroring, Hungary",
	10: "Spa, Belgium",
	11: "Monza, Italy",
	12: "Singapore, Singapore",
	13: "Suzuka, Japan",
	14: "Abu Dhabi, United Arab Emirates",
	15: "Texas, United States",
	16: "Brazil, Brazil",
	17: "Austria, Austria",
	18: "Sochi, Russia",
	19: "Mexico, Mexico",
	20: "Baku, Azerbaijan",
	21: "Sakhir Short, Bahrain",
	22: "Silverstone Short, United Kingdom",
	23: "Texas Short, United States",
	24: "Suzuka Short, Japan",
	25: "Hanoi, Vietnam",
	26: "Zandvoort, Netherlands",
	27: "Imola, Italy",
	28: "Portimão, Portugal",
	29: "Jeddah, Saudi Arabia",
	30: "Miami, United States",
	31: "Las Vegas, United States",
	32: "Losail, Qatar",
	-1: "Unknown",
}

func GetMappedTrack(trackID int8) string {
	return TrackNames[trackID]
}
