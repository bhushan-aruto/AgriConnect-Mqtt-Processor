package usecase

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/bhushan-naruto/mqtt-message-processor/internal/entity"
	"github.com/bhushan-naruto/mqtt-message-processor/internal/repo"
)

type moistureUseCase struct {
	moistureHighAlertApiUrl string
	moistureLowAlertApiUrl  string
	callFrom                string
	callTo                  string
	storageRepo             repo.StorageRepo
	voiceCallRepo           repo.VoiceCallRepo
}

func NewMoistureUseCase(
	stRepo repo.StorageRepo,
	vcRepo repo.VoiceCallRepo,
	moistureHighAlertApiUrl string,
	moistureLowAlertApiUrl string,
	callFrom string,
	callTo string,
) *moistureUseCase {
	return &moistureUseCase{
		moistureHighAlertApiUrl: moistureHighAlertApiUrl,
		moistureLowAlertApiUrl:  moistureLowAlertApiUrl,
		callFrom:                callFrom,
		callTo:                  callTo,
		storageRepo:             stRepo,
		voiceCallRepo:           vcRepo,
	}
}

func (u *moistureUseCase) MonitorMoistureUseCase(presentMoistureLevel uint32) {

	res, err := u.storageRepo.GetMoistureSettings()

	if err != nil {
		log.Println("error occurred with the storage while getting the moisture settings, Err: ", err.Error())
		return
	}

	settings := new(entity.MoistureSettings)

	if err := json.Unmarshal(res, settings); err != nil {
		log.Println("error occurred while decoding the moisture settings, Err: ", err.Error())
		return
	}

	if presentMoistureLevel >= settings.Max {
		url := fmt.Sprintf("%s?level=%s", u.moistureHighAlertApiUrl, strconv.Itoa(int(presentMoistureLevel)))
		if err := u.voiceCallRepo.MakeAlertCall(url, u.callFrom, u.callTo); err != nil {
			log.Println("error occurred while making the voice call, Err: ", err.Error())
		}
		return
	}

	if presentMoistureLevel <= settings.Min {
		url := fmt.Sprintf("%s?level=%s", u.moistureHighAlertApiUrl, strconv.Itoa(int(presentMoistureLevel)))
		if err := u.voiceCallRepo.MakeAlertCall(url, u.callFrom, u.callTo); err != nil {
			log.Println("error occurred while making the voice call, Err: ", err.Error())
		}
		return
	}

}

func (u *moistureUseCase) UpdateMoistureUseCase(payload []byte) {
	if err := u.storageRepo.UpdateMoistureSettings(payload); err != nil {
		log.Println("error occurred while updating the payload, Err: ", err.Error())
		return
	}
}
