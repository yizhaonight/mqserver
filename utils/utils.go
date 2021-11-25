package utils

import (
	"bytes"
	"strconv"
)

type Account struct {
	Name     string `json:"account_name"`
	Password string `json:"account_password"`
}

func GetAccount() (acc Account, err error) {
	// temporarily set to default account
	acc = Account{
		Name:     "test",
		Password: "test",
	}
	return
}

func GetServerIP() (ip string, err error) {
	// temporary
	ip = "127.0.0.1"
	return
}

func GetServerPort() (port int, err error) {
	// temporarily set to default port
	port = 5672
	return
}

func GetClientIP() (ip string, err error) {
	// temporary
	testIP := "192.168.1.253"
	return testIP, err
}

func GetExchangeName() string {
	// temporary
	name := "mps_subscribe_demo"
	return name
}

func GetURL(acc Account, host string, port int) (url string, err error) {
	var buffer bytes.Buffer
	buffer.WriteString("amqp://")
	buffer.WriteString(acc.Name)
	buffer.WriteString(":")
	buffer.WriteString(acc.Password)
	buffer.WriteString("@")
	buffer.WriteString(host)
	buffer.WriteString(":")
	buffer.WriteString(strconv.Itoa(port))
	url = buffer.String()
	return
}
