package main

import (
	"github.com/arifbugaresa/mnc-wallet/config"
	"github.com/arifbugaresa/mnc-wallet/databases/connection"
	"github.com/arifbugaresa/mnc-wallet/databases/migration"
	"github.com/arifbugaresa/mnc-wallet/modules/health_check"
	"github.com/arifbugaresa/mnc-wallet/modules/master/transaction"
	"github.com/arifbugaresa/mnc-wallet/modules/master/user"
	"github.com/arifbugaresa/mnc-wallet/modules/upload"
	"github.com/arifbugaresa/mnc-wallet/utils/rabbitmq"
	redisPackage "github.com/arifbugaresa/mnc-wallet/utils/session"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func main() {
	config.Initiator()

	dbConnection, err := connection.Initiator()
	if err != nil {
		return
	}
	defer dbConnection.Close()

	migration.Initiator(dbConnection)
	redisConnection := redisPackage.Initiator()

	// initiate rabbitmq publisher
	rabbitMqConn := rabbitmq.Initiator()
	defer rabbitMqConn.Channel.Close()
	defer rabbitMqConn.Conn.Close()

	// initiate rabbitmq consumer
	//_ = rabbitMqConn.Consume()

	InitiateRouter(dbConnection, redisConnection, rabbitMqConn)
}

func InitiateRouter(dbConnection *sqlx.DB, redisConnection *redis.Client, rabbitMqConn *rabbitmq.RabbitMQ) {
	router := gin.Default()

	health_check.Initiator(router)

	// module user
	user.Initiator(router, dbConnection, redisConnection)
	transaction.Initiator(router, dbConnection, redisConnection, rabbitMqConn)

	upload.Initiator(router, dbConnection)

	router.Run(viper.GetString("app.port"))
}
