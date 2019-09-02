package main

import (
	"github.com/maciejtarnowski/f1-2018-kafka-feeder"
	"github.com/maciejtarnowski/f1-2018-kafka-feeder/packet_handler"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
	"net"
	"os"
)

const LISTEN_ADDR = ":2018"

func main() {
	server, err := net.ListenPacket("udp", LISTEN_ADDR)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("UDP listening on %s", LISTEN_ADDR)
	defer server.Close()

	if len(os.Args) != 2 {
		log.Fatal("Invalid options, usage: <broker>")
	}

	broker := os.Args[1]

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})

	if err != nil {
		log.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	log.Printf("Created Producer %v\n", p)

	for {
		buf := make([]byte, 2048)
		n, _, err := server.ReadFrom(buf)
		if err != nil {
			continue
		}
		go handleData(buf[:n], p)
	}
}

func handleData(buf []byte, p *kafka.Producer) {
	jsonString, topic, skip := packet_handler.Dispatch(buf)

	if skip {
		return
	}

	doneChan := make(chan bool)

	f1_2018_kafka_feeder.PublishToKafka(p, topic, jsonString, doneChan)

	_ = <-doneChan
}
