package main

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// MQTTClient represents the MQTT client used for communication.
var MQTTClient mqtt.Client

func main() {
	InitializeMQTTClient()
	token := MQTTClient.Subscribe("/movement", byte(2), func(client mqtt.Client, message mqtt.Message) {
		fmt.Printf("Received message on topic %s: %s\n", message.Topic(), message.Payload())
	})
	token.Wait()

	Publish("forward")
}
