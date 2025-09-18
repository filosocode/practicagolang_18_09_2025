package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/tu_usuario/enerbit_golang/producer"
)

const channelName = "device_readings"

var ctx = context.Background()

// StartConsumer se suscribe al canal y procesa cada mensaje
func StartConsumer(client *redis.Client) {
	sub := client.Subscribe(ctx, channelName)
	ch := sub.Channel()

	log.Println("Consumidor suscrito a", channelName)

	for msg := range ch {
		var reading producer.DeviceReading
		if err := json.Unmarshal([]byte(msg.Payload), &reading); err != nil {
			log.Println("Error decodificando:", err)
			continue
		}
		fmt.Printf("Recibido: %+v\n", reading)
	}
}
