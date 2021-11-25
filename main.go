package main

import (
	"os"

	"kumact.com/dc_scheduler/service"
)

func main() {
	addr := os.Getenv("SERVER_ADDR")
	if addr == "" {
		return
	}
	cfg := service.GetConfig()
	cfg.Addr = addr
	service.Run(cfg.Addr)
}
