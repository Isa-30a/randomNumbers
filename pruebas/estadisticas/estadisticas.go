package estadisticas

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"pruebas/archivos"

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
	filename := archivos.Filename() + "_media_output.txt"
	file, _ := os.Create(filename)
	writer := bufio.NewWriter(file)

	media := calcMedia(data)
	n := float64(len(data))
	fmt.Println("\nMedia de los datos:", media)
	writer.WriteString(fmt.Sprintf("Media de los datos: %v\n", media))
	li := 0.5 - (1.96)/math.Sqrt(12*n)
	ls := 0.5 + (1.96)/math.Sqrt(12*n)
	show := fmt.Sprintf("Limites: [%f, %f]\n", li, ls)
	fmt.Print(show)
	writer.WriteString(show)

	fmt.Println("Ho = El conjunto tiene un media de 0.5")
	writer.WriteString("Ho = El conjunto tiene un media de 0.5\n")
	Ho := li <= media && media <= ls
	show = fmt.Sprintf("¿Se cumple Ho? Respuesta: %s\n", response(Ho))
	fmt.Print(show)
	writer.WriteString(show)

	writer.Flush()
	fmt.Println("\nSalida almacenada en el archivo:", filename)
}

func VarianzaTest(data []float64) {
	filename := archivos.Filename() + "_varianza_output.txt"
	file, _ := os.Create(filename)
	writer := bufio.NewWriter(file)

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
	writer.WriteString(fmt.Sprintf("Varianza: %v\n", varianza))
	ls := distuv.ChiSquared{K: float64(n - 1)}.Quantile(1-0.05/2) / (12 * (n - 1))
	li := distuv.ChiSquared{K: float64(n - 1)}.Quantile(0.05/2) / (12 * (n - 1))
	show := fmt.Sprintf("Limites: [%.5f, %.5f]\n", li, ls)
	fmt.Print(show)
	writer.WriteString(show)
	fmt.Println("Ho = El conjunto tiene un varianza de 1/12")
	writer.WriteString("Ho = El conjunto tiene un varianza de 1/12\n")
	Ho := li <= varianza && varianza <= ls
	show = fmt.Sprintf("¿Se cumple Ho? Respuesta: %s\n", response(Ho))
	fmt.Print(show)
	writer.WriteString(show)
	fmt.Println("Salida almacenada en el archivo:", filename)

	writer.Flush()
}
