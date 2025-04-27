package topics

import (
	"log"

	"github.com/bhushan-naruto/mqtt-message-processor/config"
	mqtt_connection "github.com/bhushan-naruto/mqtt-message-processor/internal/infra/mqtt_app/connection"
	"github.com/bhushan-naruto/mqtt-message-processor/internal/infra/mqtt_app/handler"
	"github.com/bhushan-naruto/mqtt-message-processor/internal/repo"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func InitTopicsSub(
	cfg *config.Config,
	mqttConn *mqtt_connection.MqttConnection,
	storageRepo repo.StorageRepo,
	voiceCallRepo repo.VoiceCallRepo,
) {
	if token := mqttConn.Client.Subscribe("agriconnect/moisture", 1, func(c mqtt.Client, m mqtt.Message) {
		go handler.MoistureMonitorHandler(
			cfg.MoistureHighAlertApiUrl,
			cfg.MoistureLowAlertApiUrl,
			cfg.TwilioPhoneNumber,
			cfg.PhoneNumber,
			storageRepo,
			voiceCallRepo,
			c,
			m,
		)
	}); token.Wait() && token.Error() != nil {
		log.Println("error occurred while subscribing to the topic: agriconnect/moisture, Err:  ", token.Error().Error())
		return
	}

	if token := mqttConn.Client.Subscribe("agriconnect/update/moisture", 1, func(c mqtt.Client, m mqtt.Message) {
		go handler.MoistureUpdateHandler(
			cfg.MoistureHighAlertApiUrl,
			cfg.MoistureLowAlertApiUrl,
			cfg.TwilioPhoneNumber,
			cfg.PhoneNumber,
			storageRepo,
			voiceCallRepo,
			c,
			m,
		)
	}); token.Wait() && token.Error() != nil {
		log.Println("error occurred while subscribing to the topic: agriconnect/update/moisture, Err:  ", token.Error().Error())
		return
	}
}
