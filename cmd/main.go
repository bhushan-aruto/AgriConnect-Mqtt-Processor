package main

import (
	"github.com/bhushan-naruto/mqtt-message-processor/config"
	"github.com/bhushan-naruto/mqtt-message-processor/internal/infra/mqtt_app"
)

func main() {
	config := config.LoadConfig()
	mqtt_app.StartApp(config)
}
