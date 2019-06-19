package f1_2018_kafka_feeder

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
)

func PublishToKafka(producer *kafka.Producer, topic string, payload []byte, doneChan chan<- bool) {
	go func() {
		defer close(doneChan)

		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				m := ev
				if m.TopicPartition.Error != nil {
					log.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
				}
				return

			default:
				log.Printf("Ignored event: %s\n", ev)
			}
		}
	}()

	producer.ProduceChannel() <- &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value: payload,
	}
}
