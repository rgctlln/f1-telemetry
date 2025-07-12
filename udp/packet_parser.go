package udp

import (
	"f1-telemetry/packets"
	"f1-telemetry/packets/mapper"
	"fmt"
	"strconv"
)

func parsePacket(header packets.PacketHeader, buffer []byte, trackMap map[uint8]struct{}) {
	switch {
	case header.PacketId == 1:
		parsePacketID1(header, buffer, trackMap)
	case header.PacketId == 3:
		parsePacketID3(header, buffer)
	}
}

func parsePacketID1(header packets.PacketHeader, buffer []byte, trackMap map[uint8]struct{}) {
	if _, ok := trackMap[header.PacketId]; ok {
		return
	}

	trackMap[header.PacketId] = struct{}{}
	session := packets.ParseSessionPacket(buffer)
	fmt.Println("Welcome to " + mapper.GetMappedTrack(session.TrackId) +
		", track length is approximately " + strconv.Itoa(int(session.TrackLength)) + " metres!\n")
	fmt.Println("Sector 1 length is " + strconv.Itoa(int(session.Sector1Length)) + " metres!\n")
	fmt.Println("Sector 2 length is " + strconv.Itoa(int(session.Sector2Length)) + " metres!\n")
	fmt.Println("Sector 3 length is " + strconv.Itoa(int(session.Sector3Length)) + " metres!\n")
	fmt.Println("Weather today is " + mapper.GetMappedWeather(session.Weather) +
		", temperature is " + strconv.Itoa(int(session.AirTemperature)) + "℃")
}

func parsePacketID3(header packets.PacketHeader, buffer []byte) {
	event := packets.ParseEventPacket(buffer)
	fmt.Println(string(event.EventStringCode[:]))
}
