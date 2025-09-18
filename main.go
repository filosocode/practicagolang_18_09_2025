package main

import (
	"time"

	"github.com/tu_usuario/enerbit_golang/producer"
)

func main() {
	client := producer.NewRedisClient()
	for {
		reading := producer.SimulateReading("device-001")
		if err := producer.PublishReading(client, reading); err != nil {
			panic(err)
		}
		time.Sleep(2 * time.Second)
	}
}
