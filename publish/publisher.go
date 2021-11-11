package publisher

import (
	"strings"
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
 * The methods parameter is something looks like "GET|POST|PUT|DELETE"
 */
func Publish(methods, path string, ctl interface{}) (err error) {
	methodList := strings.Split(methods, "|")
	_ = GetIP()
	for _, method := range methodList {
		if method == METHOD_GET {
			return
		} else {
			err = Queue()
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
func Queue() (err error) {
	// ip := GetIP()
	// err = amqp.Dial("amqp://")
	return
}

func GetIP() string {
	return ""
}
