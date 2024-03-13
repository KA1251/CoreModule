package core

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

// NewKafka creates a new connection to Kafka
func ConToKafka(host, port, username, password string, done chan<- struct{}, dataProd chan<- sarama.SyncProducer, dataCons chan<- sarama.ConsumerGroup, con *ConnectionHandler) {

	for {
		// Setting up Kafka Producer
		brokers := []string{host + ":" + port}

		producerConfig := sarama.NewConfig()
		producerConfig.Producer.Return.Successes = true
		/*producerConfig.Net.SASL.Enable = false
		producerConfig.Net.SASL.User = username
		producerConfig.Net.SASL.Password = password*/

		consumerConfig := sarama.NewConfig()

		/*consumerConfig.Net.SASL.Enable = false
		consumerConfig.Net.SASL.User = username
		consumerConfig.Net.SASL.Password = password*/
		consumerConfig.Consumer.Return.Errors = true
		consumer, errConsumer := sarama.NewConsumerGroup(brokers, "group", consumerConfig)
		producer, errProducer := sarama.NewSyncProducer(brokers, producerConfig)
		//consumer, errConsumer := sarama.NewConsumer(brokers, consumerConfig)

		if errConsumer == nil || errProducer == nil {
			logrus.Info("Kafka producer sucsessfull connection")
			logrus.Info("Kafka consumer sucsessful connection")
			dataProd <- producer
			dataCons <- consumer
			done <- struct{}{}

			return
		}
		if errProducer != nil || errConsumer != nil {
			if errProducer != nil {
				fmt.Print("err:", errProducer)
				//con.KafkaProducerErr = errProducer
				logrus.Error("Error during connection to KAFKA(producer)", errProducer)
			}
			if errConsumer != nil {
				//con.KafkaConsumerErr = errProducer
				logrus.Error("Error during connection to KAFKA(consumer)", errConsumer)
			}

		}
		time.Sleep(3 * time.Second)

	}
}
