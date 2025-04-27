package handler

import (
	"encoding/json"
	"log"

	"github.com/bhushan-naruto/mqtt-message-processor/internal/infra/mqtt_app/models"
	"github.com/bhushan-naruto/mqtt-message-processor/internal/repo"
	"github.com/bhushan-naruto/mqtt-message-processor/internal/usecase"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func MoistureMonitorHandler(
	moistureHighAlertApiUrl string,
	moistureLowAlertApiUrl string,
	callFrom string, callTo string,
	storageRepo repo.StorageRepo,
	voiceCallRepo repo.VoiceCallRepo,
	client mqtt.Client,
	message mqtt.Message,
) {
	u := usecase.NewMoistureUseCase(
		storageRepo,
		voiceCallRepo,
		moistureHighAlertApiUrl,
		moistureLowAlertApiUrl,
		callFrom,
		callTo,
	)

	req := new(models.MoistureRequest)

	if err := json.Unmarshal(message.Payload(), req); err != nil {
		log.Println("invalid request json pyaload, Err: ", err.Error())
		return
	}

	u.MonitorMoistureUseCase(req.MoistureLevel)
}

func MoistureUpdateHandler(
	moistureHighAlertApiUrl string,
	moistureLowAlertApiUrl string,
	callFrom string, callTo string,
	storageRepo repo.StorageRepo,
	voiceCallRepo repo.VoiceCallRepo,
	client mqtt.Client,
	message mqtt.Message,
) {
	u := usecase.NewMoistureUseCase(
		storageRepo,
		voiceCallRepo,
		moistureHighAlertApiUrl,
		moistureLowAlertApiUrl,
		callFrom,
		callTo,
	)

	u.UpdateMoistureUseCase(message.Payload())
}
