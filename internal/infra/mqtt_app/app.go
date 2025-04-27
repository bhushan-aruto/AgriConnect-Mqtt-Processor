package mqtt_app

import (
	"log"
	"time"

	"github.com/bhushan-naruto/mqtt-message-processor/config"
	mqtt_connection "github.com/bhushan-naruto/mqtt-message-processor/internal/infra/mqtt_app/connection"
	"github.com/bhushan-naruto/mqtt-message-processor/internal/infra/mqtt_app/topics"
	"github.com/bhushan-naruto/mqtt-message-processor/internal/infra/storage"
	"github.com/bhushan-naruto/mqtt-message-processor/internal/infra/twilio_app"
)

func StartApp(conf *config.Config) {
	mqttConn := mqtt_connection.Connect(
		conf.BrokerHost,
		conf.BrokerPort,
	)

	storageRepo := storage.NewStorageRepo("./settings")

	storageRepo.Init()

	twilioClient := twilio_app.NewTwilioClient()

	voiceCallRepo := twilio_app.NewTwilioRepo(twilioClient)

	topics.InitTopicsSub(
		conf,
		mqttConn,
		storageRepo,
		voiceCallRepo,
	)

	for {
		if !mqttConn.Client.IsConnected() {
			log.Println("mqtt client is disconnected from the broker")
			if token := mqttConn.Client.Connect(); token.Wait() && token.Error() != nil {
				log.Println("error occurred while connecting to broker, Err: ", token.Error().Error())
				log.Println("trying to reconnect in 5 secs...")
				time.Sleep(5 * time.Second)
				continue
			}

			log.Println("reconnected to mqtt broker")
			topics.InitTopicsSub(
				conf,
				mqttConn,
				storageRepo,
				voiceCallRepo,
			)

		}
		time.Sleep(500 * time.Millisecond)
	}
}
