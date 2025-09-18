package database

import (
	"log"

	"github.com/tu_usuario/enerbit_golang/models"
)

func InsertReading(r models.Reading) {
	query := `INSERT INTO readings (device_id, timestamp, consumption) VALUES ($1, $2, $3)`
	_, err := DB.Exec(query, r.DeviceID, r.Timestamp, r.Consumption)
	if err != nil {
		log.Println("Error insertando lectura:", err)
	}
}
