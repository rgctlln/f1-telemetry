package tracks

import (
	"f1-telemetry/packets"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

// Listens conn and gets WorldPositionX,Y,Z and writes to ./tracks/<trackName>_2024_racingline.txt
func GetTrackCoordinates(conn *net.UDPConn, trackName string) {
	path := fmt.Sprintf("./tracks/generated_tracks/%s_2024_racingline.txt", trackName)
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("failed to open %s: %v", path, err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close file: %v", err)
		}
	}()

	buffer := make([]byte, 2048)
	once := sync.Once{}

	for {
		ln, UDPaddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Fatalf("Error reading from UDP: %v", err)
		}

		once.Do(func() {
			log.Printf("Read from UDP: %v, trying to get track coordinates.", UDPaddr)
		})

		buffer = buffer[:ln]
		header := packets.ParseHeader(buffer)

		if header.PacketId == 2 {
			lapData := packets.ParseLapDataPacket(buffer)
			if lapData.LapData[header.PlayerCarIndex].CurrentLapNum == 2 {
				goto write
			}
		}
	}

write:
	startWriting(conn, f)
}

// This function writes X, Y, Z coordinates of the car on the track
func startWriting(conn *net.UDPConn, f *os.File) {
	buffer := make([]byte, 2048)

	for {
		ln, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Fatalf("Error reading from UDP: %v", err)
		}

		buffer = buffer[:ln]
		header := packets.ParseHeader(buffer)

		switch {
		case header.PacketId == 2:
			lapData := packets.ParseLapDataPacket(buffer)
			if lapData.LapData[header.PlayerCarIndex].CurrentLapNum == 3 {
				return
			}
		case header.PacketId == 0:
			motion := packets.ParseMotionPacket(buffer)
			playerCarIndex := header.PlayerCarIndex

			x := motion.CarMotionData[playerCarIndex].WorldPositionX
			y := motion.CarMotionData[playerCarIndex].WorldPositionY
			z := motion.CarMotionData[playerCarIndex].WorldPositionZ

			_, err := fmt.Fprintf(f, "%f,%f,%f\n", x, y, z)
			if err != nil {
				log.Printf("Error writing to file %v: %v", f, err)
			}

			//TODO
		}
	}
}
