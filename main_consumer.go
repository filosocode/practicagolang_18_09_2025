package main

import (
	"github.com/tu_usuario/enerbit_golang/consumer"
	"github.com/tu_usuario/enerbit_golang/producer"
)

func main() {
	client := producer.NewRedisClient()
	consumer.StartConsumer(client)
}
