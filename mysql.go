package core

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

// mariaDB
func ConToMySQL(drivername, username, password, dbname, host, port string, done chan<- struct{}, data chan<- *sql.DB) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)
	for {
		db, err := sql.Open(drivername, dsn)
		if err == nil {
			logrus.Info("MySql sucsessfull conection")
			data <- db
			done <- struct{}{}
			return
		}
		logrus.Error("Error during connection to MysqlDB", err)
		time.Sleep(3 * time.Second)

	}
}
