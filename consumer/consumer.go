package consumer

import (
	"fmt"

	"github.com/gigmsn/messagebroker"
	log "github.com/sirupsen/logrus"
)

// Subscribe RabbitMQ topic to get messages
func Subscribe(addr, queue string) {
	broker, err := messagebroker.New("amqp://guest:guest@broker:5672", "gigmsn")
	if err != nil {
		log.Fatalf("could not create broker: %s", err)
	}
	defer broker.Close()
	log.Infoln("broker connection established")

	msgs, _ := broker.Channel.Consume(
		queue, // queue name
		"",    // consumer,
		true,  // autoAck,
		false, // exclusive,
		false, //noLocal,
		false, //noWait,
		nil,   // args amqp.Table
	)

	for m := range msgs {
		fmt.Println(string(m.Body))
	}
}
