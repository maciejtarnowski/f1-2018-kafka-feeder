package packet_handler

import (
	"bytes"
	"encoding/binary"
)

func Dispatch(message []byte) (jsonPayload []byte, skip bool) {
	var packetId uint8
	err := binary.Read(bytes.NewBuffer(message[3:4]), binary.LittleEndian, &packetId)

	if err != nil {
		return []byte(""), true
	}

	switch packetId {
	case 0:
		return CarMotionHandler(message)
	case 6:
		return CarTelemetryHandler(message)
	default:
		return []byte(""), true
	}
}