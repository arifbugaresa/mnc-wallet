package rabbitmq

import (
	"github.com/arifbugaresa/mnc-wallet/utils/constant/enum"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

func (r *RabbitMQ) Publish(rabbitConfig MqConfig) (err error) {
	var (
		routingKey   = viper.GetString("app.mode")
		exchangeName = viper.GetString("name")
	)

	rabbitConfig.ExchangeName = exchangeName
	rabbitConfig.RoutingKey = routingKey

	if rabbitConfig.RoutingKey == "" {
		rabbitConfig.QueueName = enum.DefaultQueue
	}

	_ = r.DeclareExchange(rabbitConfig)
	_ = r.DeclareQueue(rabbitConfig)
	_ = r.Bind(rabbitConfig)

	// publishing a message
	err = r.Channel.Publish(
		rabbitConfig.ExchangeName, // exchange name
		rabbitConfig.RoutingKey,   // key
		false,                     // mandatory
		false,                     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(rabbitConfig.Messsage),
		},
	)

	if err != nil {
		panic(err)
	}

	return
}
