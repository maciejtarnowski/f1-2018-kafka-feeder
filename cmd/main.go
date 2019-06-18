package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/maciejtarnowski/f1-2018-kafka-feeder"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
	"net"
	"os"
)

func main() {
	// listen to incoming udp packets
	server, err := net.ListenPacket("udp", ":2018")
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <broker> <topic>\n",
			os.Args[0])
		os.Exit(1)
	}

	broker := os.Args[1]
	topic := os.Args[2]

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Producer %v\n", p)

	for {
		buf := make([]byte, 2048)
		n, _, err := server.ReadFrom(buf)
		if err != nil {
			continue
		}
		go handleData(buf[:n], p, topic)
	}
}

func handleData(buf []byte, p *kafka.Producer, topic string) {
	packet := f1_2018_kafka_feeder.PacketCarTelemetryData{}

	var packetId uint8
	err := binary.Read(bytes.NewBuffer(buf[3:4]), binary.LittleEndian, &packetId)
	if err != nil || packetId != 6 {
		return
	}

	err = binary.Read(bytes.NewBuffer(buf[:]), binary.LittleEndian, &packet)

	if err != nil {
		panic(err)
	}

	playerCarData := packet.Telemetry[packet.Header.PlayerCarIndex]

	jsonPacket := f1_2018_kafka_feeder.CarTelemetryJson{
		Speed: int(playerCarData.Speed),
		Throttle: int(playerCarData.Throttle),
		Steer: int(playerCarData.Steer),
		Brake: int(playerCarData.Brake),
		Gear: int(playerCarData.Gear),
		Drs: playerCarData.DRS == 1,
	}

	jsonString, err := json.Marshal(&jsonPacket)

	if err != nil {
		log.Println("Failed to encode JSON")
	}

	doneChan := make(chan bool)

	go func() {
		defer close(doneChan)
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				m := ev
				if m.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
				}
				return

			default:
				fmt.Printf("Ignored event: %s\n", ev)
			}
		}
	}()

	p.ProduceChannel() <- &kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny}, Value: jsonString}

	_ = <-doneChan
	//fmt.Println(packet.Telemetry[packet.Header.PlayerCarIndex].Speed)
}
