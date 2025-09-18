package producer

import (
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
)

const channelName = "device_readings"

// PublishReading convierte la lectura a JSON y la publica en Redis
func PublishReading(client *redis.Client, reading DeviceReading) error {
	data, err := json.Marshal(reading)
	if err != nil {
		return err
	}

	if err := client.Publish(ctx, channelName, data).Err(); err != nil {
		return err
	}

	log.Printf("Publicado en Redis: %s", data)
	return nil
}
