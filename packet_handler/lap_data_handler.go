package packet_handler

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	f1_2018_kafka_feeder "github.com/maciejtarnowski/f1-2018-kafka-feeder"
	"time"
)

func LapDataHandler(message []byte) (jsonPayload []byte, topic string, skip bool) {
	packet := f1_2018_kafka_feeder.PacketLapData{}

	err := binary.Read(bytes.NewBuffer(message[:]), binary.LittleEndian, &packet)

	if err != nil {
		return []byte(""), "", true
	}

	playerLapData := packet.LapData[packet.Header.PlayerCarIndex]

	jsonPacket := f1_2018_kafka_feeder.LapDataJson{
		PlayerId: int(packet.Header.PlayerCarIndex),
		TimeStamp: time.Now().UnixNano() / int64(time.Millisecond),
		EventId: 2,
		CurrentLapNumber: int(playerLapData.CurrentLapNum),
		LastLapTime: float64(playerLapData.LastLapTime),
	}

	jsonString, err := json.Marshal(&jsonPacket)

	if err != nil {
		return []byte(""), "", true
	}

	return jsonString, "raw_lap_data", false
}
