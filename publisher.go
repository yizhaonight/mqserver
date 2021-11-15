package mqserver

import (
	"github.com/streadway/amqp"
	"kumact.com/gosdk/rest"
	"kumact.com/gosdk/xutils/logutil"
)

const (
	METHOD_GET    = "GET"
	METHOD_POST   = "POST"
	METHOD_PUT    = "PUT"
	METHOD_DELETE = "DELETE"
)

/*
 * Decide where to route the requests.
 * If it is a Get method, directly route it to datacenter backend;
 * if it is a method other than Get, put it into message queue.
 *
 * The <methods> parameter is something looks like "GET|POST|PUT|DELETE"
 */
func Publish(methods, path string) (err error) {
	_, err = GetIP()
	if err != nil {
		logutil.Error(err)
		return
	}
	err = Queue(methods, path)
	if err != nil {
		logutil.Error(err)
		return
	}
	return
}

/*
 * Put request into a message queue
 */
func Queue(methods, path string) (err error) {
	ip, err := GetIP() // Get the client ip to whom we publish our message
	if err != nil {
		logutil.Error(err)
		return
	}

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
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

	err = ch.ExchangeDeclare(
		"mps_subscribe_demo", // exchangeName
		"topic",              // kind
		true,                 // durable
		false,                // autoDelete
		false,                // internal
		false,                // noWait
		nil,                  // args
	)
	if err != nil {
		logutil.Error(err)
		return
	}

	req := RequestBody{
		Method: methods,
		Path:   path,
	}
	body := req.String()

	err = ch.Publish(
		"mps_subscribe_demo", // exchange name
		ip,                   // routing key
		false,                // mandatory
		false,                // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		},
	)

	logutil.Debug("[x] Sent " + body + " to " + ip)
	// logutil.Debug(body)
	return
}

/*
 * Get IP Address of edgeserver.
 */
func GetIP() (ip string, err error) {
	// api to be completed
	testIP := "1.1.1.1"
	return testIP, err
}

type RequestBody struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

func (p *RequestBody) String() string {
	return rest.JSONString(p)
}
