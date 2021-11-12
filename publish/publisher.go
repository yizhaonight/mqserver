package publisher

import (
	"log"
	"strings"

	"github.com/streadway/amqp"
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
func Publish(methods, path string, ctl interface{}) (err error) {
	methodList := strings.Split(methods, "|")
	_, err = GetIP()
	for _, method := range methodList {
		if method == METHOD_GET {
			return
		} else {
			err = Queue(path, ctl)
			if err != nil {
				return
			}
		}
	}
	return
}

/*
 * Put request into a message queue
 */
func Queue(path string, ctl interface{}) (err error) {
	// ip, err := GetIP() // Get the client ip to whom we publish our message
	if err != nil {
		log.Fatal(err)
		return
	}
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("%s: %s", "Cannot establish connection", err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("%s: %s", "Cannot build channel", err)
		return
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"mps_subscribe", // exchangeName
		"topic",         // kind
		true,            // durable
		false,           // autoDelete
		false,           // internal
		false,           // noWait
		nil,             // args
	)
	if err != nil {
		log.Fatalf("%s: %s", "Cannot declare exchange", err)
		return
	}

	return
}

func GetIP() (string, err error) {
	return
}
