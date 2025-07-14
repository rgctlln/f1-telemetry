package packets

type PacketLapData struct {
	Header  PacketHeader
	LapData [22]lapData
}

type lapData struct {
}
