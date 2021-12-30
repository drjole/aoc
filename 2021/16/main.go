package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputBytes, err := os.ReadFile("2021/16/input.txt")
	if err != nil {
		panic(err)
	}

	input := string(inputBytes)
	builder := strings.Builder{}
	for i := 0; i < len(input); i++ {
		b, _ := strconv.ParseInt(string(input[i]), 16, 8)
		for j := 3; j >= 0; j-- {
			_, _ = fmt.Fprintf(&builder, "%d", (b>>j)&1)
		}
	}
	transmission := builder.String()
	packets := make([]packet, 0)
	parsePacket(transmission, &packets)

	// First star
	versionsSum := 0
	for _, p := range packets {
		versionsSum += p.version
	}
	fmt.Println(versionsSum)

	// Second star
	fmt.Println(packets[len(packets)-1].payload)
}

func parsePacket(transmission string, packets *[]packet) packet {
	subPackets := make([]packet, 0)
	currentPacket := packet{}
	currentState := stateVersion
	bit := 0
	buffer := strings.Builder{}
	for {
		switch currentState {
		case stateVersion:
			currentPacket.version = parseBits(transmission, bit, 3)
			bit += 3
			currentState = stateTypeID
		case stateTypeID:
			currentPacket.typeID = parseBits(transmission, bit, 3)
			bit += 3
			if currentPacket.typeID == 4 {
				currentState = statePayload
			} else {
				currentState = stateLengthTypeID
			}
		case stateLengthTypeID:
			currentPacket.lengthTypeID = parseBits(transmission, bit, 1)
			bit += 1
			currentState = statePayload
		case statePayload:
			if currentPacket.typeID == 4 {
				// Literal value packet
				more := 1
				for more == 1 {
					more = parseBits(transmission, bit, 1)
					bit += 1
					buffer.WriteString(transmission[bit : bit+4])
					bit += 4
				}
			} else {
				if currentPacket.lengthTypeID == 0 {
					totalLength := parseBits(transmission, bit, 15)
					bit += 15
					start := bit
					for bit-start < totalLength {
						p := parsePacket(transmission[bit:], packets)
						subPackets = append(subPackets, p)
						bit += p.length
					}
				} else if currentPacket.lengthTypeID == 1 {
					numberOfSubPackets := parseBits(transmission, bit, 11)
					bit += 11
					for i := 0; i < numberOfSubPackets; i++ {
						p := parsePacket(transmission[bit:], packets)
						subPackets = append(subPackets, p)
						bit += p.length
					}
				} else {
					panic("this should never happen")
				}
				switch currentPacket.typeID {
				case 0:
					// Sum packet
					if len(subPackets) == 1 {
						currentPacket.payload = subPackets[0].payload
					} else {
						for _, s := range subPackets {
							currentPacket.payload += s.payload
						}
					}
				case 1:
					// Product packet
					if len(subPackets) == 1 {
						currentPacket.payload = subPackets[0].payload
					} else {
						currentPacket.payload = 1
						for _, s := range subPackets {
							currentPacket.payload *= s.payload
						}
					}
				case 2:
					// Minimum packet
					minValue := math.MaxInt
					for _, s := range subPackets {
						if s.payload < minValue {
							minValue = s.payload
						}
					}
					currentPacket.payload = minValue
				case 3:
					// Maximum packet
					maxValue := 0
					for _, s := range subPackets {
						if s.payload > maxValue {
							maxValue = s.payload
						}
					}
					currentPacket.payload = maxValue
				case 5:
					// Greater than packet
					if subPackets[0].payload > subPackets[1].payload {
						currentPacket.payload = 1
					} else {
						currentPacket.payload = 0
					}
				case 6:
					// Less than packet
					if subPackets[0].payload < subPackets[1].payload {
						currentPacket.payload = 1
					} else {
						currentPacket.payload = 0
					}
				case 7:
					// Equal to packet
					if subPackets[0].payload == subPackets[1].payload {
						currentPacket.payload = 1
					} else {
						currentPacket.payload = 0
					}
				}
			}
			currentState = stateDone
		case stateDone:
			if currentPacket.typeID == 4 {
				payloadString := buffer.String()
				currentPacket.payload = parseBits(payloadString, 0, len(payloadString))
			}
			currentPacket.length = bit
			*packets = append(*packets, currentPacket)
			return currentPacket
		}
	}
}

func parseBits(s string, start, n int) int {
	res, err := strconv.ParseInt(s[start:start+n], 2, 64)
	if err != nil {
		panic(err)
	}
	return int(res)
}

type state int

const (
	stateVersion state = iota
	stateTypeID
	stateLengthTypeID
	statePayload
	stateDone
)

type packet struct {
	version, typeID, lengthTypeID int
	payload                       int
	length                        int
}
