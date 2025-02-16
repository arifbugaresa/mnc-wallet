package session

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"strconv"
)

type RedisData struct {
	UserId      string `json:"user_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

var (
	RedisClient *redis.Client
)

func Initiator() *redis.Client {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("connection.redis.port" + ":" + strconv.Itoa(viper.GetInt("connection.redis.port"))),
		Password: viper.GetString("connection.redis.password"),
		DB:       viper.GetInt("connection.redis.db"),
	})

	return RedisClient
}
