package udp

import (
	"f1-telemetry/packets"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"net"
	"os"
	"time"
)

/*

todo documentation

*/

type Config struct {
	Address struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}
}

func ReadConfigYaml(cfg *Config) error {
	file, err := os.Open("udp/config.yaml")
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("Error closing file: %v", err)
		}
	}(file)

	return yaml.NewDecoder(file).Decode(cfg)
}

func StartListening(conn *net.UDPConn, addr net.UDPAddr) error {
	log.Printf("Listening on %v:%v\n", addr.String(), conn.LocalAddr().(*net.UDPAddr).Port)
	err := conn.SetDeadline(time.Now().Add(1 * time.Second))
	if err != nil {
		return err
	}

	buffer := make([]byte, 2048)
	ln, UDPaddr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		log.Printf("Error reading from UDP: %v", err)
		return err
	}
	fmt.Printf("Read a message from %v %s \n", UDPaddr, buffer)

	header := packets.ParseHeader(buffer[:ln])
	if header.PacketId == 1 {
		session := packets.ParseSessionPacket(buffer[:ln])
		fmt.Printf("Session packet from %v:\n", UDPaddr)
		fmt.Printf("Weather: %d, Track Temp: %d°C, Air Temp: %d°C\n",
			session.Weather, session.TrackTemperature, session.AirTemperature)
		fmt.Printf("Total laps: %d, Track Length: %dm\n",
			session.TotalLaps, session.TrackLength)
		fmt.Printf("Session time left: %d sec, Session duration: %d sec\n",
			session.SessionTimeLeft, session.SessionDuration)
	}

	return nil
}
