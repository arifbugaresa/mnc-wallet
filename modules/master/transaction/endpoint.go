package transaction

import (
	"github.com/arifbugaresa/mnc-wallet/middlewares"
	"github.com/arifbugaresa/mnc-wallet/utils/rabbitmq"
	"github.com/arifbugaresa/mnc-wallet/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

func Initiator(router *gin.Engine, dbConnection *sqlx.DB, redisConnection *redis.Client, rabbitMqConn *rabbitmq.RabbitMQ) {
	var (
		userRepo = NewRepository(dbConnection)
		userSrv  = NewService(userRepo, redisConnection)
	)

	protectedApi := router.Group("/api/users")
	protectedApi.Use(middlewares.JwtMiddleware())
	{
		protectedApi.POST("/top-up", func(c *gin.Context) {
			TopUpEndpoint(c, userSrv)
		})
		protectedApi.POST("/transfer", func(c *gin.Context) {
			TransferEndpoint(c, userSrv, rabbitMqConn)
		})
	}
}

func TopUpEndpoint(ctx *gin.Context, userSrv Service) {
	var (
		dataBody TopUpRequest
	)

	if err := ctx.ShouldBindJSON(&dataBody); err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	err := dataBody.ValidateTopUpRequest()
	if err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	result, err := userSrv.TopUp(ctx, dataBody)
	if err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	response.GenerateSuccessResponseWithData(ctx, "top up successful", result)
}

func TransferEndpoint(ctx *gin.Context, userSrv Service, rabbitMqConn *rabbitmq.RabbitMQ) {
	var (
		dataBody TransferRequest
	)

	if err := ctx.ShouldBindJSON(&dataBody); err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	err := dataBody.ValidateTransferRequest()
	if err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	_, err = userSrv.Transfer(ctx, dataBody, rabbitMqConn)
	if err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	response.GenerateSuccessResponse(ctx, "transfer will be process in background, have fun!")
}
