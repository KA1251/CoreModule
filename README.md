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
## Example of usage 1
```text
В примере я использую conf.txt для установки переменных окружения.(всегда создаем conf.txt, если
переменные из окружения берутся из другого места, то оставляем его пустым)
config configuration:
REDIS_ENABLED:T
REDIS_HOST:localhost
REDIS_PORT:6379
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
Ниже предоставлен простой код, который инициализирует необходимое подключение кладет и достает из него данные:
```
```go
package main

import (
	"context"
	"fmt"
	"log"

	core "github.com/KA1251/CoreModule"
)

func main() {
	var testCon core.ConnectionHandler
	core.Initiallizing(&testCon)
	ctx := context.Background()
	err := testCon.Redis.Set(ctx, "newkey2", "45", 0).Err()
	if err != nil {

		log.Fatal("qq1")
	}
	val, err := testCon.Redis.Get(ctx, "newkey2").Result()
	if err != nil {
		log.Fatal("qq")
	}

	fmt.Println(val)
	testCon.CloseAllConnections()
}
```

```text
Output:
time="2024-03-07T13:50:46+03:00" level=info msg="Redis sucsessfull conection"
time="2024-03-07T13:50:48+03:00" level=info msg="sucsessfull conection to sql"
45
```
## Example of usage 2
Dockerfile:
```text

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
docker-compose file:
```text
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
/*
package main

import (

	"connection_test/core"
	"context"
	"database/sql"
	"fmt"
	"log"

)

	type Product struct {
		ProductID   int     `db:"product_id" json:"product_id"`
		Naming      string  `db:"naming" json:"naming"`
		Weight      float64 `json:"weight"`
		Description string  `json:"description"`
	}

func main() {

	var test_connection core.ConnectionHandler

	core.Initiallizing(&test_connection)

	ctx := context.Background()

	err := test_connection.Redis.Set(ctx, "newkey2", "45", 0).Err()
	if err != nil {

		log.Fatal("qq1")
	}
	val, err := test_connection.Redis.Get(ctx, "newkey2").Result()
	if err != nil {
		log.Fatal("qq")
	}

	var p Product
	err = test_connection.SQLDB.Get(&p, "SELECT product_id, naming, weight, description FROM products WHERE product_id = $1", 1)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Нет данных для указанного id")
		} else {
			log.Fatalln("Ошибка при выполнении запроса:", err)
		}
	}

	if err != nil {
		fmt.Println("+")
		log.Fatal(err)
	}
	fmt.Println(p)
	fmt.Println(val)
	//core.StartHealthCheckServer("1251")

}
*/
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
