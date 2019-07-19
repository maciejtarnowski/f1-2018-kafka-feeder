package packet_handler

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	f1_2018_kafka_feeder "github.com/maciejtarnowski/f1-2018-kafka-feeder"
	"time"
)

func CarTelemetryHandler(message []byte) (jsonPayload []byte, topic string, skip bool) {
	packet := f1_2018_kafka_feeder.PacketCarTelemetryData{}

	err := binary.Read(bytes.NewBuffer(message[:]), binary.LittleEndian, &packet)

	if err != nil {
		return []byte(""), "", true
	}

	playerCarData := packet.Telemetry[packet.Header.PlayerCarIndex]

	jsonPacket := f1_2018_kafka_feeder.CarTelemetryJson{
		TimeStamp: time.Now().UnixNano() / int64(time.Millisecond),
		EventId: 6,
		Speed: int(playerCarData.Speed),
		Throttle: int(playerCarData.Throttle),
		Steer: int(playerCarData.Steer),
		Brake: int(playerCarData.Brake),
		Gear: int(playerCarData.Gear),
		Drs: playerCarData.DRS == 1,
	}

	jsonString, err := json.Marshal(&jsonPacket)

	if err != nil {
		return []byte(""), "", true
	}

	return jsonString, "raw_car_telemetry_data", false
}
