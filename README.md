# CORE INITIALLIZER 
golang модуль для быстрого развертывания сервисов. Совместим с sql/mysql базами данных, брокерами сообщений rabbitMQ/kafka и noSQL DB Redis.
## How to install 
go get github.com/KA1251/CoreModule
## how it works
```text
Для инциализации используется
var conn core.ConnectionHandler // объект хранящий инф-ию о подключениях
core.Initiallizing(&conn) //  инициализация подключения
conn.CloseAllConnections() // инициализация отключения
в случае неудачного подключения происходит реконнект и формируютя логи
 о неудачном подключении (в случае если подключение было успешным также формируются логи о подключении)
```
## Config list
эти переменные окружения используются для задания параметров подключения (при использовании какого либо подключения нужно установить параметр ..._ENABLED:T) :
```text
REDIS_ENABLED: 
REDIS_HOST:
REDIS_PORT:
REDIS_PASSWORD: 
RABBITMQ_ENABLED:
RABBITMQ_HOST: 
RABBITMQ_PORT:
RABBITMQ_USERNAME: 
RABBITMQ_PASSWORD:
PROMETHEUS_ENABLED:
PROMETHEUS_HOST:
PROMETHEUS_PORT: 
KAFKA_ENABLED:
KAFKA_PORT:
KAFKA_HOST:
KAFKA_USERNAME:
KAFKA_PASSWORD:
SQL_ENABLED:
SQL_PORT:
SQL_USERNAME:
SQL_PASSWORD:
SQL_HOST:
SQL_DB:
SQL_DRIVER:
MYSQL_ENABLED:
MYSQL_PORT:
MYSQL_HOST:
MYSQL_USERNAME:
MYSQL_PASSWORD:
MYSQL_DB:
MYSQL_DRIVER:
COCKROACH_ENABLED: 
COCKROACH_HOST: 
COCKROACH_USERNAME:
COCKROACH_PORT: 
COCKROACH_DB: 
COCKROACH_DRIVER: 
COCKROACH_APP: 
```
## Example of usage 1(redis)
Dockerfile:
```Dockerfile

FROM golang:latest

ENV REDIS_HOST redis
ENV REDIS_PORT 6379
ENV REDIS_PASSWORD=
ENV REDIS_ENABLED T

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main .

CMD ["./main"]
```
docker-compose:
``` .yml
  version: '3'
  services:
    redis:
      image: redis
      ports:
        - "6379:6379"
    goapp:
      build:
        context: .
        dockerfile: Dockerfile
      depends_on:
        - redis
```
main.go:
```go
package main

import (
	"connection_test/core"
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

func main() {
	var conn core.ConnectionHandler
	core.Initiallizing(&conn)

	// Простая операция с Redis
	ctx := context.Background()

	err := conn.Redis.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		logrus.Error("Failed to set key:", err)
		return
	}

	val, err := conn.Redis.Get(ctx, "key").Result()
	if err != nil {
		logrus.Error("Failed to get key:", err)
		return
	}

	fmt.Printf("Value: %s\n", val)
}
```
Output:
```text
redis-1  | 1:C 07 Mar 2024 11:48:41.281 # WARNING Memory overcommit must be enabled! Without it, a background save or replication may fail under low memory condition. Being disabled, it can also cause failures without low memory condition, see https://github.com/jemalloc/jemalloc/issues/1328. To fix this issue add 'vm.overcommit_memory = 1' to /etc/sysctl.conf and then reboot or run the command 'sysctl vm.overcommit_memory=1' for this to take effect.
redis-1  | 1:C 07 Mar 2024 11:48:41.281 * oO0OoO0OoO0Oo Redis is starting oO0OoO0OoO0Oo
redis-1  | 1:C 07 Mar 2024 11:48:41.281 * Redis version=7.2.4, bits=64, commit=00000000, modified=0, pid=1, just started
redis-1  | 1:C 07 Mar 2024 11:48:41.281 # Warning: no config file specified, using the default config. In order to specify a config file use redis-server /path/to/redis.conf
redis-1  | 1:M 07 Mar 2024 11:48:41.282 * monotonic clock: POSIX clock_gettime
redis-1  | 1:M 07 Mar 2024 11:48:41.282 * Running mode=standalone, port=6379.
redis-1  | 1:M 07 Mar 2024 11:48:41.283 * Server initialized
redis-1  | 1:M 07 Mar 2024 11:48:41.283 * Ready to accept connections tcp
goapp-1  | Value: value
goapp-1  | time="2024-03-07T11:48:42Z" level=info msg="Redis sucsessfull conection"
goapp-1 exited with code 0
redis-1  | 1:M 07 Mar 2024 12:48:42.099 * 1 changes in 3600 seconds. Saving...
redis-1  | 1:M 07 Mar 2024 12:48:42.124 * Background saving started by pid 19
redis-1  | 19:C 07 Mar 2024 12:48:42.144 * DB saved on disk
redis-1  | 19:C 07 Mar 2024 12:48:42.144 * Fork CoW for RDB: current 0 MB, peak 0 MB, average 0 MB
redis-1  | 1:M 07 Mar 2024 12:48:42.224 * Background saving terminated with success
```
## example of usage 2 (cockroachdb)
DockerFile:
``` DockerFile
FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main .

CMD ["./main"]
```
docker-compose:
``` .yml
version: '3'

services:
  cockroachdb:
    restart: always
    image: cockroachdb/cockroach:v21.1.10
    environment:
      COCKROACH_DB: defaultdb
      COCKROACH_USER: root 
    ports:
      - "26257:26257"
    volumes: 
      - cockroach-data:/cockroach/cockroach-data
    command: start-single-node --insecure

  go-app:
    build: ./
    command: ./main
    environment:
      COCKROACH_ENABLED: T
      COCKROACH_HOST: cockroachdb
      COCKROACH_USERNAME: root
      COCKROACH_PORT: 26257
      COCKROACH_DB: defaultdb
      COCKROACH_DRIVER: postgres
    depends_on:
      - cockroachdb

volumes:
  cockroach-data:
```
Код: создает бд, кидает данные, получает данные
```go
package main

import (
	"connection_test/core"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	var conn core.ConnectionHandler
	core.Initiallizing(&conn)
	if _, err := conn.Cockroach.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", os.Getenv("COCKROACH_DB"))); err != nil {
		log.Fatal("Error creating database:", err)
	}
	fmt.Println("created")
	var exists bool
	err := conn.Cockroach.QueryRow(`SELECT EXISTS(SELECT 1 FROM information_schema.schemata WHERE catalog_name = $1)`, os.Getenv("COCKROACH_DB")).Scan(&exists)
	if err != nil {
		log.Fatal("Error checking database existence:", err)
	}
	if !exists {
		log.Fatalf("Database %s does not exist11", conn.Cockroach)
	}

	if _, err := conn.Cockroach.Exec(`CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name STRING)`); err != nil {
		log.Fatal("Error creating table:", err)
	}

	if _, err := conn.Cockroach.Exec(`INSERT INTO users (name) VALUES ($1)`, "John Doe"); err != nil {
		log.Fatal("Error inserting data:", err)
	}

	rows, err := conn.Cockroach.Query(`SELECT id, name FROM users`)
	if err != nil {
		log.Fatal("Error querying data:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal("Error scanning row:", err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}
	if err := rows.Err(); err != nil {
		log.Fatal("Error iterating rows:", err)
	}
}

```
output:
```text
go-app-1       | time="2024-03-08T13:35:20Z" level=info msg="COCKROACH sucsessfull conection"
go-app-1       | created
go-app-1       | ID: 949729758244175873, Name: John Doe
```
