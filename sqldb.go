package core

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// NewSqlDB creates a new connector to SQL Database
func ConToSql(host, port, user, password, dbName, driverName string, done chan<- struct{}, data chan<- *sqlx.DB, con *ConnectionHandler) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
	if driverName == "" {
		driverName = "postgres" // Default driver
	}

	for {
		db, err := sqlx.Connect(driverName, dsn)
		if err == nil {
			logrus.Info("sucsessfull conection to sql")
			data <- db
			done <- struct{}{}
			return
		}
		con.SQLDBErr = err
		logrus.Error("Error during connection to sqlDB: ", con.SQLDBErr)
		time.Sleep(3 * time.Second)
	}

}
