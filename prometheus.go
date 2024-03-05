package core

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/sirupsen/logrus"
)

// NewPrometheus creates a new client to interact with the Prometheus API
func ConToPrometheus(host, port string, done chan<- struct{}, data chan<- *v1.API) {
	for {
		address := fmt.Sprintf("http://%s:%s", host, port)
		// Setting up an HTTP client
		httpClient := &http.Client{
			Timeout: 30 * time.Second,
		}

		// Setting up an HTTP client
		clientConfig := api.Config{
			Address:      address,
			RoundTripper: httpClient.Transport,
		}
		client, err := api.NewClient(clientConfig)
		if err == nil {
			v1api := v1.NewAPI(client)
			data <- &v1api
			done <- struct{}{}
			return
		}
		logrus.Error("Error during connection to Prometheus", err)
		time.Sleep(3 * time.Second)
	}
}
