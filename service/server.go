package service

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"gitea.com/lunny/tango"
	"github.com/streadway/amqp"
	"kumact.com/dc_scheduler/utils"
	"kumact.com/gosdk/rest"
	"kumact.com/gosdk/xutils/errutil"
	"kumact.com/gosdk/xutils/logutil"
	"kumact.com/gosdk/xutils/uuid"
)

func ServerSubscribe(method string, path *url.URL, token string, ctx *tango.Context) (err error) {
	err = DoServerSubscribe(method, path, token, ctx)
	if err != nil {
		logutil.Error(err)
		return
	}
	return
}

func DoServerSubscribe(method string, path *url.URL, token string, ctx *tango.Context) (err error) {
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

	ch2, err := conn.Channel()
	if err != nil {
		logutil.Error(err)
		return
	}
	defer ch2.Close()

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

	err = ch2.ExchangeDeclare(
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

	client, _ := utils.GetClientIP()
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

	// ServerPublish(method, path, token)
	msgID := uuid.GetUUID()

	req := RequestBody{
		Method: method,
		URL:    path,
		Token:  token,
		Time:   time.Now().UnixMilli(),
		IP:     client,
		MsgID:  msgID,
	}
	body := req.String()

	err = ch2.Publish(
		exchange, // exchange name
		req.IP,   // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		},
	)
	if err != nil {
		logutil.Error(err)
	}
	logutil.Info("[x] Server Sent", body, "to", req.IP)

	logutil.Info("[x] Waiting for response")
	go func() {
		for d := range msgs {
			var msg ClientResponse
			_ = json.Unmarshal(d.Body, &msg)
			logutil.Info("[x] Received response from client:", msg.String())
			if !HasID(msg.MsgID) {
				err = errutil.NewError(ID_NOT_FOUND, "message id not found")
				break
			} else {
				baseURL := "http://127.0.0.1:8081"
				client, _ := rest.NewClientWithProxy(
					rest.NewClient(baseURL, nil),
					"token=",
				)
				ctx.Req().URL.RawQuery += "from=server"
				client.Proxy(
					ctx.ResponseWriter,
					ctx.Req(),
					token,
				)
				break
			}
		}
	}()
	return
}

const ID_NOT_FOUND = 12345

func HasID(msgID string) bool {
	return true
}

func GetOriginalRequest(msgID string) (
	wr http.ResponseWriter, r *http.Request, token string) {
	return
}
