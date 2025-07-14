package udp

import (
	"f1-telemetry/packets"
	"gopkg.in/yaml.v3"
	"log"
	"net"
	"os"
	"sync"
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
	Drawer struct {
		TrackName string `yaml:"track_name"`
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

	err := conn.SetDeadline(time.Now().Add(20 * time.Second))
	if err != nil {
		return err
	}
	buffer := make([]byte, 2048)
	trackMap := make(map[uint8]struct{})
	var once sync.Once

	for {
		ln, UDPaddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("Error reading from UDP: %v", err)
			return err
		}

		once.Do(func() {
			log.Printf("Read from UDP: %v", UDPaddr)
		})

		header := packets.ParseHeader(buffer[:ln])
		parsePacket(header, buffer[:ln], trackMap)
	}
}
