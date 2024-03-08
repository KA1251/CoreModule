package core

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func ConToCockRoach(drivername, username, password, dbname, host, port, appname string, done chan<- struct{}, data chan<- *sql.DB) {
	for {
		dsn := fmt.Sprintf("postgresql://%s@%s:%s/%s?sslmode=disable&application_name=%s", username, host, port, dbname, appname)
		db, err := sql.Open(drivername, dsn)
		if err == nil {
			logrus.Info("COCKROACH sucsessfull conection")
			data <- db
			done <- struct{}{}
			return
		}
		logrus.Error("Error during connection to COCKROACHDB", err)
		time.Sleep(3 * time.Second)

	}
}
