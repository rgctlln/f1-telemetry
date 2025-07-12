package packets

type PacketEventData struct {
	Header          PacketHeader
	EventStringCode [4]byte
}

func ParseEventPacket(data []byte) PacketEventData {
	header := ParseHeader(data)

	return PacketEventData{
		Header:          header,
		EventStringCode: [4]byte(data[29:33]),
	}
}
