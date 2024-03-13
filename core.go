package core

import (
	"database/sql"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

// ConnectionHandler contains initialized connectors and their statuses
type ConnectionHandler struct {
	Mu sync.RWMutex

	// Prometheus

	PrometheusIsInitialized bool
	Prometheus              *v1.API
	Prometheus_err          error

	// SQLDB
	SQLDBIsInitialized bool
	SQLDB              *sqlx.DB
	SQLDBErr           error

	// Redis
	RedisIsInitialized bool
	Redis              *redis.Client
	RedisErr           error

	// RabbitMQ
	RabbitMQIsInitialized bool
	RabbitMQ              *amqp.Connection
	RabbitMQErr           error

	// Kafka
	KafkaIsInitialized bool
	KafkaProducer      sarama.SyncProducer
	KafkaConsumer      sarama.ConsumerGroup
	KafkaConsumerErr   error
	KafkaProducerErr   error

	// HealthCheck
	HealthCheckIsInitialized bool

	//MySQL
	MySQLIsInitialized bool
	MySQLDB            *sql.DB
	MySQLDBErr         error

	//Cockroach
	CockroachIsInitialized bool
	Cockroach              *sql.DB
	CockroachErr           error
}

var Handler ConnectionHandler

// CloseAllConnections closes all open connections
func (h *ConnectionHandler) CloseAllConnections() {
	h.Mu.Lock()
	defer h.Mu.Unlock()

	if h.Prometheus != nil {
		h.PrometheusIsInitialized = false
		// ToDo At this time, the Prometheus client does not require the connection to be closed explicitly. Leaving this space for future updates if this changes.
	}

	if h.SQLDB != nil {
		h.SQLDBIsInitialized = false
		if err := h.SQLDB.Close(); err != nil {
			logrus.Error("Error when closing connection with SQLDB:", err)
		} else {
			logrus.Info("SQLDB was closed")
		}
	}

	if h.Redis != nil {
		h.RedisIsInitialized = false
		if err := h.Redis.Close(); err != nil {
			logrus.Error("Error when closing connection with Redis:", err)
		} else {
			logrus.Info("Redis was closed")
		}
	}

	if h.RabbitMQ != nil {
		h.RabbitMQIsInitialized = false
		if err := h.RabbitMQ.Close(); err != nil {
			logrus.Error("Error when closing connection with RabbitMQ:", err)
		} else {
			logrus.Info("CockroachDB was closed")
		}
	}
	if h.Cockroach != nil {
		h.CockroachIsInitialized = false
		if err := h.Cockroach.Close(); err != nil {
			logrus.Error("Error when closing connection with CockroachDB:", err)
		} else {
			logrus.Info("CockroachDB closed")
		}
	}
	if h.KafkaProducer != nil {
		if err := h.KafkaProducer.Close(); err != nil {
			logrus.Error("Error when closing connection with KafkaProducer:", err)
		} else {
			logrus.Info("KafkaProducer was closed")
		}
	}

	if h.KafkaConsumer != nil {
		if err := h.KafkaConsumer.Close(); err != nil {
			logrus.Error("Error when closing connection with KafkaConsumer:", err)
		} else {
			logrus.Info("KafkaConsumer was closed")
		}
	}

	h.KafkaIsInitialized = false

}
