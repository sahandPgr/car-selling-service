package main

import (
	"github.com/sahandPgr/car-selling-service/api"
	"github.com/sahandPgr/car-selling-service/config"
)

func main() {
	config := config.GetConfig()
	api.InitServer(config)
}
