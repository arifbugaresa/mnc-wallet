package transaction

import (
	"encoding/json"
	"fmt"
	"github.com/arifbugaresa/mnc-wallet/middlewares"
	"github.com/arifbugaresa/mnc-wallet/utils/common"
	"github.com/arifbugaresa/mnc-wallet/utils/constant/enum"
	"github.com/arifbugaresa/mnc-wallet/utils/rabbitmq"
	"github.com/arifbugaresa/mnc-wallet/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type Service interface {
	TopUp(ctx *gin.Context, dataBody TopUpRequest) (data TopUpResponse, err error)
	Transfer(ctx *gin.Context, dataBody TransferRequest, rabbitMqConn *rabbitmq.RabbitMQ) (data TransferResponse, err error)
}

type UserService struct {
	repo  *UserRepository
	redis *redis.Client
}

func NewService(repo *UserRepository, redis *redis.Client) *UserService {
	return &UserService{
		repo:  repo,
		redis: redis,
	}
}

func (s UserService) TopUp(ctx *gin.Context, dataBody TopUpRequest) (result TopUpResponse, err error) {
	auth, err := middlewares.GetSession(ctx)
	if err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
	}

	data := TopUpModel{
		UserId:       auth.UserId,
		Amount:       dataBody.Amount,
		DefaultTable: common.DefaultTable{}.GetDefaultTable(ctx),
	}

	output, err := s.repo.TopUp(ctx, data)
	if err != nil {
		return
	}

	result = TopUpResponse{
		TopUpId:       output.TopUpId,
		Amount:        output.Amount,
		BalanceBefore: output.BalanceBefore,
		BalanceAfter:  output.BalanceAfter,
		CreatedDate:   output.CreatedAt,
	}

	return
}

func (s UserService) Transfer(ctx *gin.Context, dataBody TransferRequest, rabbitMqConn *rabbitmq.RabbitMQ) (result TransferResponse, err error) {
	auth, err := middlewares.GetSession(ctx)
	if err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
	}

	data := TransferModel{
		UserId:        auth.UserId,
		UserIdTarget:  dataBody.TargetUser,
		Amount:        dataBody.Amount,
		Remarks:       dataBody.Remarks,
		BalanceBefore: 0,
		BalanceAfter:  0,
		DefaultTable:  common.DefaultTable{}.GetDefaultTable(ctx),
	}

	// will be process background on queue rabbit mq
	messageByte, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = rabbitMqConn.Publish(rabbitmq.MqConfig{
		QueueName: enum.TransferQueue,
		Messsage:  string(messageByte),
	})
	if err != nil {
		return
	}

	return
}
