package main

import (
	"encoding/json"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Pessoa struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func onMessageReceived(client mqtt.Client, msg mqtt.Message) {
	var pessoa Pessoa
	err := json.Unmarshal(msg.Payload(), &pessoa)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %v\n", err)
		return
	}

	fmt.Printf("Person received: Name: %s, Age: %d\n", pessoa.Name, pessoa.Age)
}

func main() {
	// mqtt topic
	topic := "test/topic"
	// mqtt broker
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://localhost:1883")
	opts.SetClientID("go_mqtt_client_subscriber")

	// create a new MQTT client
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(fmt.Sprintf("Error connecting to broker: %v", token.Error()))
	}

	if token := client.Subscribe(topic, 0, onMessageReceived); token.Wait() && token.Error() != nil {
		panic(fmt.Sprintf("Error subscribing to topic: %v", token.Error()))
	}

	fmt.Printf("Subscribed to topic: %s\n", topic)

	// Keep the program running
	select {}
}
