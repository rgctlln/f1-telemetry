package packets

import (
	"encoding/binary"
	"math"
)

type PacketHeader struct {
	PacketFormat            uint16  // Packet format (e.g., 2024)
	GameYear                uint8   // Game year (last two digits, e.g., 24)
	GameMajorVersion        uint8   // Game major version (e.g., X.00)
	GameMinorVersion        uint8   // Game minor version (e.g., 1.XX)
	PacketVersion           uint8   // Version of this packet type (starts at 1)
	PacketId                uint8   // Packet type identifier (0 - 14)
	SessionUID              uint64  // Unique identifier for the session
	SessionTime             float32 // Session timestamp
	FrameIdentifier         uint32  // Identifier of the frame when data was retrieved
	OverallFrameIdentifier  uint32  // Overall frame identifier (does not reset after flashbacks)
	PlayerCarIndex          uint8   // Index of the player’s car in the array
	SecondaryPlayerCarIndex uint8   // Index of the secondary player’s car (split-screen), 255 if none
}

func ParseHeader(data []byte) PacketHeader {
	return PacketHeader{
		PacketFormat:            binary.LittleEndian.Uint16(data[0:2]),
		GameYear:                data[2],
		GameMajorVersion:        data[3],
		GameMinorVersion:        data[4],
		PacketVersion:           data[5],
		PacketId:                data[6],
		SessionUID:              binary.LittleEndian.Uint64(data[7:15]),
		SessionTime:             math.Float32frombits(binary.LittleEndian.Uint32(data[15:19])),
		FrameIdentifier:         binary.LittleEndian.Uint32(data[19:23]),
		OverallFrameIdentifier:  binary.LittleEndian.Uint32(data[23:27]),
		PlayerCarIndex:          data[27],
		SecondaryPlayerCarIndex: data[28],
	}
}
