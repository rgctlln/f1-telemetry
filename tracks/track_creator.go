package tracks

import (
	"f1-telemetry/packets"
	"fmt"
	"log"
	"net"
	"os"
)

// GetTrackCoordinates слушает conn и дописывает
// WorldPositionX,Y,Z первой машины в файл ./tracks/<trackName>_2024_racingline.txt
func GetTrackCoordinates(conn *net.UDPConn, trackName string) {
	// Открываем (или создаём) файл в режиме append
	path := fmt.Sprintf("./tracks/%s_2024_racingline.txt", trackName)
	f, err := os.OpenFile(path,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("failed to open %s: %v", path, err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close file: %v", err)
		}
	}()

	buf := make([]byte, 2048)
	for {
		// читаем один UDP-пакет
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Printf("read error: %v", err)
			continue
		}

		// парсим MotionPacket
		motion := packets.ParseMotionPacket(buf[:n])

		// берём координаты первой машины
		x := motion.CarMotionData[0].WorldPositionX
		y := motion.CarMotionData[0].WorldPositionY
		z := motion.CarMotionData[0].WorldPositionZ

		// дописываем в файл
		line := fmt.Sprintf("%f,%f,%f\n", x, y, z)
		if _, err := f.WriteString(line); err != nil {
			log.Printf("write error: %v", err)
		}
	}
}
