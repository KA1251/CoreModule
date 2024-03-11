package core

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

// NewRedis creates a new connector to Redis
func ConToRedis(host, port, password string, done chan<- struct{}, data chan<- *redis.Client, connection *ConnectionHandler) {
	addr := fmt.Sprintf("%s:%s", host, port)

	for {
		client := redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password, // Blank password means no password
			DB:       0,        // Default is DB 0
		})
		_, connection.RedisErr = client.Ping(client.Context()).Result()
		if connection.RedisErr == nil {

			logrus.Info("Redis sucsessfull conection")
			data <- client
			done <- struct{}{}
			return
		}
		logrus.Error("Error during connection to Redis: ", connection.RedisErr)

		time.Sleep(3 * time.Second)
	}
}
