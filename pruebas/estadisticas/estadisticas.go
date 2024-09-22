package estadisticas

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/stat/distuv"
)

func response(b bool) string {
	if b {
		return "Sí"
	}
	return "No"
}

func calcMedia(data []float64) float64 {
	media := 0.0
	n := float64(len(data))
	for _, n := range data {
		media += n
	}
	media /= n
	return media
}

func MediaTest(data []float64) {
	media := calcMedia(data)
	n := float64(len(data))
	fmt.Println("\nMedia de los datos:", media)
	li := 0.5 - (1.96)/math.Sqrt(12*n)
	ls := 0.5 + (1.96)/math.Sqrt(12*n)
	fmt.Printf("Limites: [%f, %f]\n", li, ls)

	fmt.Println("Ho = El conjunto tiene un media de 0.5")
	Ho := li <= media && media <= ls
	fmt.Printf("¿Se cumple Ho? Respuesta: %s\n", response(Ho))
}

func VarianzaTest(data []float64) {
	n := float64(len(data))
	media := calcMedia(data)
	varianza := func() float64 {
		sum := 0.0
		for _, v := range data {
			sum += math.Pow((v - media), 2)
		}
		return sum / (n - 1)
	}()

	fmt.Println("Varianza:", varianza)
	ls := distuv.ChiSquared{K: float64(n - 1)}.Quantile(1-0.05/2) / (12 * (n - 1))
	li := distuv.ChiSquared{K: float64(n - 1)}.Quantile(0.05/2) / (12 * (n - 1))
	fmt.Printf("Limites: [%.5f, %.5f]\n", li, ls)
	fmt.Println("Ho = El conjunto tiene un varianza de 1/12")
	Ho := li <= varianza && varianza <= ls
	fmt.Printf("¿Se cumple Ho? Respuesta: %s\n", response(Ho))
}
