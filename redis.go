package core

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

// NewRedis creates a new connector to Redis
func ConToRedis(host, port, password string, done chan<- struct{}, data chan<- *redis.Client) {
	addr := fmt.Sprintf("%s:%s", host, port)

	for {
		client := redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password, // Blank password means no password
			DB:       0,        // Default is DB 0
		})
		_, err := client.Ping(client.Context()).Result()
		if err == nil {
			fmt.Println("+++")
			logrus.Info("Redis sucsessfull conection")
			data <- client
			done <- struct{}{}
			return
		}
		logrus.Error("Error during connection to Redis: ", err)
		time.Sleep(3 * time.Second)
	}
}
