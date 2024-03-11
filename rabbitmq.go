package core

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

// NewRabbitMQ creates a new connection to RabbitMQ
func ConnToRabbitMQ(host, port, user, password string, done chan<- struct{}, data chan<- *amqp.Connection, con *ConnectionHandler) {
	for {
		amqpURI := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port)
		conn, err := amqp.Dial(amqpURI)
		if err == nil {
			logrus.Info("RabbitMQ sucsessful connection")
			data <- conn
			done <- struct{}{}
			return
		}
		con.RabbitMQErr = err
		logrus.Error("Errror during connection to RabbitMQ", con.RabbitMQ)
		time.Sleep(3 * time.Second)
	}
}
