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

func main() {
	opts := mqtt.NewClientOptions()

	opts.AddBroker("tcp://localhost:1883")
	opts.SetClientID("go_mqtt_client123")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(fmt.Sprintf("Error connecting to broker: %v", token.Error()))
	}

	people := Pessoa{
		Name: "John Doe",
		Age:  30,
	}

	payLoad, err := json.Marshal(people)
	if err != nil {
		panic(fmt.Sprintf("Error marshaling JSON: %v", err))
	}

	// Publish a message
	token := client.Publish("test/topic", 0, false, payLoad)

	token.Wait()

	fmt.Printf("Published message: %s to topic: %s\n", payLoad, "test/topic")

	client.Disconnect(250)
}
