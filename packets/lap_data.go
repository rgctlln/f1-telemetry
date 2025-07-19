package packets

import (
	"encoding/binary"
	"math"
)

type PacketLapData struct {
	Header  PacketHeader
	LapData [22]LapData
}

type LapData struct {
	LastLapTimeInMS              uint32  // Last lap time in milliseconds
	CurrentLapTimeInMS           uint32  // Current time around the lap in milliseconds
	Sector1TimeInMS              uint16  // Sector 1 time milliseconds part
	Sector1TimeMinutesPart       uint8   // Sector 1 whole minute part
	Sector2TimeInMS              uint16  // Sector 2 time milliseconds part
	Sector2TimeMinutesPart       uint8   // Sector 2 whole minute part
	DeltaToCarInFrontMSPart      uint16  // Time delta to car in front milliseconds part
	DeltaToCarInFrontMinutesPart uint8   // Time delta to car in front whole minute part
	DeltaToRaceLeaderMSPart      uint16  // Time delta to race leader milliseconds part
	DeltaToRaceLeaderMinutesPart uint8   // Time delta to race leader whole minute part
	LapDistance                  float32 // Distance around current lap in metres
	TotalDistance                float32 // Total distance travelled in session in metres
	SafetyCarDelta               float32 // Delta in seconds for safety car
	CarPosition                  uint8   // Car race position
	CurrentLapNum                uint8   // Current lap number
	PitStatus                    uint8   // 0 = none, 1 = pitting, 2 = in pit area
	NumPitStops                  uint8   // Number of pit stops taken in this race
	Sector                       uint8   // 0 = sector1, 1 = sector2, 2 = sector3
	CurrentLapInvalid            uint8   // 0 = valid, 1 = invalid
	Penalties                    uint8   // Accumulated time penalties in seconds
	TotalWarnings                uint8   // Accumulated number of warnings issued
	CornerCuttingWarnings        uint8   // Number of corner cutting warnings issued
	NumUnservedDriveThroughPens  uint8   // Drive through penalties left to serve
	NumUnservedStopGoPens        uint8   // Stop & go penalties left to serve
	GridPosition                 uint8   // Grid position at race start
	DriverStatus                 uint8   // 0=garage,1=flying lap,2=in lap,3=out lap,4=on track
	ResultStatus                 uint8   // 0=invalid,1=inactive,2=active,3=finished,4=dnf,5=dsq,6=not classified,7=retired
	PitLaneTimerActive           uint8   // Pit lane timing: 0=inactive,1=active
	PitLaneTimeInLaneInMS        uint16  // Time spent in pit lane in ms
	PitStopTimerInMS             uint16  // Time of the actual pit stop in ms
	PitStopShouldServePen        uint8   // Whether the car should serve a penalty at this stop
	SpeedTrapFastestSpeed        float32 // Fastest speed through speed trap in km/h
	SpeedTrapFastestLap          uint8   // Lap number fastest speed was achieved (255 = not set)
}

func ParseLapDataPacket(data []byte) PacketLapData {
	header := ParseHeader(data)
	var lapDataArr [22]LapData

	const (
		base      = 29
		entrySize = 57
	)

	for i := 0; i < 22; i++ {
		offset := base + i*entrySize
		lapDataArr[i] = LapData{
			LastLapTimeInMS:    binary.LittleEndian.Uint32(data[offset+0 : offset+4]),
			CurrentLapTimeInMS: binary.LittleEndian.Uint32(data[offset+4 : offset+8]),

			Sector1TimeInMS:        binary.LittleEndian.Uint16(data[offset+8 : offset+10]),
			Sector1TimeMinutesPart: data[offset+10],

			Sector2TimeInMS:        binary.LittleEndian.Uint16(data[offset+11 : offset+13]),
			Sector2TimeMinutesPart: data[offset+13],

			DeltaToCarInFrontMSPart:      binary.LittleEndian.Uint16(data[offset+14 : offset+16]),
			DeltaToCarInFrontMinutesPart: data[offset+16],

			DeltaToRaceLeaderMSPart:      binary.LittleEndian.Uint16(data[offset+17 : offset+19]),
			DeltaToRaceLeaderMinutesPart: data[offset+19],

			LapDistance:    math.Float32frombits(binary.LittleEndian.Uint32(data[offset+20 : offset+24])),
			TotalDistance:  math.Float32frombits(binary.LittleEndian.Uint32(data[offset+24 : offset+28])),
			SafetyCarDelta: math.Float32frombits(binary.LittleEndian.Uint32(data[offset+28 : offset+32])),

			CarPosition:                 data[offset+32],
			CurrentLapNum:               data[offset+33],
			PitStatus:                   data[offset+34],
			NumPitStops:                 data[offset+35],
			Sector:                      data[offset+36],
			CurrentLapInvalid:           data[offset+37],
			Penalties:                   data[offset+38],
			TotalWarnings:               data[offset+39],
			CornerCuttingWarnings:       data[offset+40],
			NumUnservedDriveThroughPens: data[offset+41],
			NumUnservedStopGoPens:       data[offset+42],
			GridPosition:                data[offset+43],
			DriverStatus:                data[offset+44],
			ResultStatus:                data[offset+45],
			PitLaneTimerActive:          data[offset+46],

			PitLaneTimeInLaneInMS: binary.LittleEndian.Uint16(data[offset+47 : offset+49]),
			PitStopTimerInMS:      binary.LittleEndian.Uint16(data[offset+49 : offset+51]),
			PitStopShouldServePen: data[offset+51],

			SpeedTrapFastestSpeed: math.Float32frombits(binary.LittleEndian.Uint32(data[offset+52 : offset+56])),
			SpeedTrapFastestLap:   data[offset+56],
		}
	}

	return PacketLapData{
		Header:  header,
		LapData: lapDataArr,
	}
}
