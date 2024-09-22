package estadisticas

import (
	"fmt"
	"math"
)

func response(b bool) string {
	if b {
		return "Sí"
	}
	return "No"
}

func MediaTest(data []float64) {
	media := 0.0
	n := float64(len(data))
	for _, n := range data {
		media += n
	}
	media /= n
	fmt.Println("\nMedia de los datos:", media)
	li := 0.5 - (1.96)/math.Sqrt(12*n)
	ls := 0.5 + (1.96)/math.Sqrt(12*n)
	fmt.Printf("Limites: [%f, %f]\n", li, ls)

	fmt.Println("Ho = El conjunto tiene un media de 0.5")
	Ho := li <= media && media <= ls
	fmt.Printf("¿Se cumple Ho? Respuesta: %s\n", response(Ho))
}
