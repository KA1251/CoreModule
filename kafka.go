package core

import (
	"time"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

// NewKafka creates a new connection to Kafka
func ConToKafka(host, port, username, password string, done chan<- struct{}, dataProd chan<- sarama.SyncProducer, dataCons chan<- sarama.Consumer, con *ConnectionHandler) {

	for {
		// Setting up Kafka Producer
		brokers := []string{host + ":" + port}

		producerConfig := sarama.NewConfig()
		producerConfig.Net.SASL.Enable = true
		producerConfig.Net.SASL.User = username
		producerConfig.Net.SASL.Password = password

		consumerConfig := sarama.NewConfig()
		consumerConfig.Net.SASL.Enable = true
		consumerConfig.Net.SASL.User = username
		consumerConfig.Net.SASL.Password = password

		producer, errProducer := sarama.NewSyncProducer(brokers, producerConfig)
		consumer, errConsumer := sarama.NewConsumer(brokers, consumerConfig)

		if errConsumer == nil || errProducer == nil {
			logrus.Info("Kafka producer sucsessfull connection")
			dataProd <- producer
			logrus.Info("Kafka consumer sucsessful connection")
			dataCons <- consumer
			done <- struct{}{}
			return
		}
		if errProducer != nil || errConsumer != nil {
			if errProducer != nil {
				con.KafkaProducerErr = errProducer
				logrus.Error("Error during connection to KAFKA(producer)", con.KafkaProducerErr)
			}
			if errConsumer != nil {
				con.KafkaConsumerErr = errProducer
				logrus.Error("Error during connection to KAFKA(consumer)", con.KafkaConsumerErr)
			}

		}
		time.Sleep(3 * time.Second)

	}
}
