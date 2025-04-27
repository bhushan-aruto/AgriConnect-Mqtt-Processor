package mqtt_connection

import (
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttConnection struct {
	Client mqtt.Client
}

func Connect(brokerHost, brokerPort string) *MqttConnection {
	brokerUrl := fmt.Sprintf("tcp://%s:%s", brokerHost, brokerPort)

	opts := mqtt.NewClientOptions()

	opts.AddBroker(brokerUrl)
	opts.SetClientID("agriconnect/message/processor")

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("error occurred while connecting to mqtt broker: %s, Err: %s\n", brokerUrl, token.Error().Error())
	}

	log.Printf("connected to mqtt broker: %s\n", brokerUrl)

	return &MqttConnection{
		Client: client,
	}

}
