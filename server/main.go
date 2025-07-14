package main

import (
	"f1-telemetry/tracks"
	"f1-telemetry/udp"
	"github.com/joho/godotenv"
	"log"
	"net"
	"os"
	"strconv"
)

func init() {
	err := godotenv.Load("server/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	var (
		cfg udp.Config
		err error
	)

	// Reads config
	if err := udp.ReadConfigYaml(&cfg); err != nil {
		log.Fatal(err)
	}
	addr := net.UDPAddr{
		Port: cfg.Address.Port,
		IP:   net.ParseIP(cfg.Address.Host),
	}

	// Reads env variable
	drw := false
	if val, ok := os.LookupEnv("DRW"); ok {
		drw, err = strconv.ParseBool(val)
		if err != nil {
			log.Println("Error parsing DRW env variable. Set as default.")
		}
	}

	if drw {
		log.Println("Track drawing mode is enabled")
		drwMode(addr, cfg.Drawer.TrackName)
	} else {
		log.Println("Track drawing mode is disabled")
		udpMode(addr)
	}
}

func drwMode(addr net.UDPAddr, trackName string) {
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Fatal(err)
	}

	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	tracks.GetTrackCoordinates(conn, trackName)
}

func udpMode(addr net.UDPAddr) {
	// Connects to client through udp
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Fatal(err)
	}

	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	// Starts listening procedure
	err = udp.StartListening(conn, addr)
	if err != nil {
		log.Fatal(err)
	}
}
