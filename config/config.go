package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PhoneNumber             string
	BrokerHost              string
	BrokerPort              string
	TwilioPhoneNumber       string
	MoistureHighAlertApiUrl string
	MoistureLowAlertApiUrl  string
}

func getEnv(key string) string {
	value := os.Getenv(key)

	if value == "" {
		log.Fatalf("missing or empty env variable -> %s\n", key)
	}

	return value
}

func LoadConfig() *Config {

	err := godotenv.Load()

	if err != nil {
		log.Fatalln(err)
	}

	return &Config{
		PhoneNumber:             getEnv("PHONE_NUMBER"),
		BrokerHost:              getEnv("BROKER_HOST"),
		BrokerPort:              getEnv("BROKER_PORT"),
		TwilioPhoneNumber:       getEnv("TWILIO_PHONE_NUMBER"),
		MoistureHighAlertApiUrl: getEnv("MOISTURE_HIGH_API_URL"),
		MoistureLowAlertApiUrl:  getEnv("MOISTURE_LOW_API_URL"),
	}
}
