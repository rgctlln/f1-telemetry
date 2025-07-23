package tracks

import (
	"f1-telemetry/packets"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

//TODO add description how this class works

// Listens conn and gets WorldPositionX,Y,Z and writes to ./tracks/<trackName>_2024_racingline.txt
func GetTrackCoordinates(conn *net.UDPConn, trackName string) {
	path := fmt.Sprintf("./tracks/generated_tracks/%s_2024_racingline.txt", trackName)

	if info, err := os.Stat(path); err == nil && info.Size() > 0 {
		fmt.Printf("File %s already exists. Delete and rewrite? (Y/N): ", path)

		var response string
		_, err := fmt.Scanln(&response)
		if err != nil {
			log.Fatalf("Input error: %v", err)
		}

		switch response {
		case "Y", "y", "Yes", "YES", "yes":
			err := os.Remove(path)
			if err != nil {
				log.Fatalf("Unable to delete file: %v", err)
			}
		default:
			log.Println("Operation canceled.")
			return
		}
	}

	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("failed to open %s: %v", path, err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close file: %v", err)
		}
	}()

	if _, err = fmt.Fprint(f, "x,y,z,dist\n"); err != nil {
		log.Fatalf("failed to write to file: %v", err)
	}

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

		header := packets.ParseHeader(buffer[:ln])

		if header.PacketId == 2 {
			lapData := packets.ParseLapDataPacket(buffer[:ln])
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
	collected := false
	trackLength := 0

	for {
		ln, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Fatalf("Error reading from UDP: %v", err)
		}

		header := packets.ParseHeader(buffer[:ln])

		if header.PacketId == 1 {
			session := packets.ParseSessionPacket(buffer[:ln])
			trackLength = int(session.TrackLength)
			goto collector
		}
	}

collector:
	for {
		ln, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Fatalf("Error reading from UDP: %v", err)
		}

		header := packets.ParseHeader(buffer[:ln])

		switch {
		case header.PacketId == 2:
			if !collected {
				continue
			}

			lapData := packets.ParseLapDataPacket(buffer[:ln])
			if lapData.LapData[header.PlayerCarIndex].CurrentLapNum == 3 {
				if collected {
					err = removeLastLine(f.Name())
					if err != nil {
						log.Fatalf("Unable to remove last line from file: %v", err)
					}
				}
				return
			}

			dist := float32(int(lapData.LapData[header.PlayerCarIndex].TotalDistance) % trackLength)
			_, err := fmt.Fprintf(f, "%f\n", dist)
			if err != nil {
				log.Fatalf("Error writing to file: %v", err)
			}

			collected = false
		case header.PacketId == 0:
			if collected {
				continue
			}

			motion := packets.ParseMotionPacket(buffer[:ln])
			playerCarIndex := header.PlayerCarIndex

			x := motion.CarMotionData[playerCarIndex].WorldPositionX
			y := motion.CarMotionData[playerCarIndex].WorldPositionY
			z := motion.CarMotionData[playerCarIndex].WorldPositionZ

			_, err := fmt.Fprintf(f, "%f,%f,%f,", x, y, z)
			if err != nil {
				log.Printf("Error writing to file %v: %v", f, err)
			}

			collected = true
		}
	}
}

func removeLastLine(path string) error {
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("Error closing file: %v", err)
		}
	}(file)

	stat, err := file.Stat()
	if err != nil {
		return err
	}
	size := stat.Size()
	if size == 0 {
		return nil
	}

	var truncateAt int64 = -1
	buf := make([]byte, 1)

	for i := size - 1; i >= 0; i-- {
		_, err := file.ReadAt(buf, i)
		if err != nil {
			return err
		}

		if buf[0] == '\n' {
			if i == size-1 {
				continue
			}

			truncateAt = i + 1
			break
		}
	}

	if truncateAt == -1 {
		truncateAt = 0
	}

	return file.Truncate(truncateAt)
}
