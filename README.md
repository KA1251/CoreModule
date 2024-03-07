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
## Example of usage
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
