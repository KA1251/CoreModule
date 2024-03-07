package core

import (
	"database/sql"
	"os"

	"github.com/Shopify/sarama"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func Initiallizing(con *ConnectionHandler) {

	logrus.SetLevel(logrus.DebugLevel)

	doneRedis := make(chan struct{}, 1)
	updatesRedis := make(chan *redis.Client, 1)

	doneSQL := make(chan struct{}, 1)
	updatesSQL := make(chan *sqlx.DB, 1)

	doneKafka := make(chan struct{}, 1)
	updatesKafkaProd := make(chan sarama.SyncProducer, 1)
	updatesKafkaCons := make(chan sarama.Consumer, 1)

	doneRabbit := make(chan struct{}, 1)
	updatesRabbit := make(chan *amqp.Connection, 1)

	donePr := make(chan struct{}, 1)
	updatesPr := make(chan *v1.API, 1)

	doneMySQL := make(chan struct{}, 1)
	updatesMySQL := make(chan *sql.DB, 1)

	//sql connection checker

	if os.Getenv("SQL_ENABLED") == "T" {
		go ConToSql(os.Getenv("SQL_HOST"), os.Getenv("SQL_PORT"), os.Getenv("SQL_USERNAME"), os.Getenv("SQL_PASSWORD"), os.Getenv("SQL_DB"), os.Getenv("SQL_DRIVER"), doneSQL, updatesSQL)

	}

	//kafka connection checker
	if os.Getenv("KAFKA_ENABLED") == "T" {
		go ConToKafka(os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT"), os.Getenv("KAFKA_USERNAME"), os.Getenv("KAFKA_PASSWORD"), doneKafka, updatesKafkaProd, updatesKafkaCons)

	}

	//rabbitMQ connection checker
	if os.Getenv("RABBITMQ_ENABLED") == "T" {
		go ConnToRabbitMQ(os.Getenv("RABBITMQ_HOST"), os.Getenv("RABBITMQ_PORT"), os.Getenv("RABBITMQ_USERNAME"), os.Getenv("RABBITMQ_PASSWORD"), doneRabbit, updatesRabbit)

	}

	//Prometheus connection checker
	if os.Getenv("PROMETHEUS_ENABLED") == "T" {
		go ConToPrometheus(os.Getenv("PROMETHEUS_HOST"), os.Getenv("PROMETHEUS_PORT"), donePr, updatesPr)

	}
	//reddis connection checker
	redis_str := os.Getenv("REDIS_ENABLED")
	if redis_str == "T" {
		logrus.Info("+")
		go ConToRedis(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"), os.Getenv("REDIS_PASSWORD"), doneRedis, updatesRedis)

	}

	//MySQL connection checker
	if os.Getenv("MYSQL_ENABLED") == "T" {
		go ConToMySQL(os.Getenv("MYSQL_DRIVER"), os.Getenv("MYSQL_USERNAME"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DB"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), doneMySQL, updatesMySQL)

	}

	if os.Getenv("SQL_ENABLED") == "T" {
		<-doneSQL
		db := <-updatesSQL
		con.SQLDB = db
		con.SQLDBIsInitialized = true
	}

	if os.Getenv("KAFKA_ENABLED") == "T" {
		<-doneKafka
		kafkaProd := <-updatesKafkaProd
		kafkaCons := <-updatesKafkaCons
		con.KafkaProducer = kafkaProd
		con.KafkaConsumer = kafkaCons
		con.KafkaIsInitialized = true
	}

	if os.Getenv("RABBITMQ_ENABLED") == "T" {
		<-doneRabbit
		conRabbit := <-updatesRabbit
		con.RabbitMQ = conRabbit
		con.RabbitMQIsInitialized = true
	}

	if os.Getenv("PROMETHEUS_ENABLED") == "T" {
		<-donePr
		promCon := <-updatesPr
		con.Prometheus = promCon
		con.PrometheusIsInitialized = true
	}

	if os.Getenv("REDIS_ENABLED") == "T" {
		<-doneRedis
		client := <-updatesRedis
		con.Redis = client
		con.RedisIsInitialized = true
	}

	if os.Getenv("MYSQL_ENABLED") == "T" {
		dbMySql := <-updatesMySQL

		con.MySQLDB = dbMySql
		con.MySQLIsInitialized = true
		<-doneMySQL
	}

}
