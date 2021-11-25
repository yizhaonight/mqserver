package service

import (
	"encoding/json"

	"github.com/streadway/amqp"
	"kumact.com/dc_scheduler/utils"
	"kumact.com/gosdk/xutils/logutil"
)

func ClientSubscribe() (err error) {
	return DoClientSubscribe()
}

func DoClientSubscribe() (err error) {
	acc, err := utils.GetAccount()
	if err != nil {
		logutil.Error(err)
		return
	}
	host, err := utils.GetServerIP()
	if err != nil {
		logutil.Error(err)
		return
	}
	port, err := utils.GetServerPort()
	if err != nil {
		logutil.Error(err)
		return
	}
	client, _ := utils.GetClientIP()
	// if err != nil {
	// 	logutil.Error(err)
	// 	return
	// }
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
	if err != nil {
		logutil.Error(err)
		return
	}

	q, err := ch.QueueDeclare(
		"",
		false, // durable
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		logutil.Error(err)
		return
	}

	// clientIP, _ := utils.GetClientIP()
	err = ch.QueueBind(
		q.Name,
		client,
		exchange,
		false,
		nil,
	)

	msgs, err := ch.Consume(
		q.Name,
		"",
		true, //auto acknowledgement
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logutil.Error(err)
		return
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			var msg RequestBody
			_ = json.Unmarshal(d.Body, &msg)
			//wr, r, token := GetOriginalRequest(msg.MsgID)
		}
	}()
	<-forever
	return
}
