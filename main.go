package main

import (
	"github.com/arifbugaresa/mnc-wallet/config"
	"github.com/arifbugaresa/mnc-wallet/databases/connection"
	"github.com/arifbugaresa/mnc-wallet/databases/migration"
	"github.com/arifbugaresa/mnc-wallet/modules/health_check"
	"github.com/arifbugaresa/mnc-wallet/modules/master/user"
	"github.com/arifbugaresa/mnc-wallet/modules/upload"
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

	InitiateRouter(dbConnection, redisConnection)
}

func InitiateRouter(dbConnection *sqlx.DB, redisConnection *redis.Client) {
	router := gin.Default()

	health_check.Initiator(router)

	// module user
	user.Initiator(router, dbConnection, redisConnection)
	upload.Initiator(router, dbConnection)

	router.Run(viper.GetString("app.port"))
}
