package packets

import (
	"encoding/binary"
	"math"
)

// TODO add appendix
type PacketSessionData struct {
	Header  PacketHeader
	Weather uint8 // Weather - 0 = clear, 1 = light cloud, 2 = overcast
	// 3 = light rain, 4 = heavy rain, 5 = storm

	TrackTemperature int8
	AirTemperature   int8
	TotalLaps        uint8
	TrackLength      uint16
	SessionType      uint8
	TrackId          int8
	Formula          uint8 // Formula, 0 = F1 Modern, 1 = F1 Classic, 2 = F2,
	// 3 = F1 Generic, 4 = Beta, 6 = Esports
	// 8 = F1 World, 9 = F1 Elimination

	SessionTimeLeft   uint16
	SessionDuration   uint16
	PitSpeedLimit     uint8
	GamePaused        uint8
	IsSpectating      uint8
	SpectatorCarIndex uint8
	Sector1Length     float32
	Sector2Length     float32
	Sector3Length     float32
}

// Size: 753 bytes
// Frequency: 2 per second
func ParseSessionPacket(data []byte) PacketSessionData {
	header := ParseHeader(data)

	tail := data[len(data)-8:]
	sector2Start := math.Float32frombits(binary.LittleEndian.Uint32(tail[:4]))
	sector3Start := math.Float32frombits(binary.LittleEndian.Uint32(tail[4:]))

	trackLength := binary.LittleEndian.Uint16(data[33:35])

	return PacketSessionData{
		Header:            header,
		Weather:           data[29],
		TrackTemperature:  int8(data[30]),
		AirTemperature:    int8(data[31]),
		TotalLaps:         data[32],
		TrackLength:       trackLength,
		SessionType:       data[35],
		TrackId:           int8(data[36]),
		Formula:           data[37],
		SessionTimeLeft:   binary.LittleEndian.Uint16(data[38:40]),
		SessionDuration:   binary.LittleEndian.Uint16(data[40:42]),
		PitSpeedLimit:     data[42],
		GamePaused:        data[43],
		IsSpectating:      data[44],
		SpectatorCarIndex: data[45],
		Sector1Length:     sector2Start,
		Sector2Length:     sector3Start - sector2Start,
		Sector3Length:     float32(trackLength) - sector3Start,
	}
}
