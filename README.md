usage:

create a config.txt file in your working directory with such format:

REDIS_ENABLED:

REDIS_HOST:localhost

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

use params which you need (in ENABLED params use "T" to make them enabled)

than use:

var YourCon core.ConnectionHandler

core.Initiallizing(&YourCon)

you'll get logs about all your connections in stdout
