package main

import (
	"flag"

	"github.com/gigmsn/subscriber/consumer"
)

func main() {

	addrPtr := flag.String("addr", "amqp://guest:guest@broker:5672", "queue address")
	queuePrt := flag.String("queue", "gigmsn", "queue name")
	flag.Parse()

	consumer.Subscribe(*addrPtr, *queuePrt)
}
