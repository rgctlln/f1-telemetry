package packets

import (
	"encoding/binary"
	"math"
)

type PacketMotionData struct {
	Header        PacketHeader
	CarMotionData [22]CarMotionData
}

type CarMotionData struct {
	WorldPositionX     float32 // World space X position - metres
	WorldPositionY     float32 // World space Y position - metres
	WorldPositionZ     float32 // World space Z position - metres
	WorldVelocityX     float32 // Velocity in world space X – metres/s
	WorldVelocityY     float32 // Velocity in world space Y
	WorldVelocityZ     float32 // Velocity in world space Z
	WorldForwardDirX   int16   // World space forward X direction (normalised)
	WorldForwardDirY   int16   // World space forward Y direction (normalised)
	WorldForwardDirZ   int16   // World space forward Z direction (normalised)
	WorldRightDirX     int16   // World space right X direction (normalised)
	WorldRightDirY     int16   // World space right Y direction (normalised)
	WorldRightDirZ     int16   // World space right Z direction (normalised)
	GForceLateral      float32 // Lateral G-Force component
	GForceLongitudinal float32 // Longitudinal G-Force component
	GForceVertical     float32 // Vertical G-Force component
	Yaw                float32 // Yaw angle in radians
	Pitch              float32 // Pitch angle in radians
	Roll               float32 // Roll angle in radians
}

func ParseMotionPacket(data []byte) PacketMotionData {
	header := ParseHeader(data)
	var carData [22]CarMotionData
	base := 29 // Header is 29 bytes

	for i := 0; i < 22; i++ {
		offset := base + i*60 // Each CarMotionData struct is 60 bytes
		carData[i] = CarMotionData{
			WorldPositionX:     math.Float32frombits(binary.LittleEndian.Uint32(data[offset+0 : offset+4])),
			WorldPositionY:     math.Float32frombits(binary.LittleEndian.Uint32(data[offset+4 : offset+8])),
			WorldPositionZ:     math.Float32frombits(binary.LittleEndian.Uint32(data[offset+8 : offset+12])),
			WorldVelocityX:     math.Float32frombits(binary.LittleEndian.Uint32(data[offset+12 : offset+16])),
			WorldVelocityY:     math.Float32frombits(binary.LittleEndian.Uint32(data[offset+16 : offset+20])),
			WorldVelocityZ:     math.Float32frombits(binary.LittleEndian.Uint32(data[offset+20 : offset+24])),
			WorldForwardDirX:   int16(binary.LittleEndian.Uint16(data[offset+24 : offset+26])),
			WorldForwardDirY:   int16(binary.LittleEndian.Uint16(data[offset+26 : offset+28])),
			WorldForwardDirZ:   int16(binary.LittleEndian.Uint16(data[offset+28 : offset+30])),
			WorldRightDirX:     int16(binary.LittleEndian.Uint16(data[offset+30 : offset+32])),
			WorldRightDirY:     int16(binary.LittleEndian.Uint16(data[offset+32 : offset+34])),
			WorldRightDirZ:     int16(binary.LittleEndian.Uint16(data[offset+34 : offset+36])),
			GForceLateral:      math.Float32frombits(binary.LittleEndian.Uint32(data[offset+36 : offset+40])),
			GForceLongitudinal: math.Float32frombits(binary.LittleEndian.Uint32(data[offset+40 : offset+44])),
			GForceVertical:     math.Float32frombits(binary.LittleEndian.Uint32(data[offset+44 : offset+48])),
			Yaw:                math.Float32frombits(binary.LittleEndian.Uint32(data[offset+48 : offset+52])),
			Pitch:              math.Float32frombits(binary.LittleEndian.Uint32(data[offset+52 : offset+56])),
			Roll:               math.Float32frombits(binary.LittleEndian.Uint32(data[offset+56 : offset+60])),
		}
	}

	return PacketMotionData{
		Header:        header,
		CarMotionData: carData,
	}
}
