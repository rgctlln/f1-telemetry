package main

import (
	"f1-telemetry/udp"
	"log"
	"net"
)

func main() {
	// Reads config
	var cfg udp.Config
	if err := udp.ReadConfigYaml(&cfg); err != nil {
		log.Fatal(err)
	}

	addr := net.UDPAddr{
		Port: cfg.Address.Port,
		IP:   net.ParseIP(cfg.Address.Host),
	}

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
