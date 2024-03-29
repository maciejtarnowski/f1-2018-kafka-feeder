package packet_handler

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"github.com/maciejtarnowski/f1-2018-kafka-feeder"
	"time"
)

func CarMotionHandler(message []byte) (jsonPayload []byte, topic string, skip bool) {
	packet := f1_2018_kafka_feeder.PacketMotionData{}

	err := binary.Read(bytes.NewBuffer(message[:]), binary.LittleEndian, &packet)

	if err != nil {
		return []byte(""), "", true
	}

	playerCarData := packet.MotionData[packet.Header.PlayerCarIndex]

	jsonPacket := f1_2018_kafka_feeder.CarMotionJson{
		PlayerId: int(packet.Header.PlayerCarIndex),
		TimeStamp: time.Now().UnixNano() / int64(time.Millisecond),
		EventId: 0,
		PosX: playerCarData.WorldPositionX,
		PosY: playerCarData.WorldPositionY,
		PosZ: playerCarData.WorldPositionZ,
	}

	jsonString, err := json.Marshal(&jsonPacket)

	if err != nil {
		return []byte(""), "", true
	}

	return jsonString, "raw_car_motion_data", false
}