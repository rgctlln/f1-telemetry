package packets

import "encoding/binary"

type PacketSessionData struct {
	Header            PacketHeader
	Weather           uint8
	TrackTemperature  int8
	AirTemperature    int8
	TotalLaps         uint8
	TrackLength       uint16
	SessionType       uint8
	TrackId           int8
	Formula           uint8
	SessionTimeLeft   uint16
	SessionDuration   uint16
	PitSpeedLimit     uint8
	GamePaused        uint8
	IsSpectating      uint8
	SpectatorCarIndex uint8
}

func ParseSessionPacket(data []byte) PacketSessionData {
	header := ParseHeader(data)

	return PacketSessionData{
		Header:            header,
		Weather:           data[25],
		TrackTemperature:  int8(data[26]),
		AirTemperature:    int8(data[27]),
		TotalLaps:         data[28],
		TrackLength:       binary.LittleEndian.Uint16(data[29:31]),
		SessionType:       data[31],
		TrackId:           int8(data[32]),
		Formula:           data[33],
		SessionTimeLeft:   binary.LittleEndian.Uint16(data[34:36]),
		SessionDuration:   binary.LittleEndian.Uint16(data[36:38]),
		PitSpeedLimit:     data[38],
		GamePaused:        data[39],
		IsSpectating:      data[40],
		SpectatorCarIndex: data[41],
	}
}
