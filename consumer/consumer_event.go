package consumer

import (
	"github.com/tu_usuario/enerbit_golang/database"
	"github.com/tu_usuario/enerbit_golang/models"
	"github.com/tu_usuario/enerbit_golang/producer"
)

func HandleReading(reading producer.DeviceReading) {
	r := models.Reading{
		DeviceID:    reading.DeviceID,
		Timestamp:   reading.Timestamp,
		Consumption: reading.Consumption,
	}
	database.InsertReading(r)
}
