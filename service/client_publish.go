package service

import (
	"github.com/streadway/amqp"
	"kumact.com/dc_scheduler/utils"
	"kumact.com/gosdk/xutils/logutil"
)

func ClientPublish(response ClientResponse) (err error) {
	return DoClientPublish(response)
}

func DoClientPublish(response ClientResponse) (err error) {
	acc, _ := utils.GetAccount()
	host, _ := utils.GetServerIP()
	port, _ := utils.GetServerPort()
	url, _ := utils.GetURL(acc, host, port)

	conn, err := amqp.Dial(url)
	if err != nil {
		logutil.Error(err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		logutil.Error(err)
		return
	}
	defer ch.Close()

	exchange := utils.GetExchangeName()
	err = ch.ExchangeDeclare(
		exchange, // exchangeName
		"topic",  // kind
		true,     // durable
		false,    // autoDelete
		false,    // internal
		false,    // noWait
		nil,      // args
	)

	ch.Publish(
		exchange,
		"server",
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(response.String()),
		},
	)
	logutil.Info("[x] Client sent", response, "to server")

	return
}
