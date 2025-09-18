package consumer

import (
	"fmt"

	"github.com/tu_usuario/enerbit_golang/producer"
)

// HandleReading procesa cada lectura recibida del canal de Redis.
// Por ahora solo la imprime, luego la extenderemos para guardarla en PostgreSQL.
func HandleReading(reading producer.DeviceReading) {
	fmt.Printf("[EVENT] Procesando lectura: %+v\n", reading)
}
