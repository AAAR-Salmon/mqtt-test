package main

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var options *mqtt.ClientOptions = mqtt.NewClientOptions().AddBroker("tcp://mqtt-broker:1883")
var msgCh = make(chan mqtt.Message)

func onReceive(client mqtt.Client, msg mqtt.Message) {
	msgCh <- msg
}

func publish(client mqtt.Client) {
	if token := client.Publish("mqtt-test", 0, false, "a boring message"); token.Wait() && token.Error() != nil {
		log.Fatalln(token.Error())
	}
}

func main() {
	client := mqtt.NewClient(options)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalln(token.Error())
	}

	if token := client.Subscribe("mqtt-test", 0, onReceive); token.Wait() && token.Error() != nil {
		log.Fatalln(token.Error())
	}

	go func() {
		for {
			publish(client)
			time.Sleep(time.Duration(1) * time.Second)
		}
	}()

	for {
		m := <-msgCh
		fmt.Printf("topic: %v, payload: %v\n", m.Topic(), string(m.Payload()))
	}
}
