package core

import (
	"database/sql"
	"log"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/coder/redjet"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/streadway/amqp"
)

// ConnectionHandler contains initialized connectors and their statuses
type ConnectionHandler struct {
	Mu sync.RWMutex

	// Prometheus

	PrometheusIsInitialized bool
	Prometheus              *v1.API

	// SQLDB
	SQLDBIsInitialized bool
	SQLDB              *sqlx.DB

	// Redis
	RedisIsInitialized bool
	Redis              *redis.Client

	// Redis-Redjet
	RedisRedjetIsInitialized bool
	RedisRedjet              *redjet.Client

	// RabbitMQ
	RabbitMQIsInitialized bool
	RabbitMQ              *amqp.Connection

	// Kafka
	KafkaIsInitialized bool
	KafkaProducer      sarama.SyncProducer
	KafkaConsumer      sarama.Consumer

	// HealthCheck
	HealthCheckIsInitialized bool

	//MySQL
	MySQLIsInitialized bool
	MySQLDB            *sql.DB

	//Cockroach
	CockroachIsInitialized bool
	Cockroach              *sql.DB
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
			log.Printf("Error when closing connection with CockroachDB: %s\n", err)
		}
	}

	if h.Redis != nil {
		h.RedisIsInitialized = false
		if err := h.Redis.Close(); err != nil {
			log.Printf("Error when closing connection with Redis: %s\n", err)
		}
	}

	if h.RedisRedjet != nil {
		h.RedisRedjetIsInitialized = false
		if err := h.RedisRedjet.Close(); err != nil {
			log.Printf("Error when closing connection with Redis-Redjet: %s\n", err)
		}
	}

	if h.RabbitMQ != nil {
		h.RabbitMQIsInitialized = false
		if err := h.RabbitMQ.Close(); err != nil {
			log.Printf("Error when closing connection with RabbitMQ: %s\n", err)
		}
	}
	if h.Cockroach != nil {
		h.CockroachIsInitialized = false
		if err := h.Cockroach.Close(); err != nil {
			log.Printf("Error when closing connection with Cockroach: %s", err)
		}
	}
	if h.KafkaProducer != nil {
		if err := h.KafkaProducer.Close(); err != nil {
			log.Printf("Error when closing Kafka Producer: %s\n", err)
		}
	}

	if h.KafkaConsumer != nil {
		if err := h.KafkaConsumer.Close(); err != nil {
			log.Printf("Error when closing Kafka Consumer: %s\n", err)
		}
	}

	h.KafkaIsInitialized = false

}
