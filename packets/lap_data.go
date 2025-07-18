package packets

import (
	"encoding/binary"
	"math"
)

type PacketLapData struct {
	Header  PacketHeader
	LapData [22]lapData
}

type lapData struct {
	lastLapTimeInMS              uint32  // Last lap time in milliseconds
	currentLapTimeInMS           uint32  // Current time around the lap in milliseconds
	sector1TimeInMS              uint16  // Sector 1 time milliseconds part
	sector1TimeMinutesPart       uint8   // Sector 1 whole minute part
	sector2TimeInMS              uint16  // Sector 2 time milliseconds part
	sector2TimeMinutesPart       uint8   // Sector 2 whole minute part
	deltaToCarInFrontMSPart      uint16  // Time delta to car in front milliseconds part
	deltaToCarInFrontMinutesPart uint8   // Time delta to car in front whole minute part
	deltaToRaceLeaderMSPart      uint16  // Time delta to race leader milliseconds part
	deltaToRaceLeaderMinutesPart uint8   // Time delta to race leader whole minute part
	lapDistance                  float32 // Distance around current lap in metres
	totalDistance                float32 // Total distance travelled in session in metres
	safetyCarDelta               float32 // Delta in seconds for safety car
	carPosition                  uint8   // Car race position
	currentLapNum                uint8   // Current lap number
	pitStatus                    uint8   // 0 = none, 1 = pitting, 2 = in pit area
	numPitStops                  uint8   // Number of pit stops taken in this race
	sector                       uint8   // 0 = sector1, 1 = sector2, 2 = sector3
	currentLapInvalid            uint8   // 0 = valid, 1 = invalid
	penalties                    uint8   // Accumulated time penalties in seconds
	totalWarnings                uint8   // Accumulated number of warnings issued
	cornerCuttingWarnings        uint8   // Number of corner cutting warnings issued
	numUnservedDriveThroughPens  uint8   // Drive through penalties left to serve
	numUnservedStopGoPens        uint8   // Stop & go penalties left to serve
	gridPosition                 uint8   // Grid position at race start
	driverStatus                 uint8   // 0=garage,1=flying lap,2=in lap,3=out lap,4=on track
	resultStatus                 uint8   // 0=invalid,1=inactive,2=active,3=finished,4=dnf,5=dsq,6=not classified,7=retired
	pitLaneTimerActive           uint8   // Pit lane timing: 0=inactive,1=active
	pitLaneTimeInLaneInMS        uint16  // Time spent in pit lane in ms
	pitStopTimerInMS             uint16  // Time of the actual pit stop in ms
	pitStopShouldServePen        uint8   // Whether the car should serve a penalty at this stop
	speedTrapFastestSpeed        float32 // Fastest speed through speed trap in km/h
	speedTrapFastestLap          uint8   // Lap number fastest speed was achieved (255 = not set)
}

func ParseLapDataPacket(data []byte) PacketLapData {
	header := ParseHeader(data)
	var lapDataArr [22]lapData

	const (
		base      = 29
		entrySize = 57
	)

	for i := 0; i < 22; i++ {
		offset := base + i*entrySize
		lapDataArr[i] = lapData{
			lastLapTimeInMS:    binary.LittleEndian.Uint32(data[offset+0 : offset+4]),
			currentLapTimeInMS: binary.LittleEndian.Uint32(data[offset+4 : offset+8]),

			sector1TimeInMS:        binary.LittleEndian.Uint16(data[offset+8 : offset+10]),
			sector1TimeMinutesPart: data[offset+10],

			sector2TimeInMS:        binary.LittleEndian.Uint16(data[offset+11 : offset+13]),
			sector2TimeMinutesPart: data[offset+13],

			deltaToCarInFrontMSPart:      binary.LittleEndian.Uint16(data[offset+14 : offset+16]),
			deltaToCarInFrontMinutesPart: data[offset+16],

			deltaToRaceLeaderMSPart:      binary.LittleEndian.Uint16(data[offset+17 : offset+19]),
			deltaToRaceLeaderMinutesPart: data[offset+19],

			lapDistance:    math.Float32frombits(binary.LittleEndian.Uint32(data[offset+20 : offset+24])),
			totalDistance:  math.Float32frombits(binary.LittleEndian.Uint32(data[offset+24 : offset+28])),
			safetyCarDelta: math.Float32frombits(binary.LittleEndian.Uint32(data[offset+28 : offset+32])),

			carPosition:                 data[offset+32],
			currentLapNum:               data[offset+33],
			pitStatus:                   data[offset+34],
			numPitStops:                 data[offset+35],
			sector:                      data[offset+36],
			currentLapInvalid:           data[offset+37],
			penalties:                   data[offset+38],
			totalWarnings:               data[offset+39],
			cornerCuttingWarnings:       data[offset+40],
			numUnservedDriveThroughPens: data[offset+41],
			numUnservedStopGoPens:       data[offset+42],
			gridPosition:                data[offset+43],
			driverStatus:                data[offset+44],
			resultStatus:                data[offset+45],
			pitLaneTimerActive:          data[offset+46],

			pitLaneTimeInLaneInMS: binary.LittleEndian.Uint16(data[offset+47 : offset+49]),
			pitStopTimerInMS:      binary.LittleEndian.Uint16(data[offset+49 : offset+51]),
			pitStopShouldServePen: data[offset+51],

			speedTrapFastestSpeed: math.Float32frombits(binary.LittleEndian.Uint32(data[offset+52 : offset+56])),
			speedTrapFastestLap:   data[offset+56],
		}
	}

	return PacketLapData{
		Header:  header,
		LapData: lapDataArr,
	}
}
