package intervalos

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"pruebas/archivos"

	"gonum.org/v1/gonum/stat/distuv"
)

type intervalo struct {
	n     int
	group []float64
}

func IntervalosTest(data []float64, n int) {
	filename := archivos.Filename() + "_intervalos_output.txt"
	file, _ := os.Create(filename)
	writer := bufio.NewWriter(file)

	intervalo_value := 1 / float64(n)
	intervalos := make(map[float64]intervalo)
	for i := intervalo_value; i < 1; i += intervalo_value {
		intervalos[i] = intervalo{n: 0, group: make([]float64, 0, 10)}
	}

	if len(intervalos) == n-1 {
		intervalos[1] = intervalo{n: 0, group: make([]float64, 0, 10)}
	}

	for _, d := range data {
		var possible_range float64
		possible_range = 1
		for k := range intervalos {
			if d < k && k < possible_range {
				possible_range = k
			}
		}
		i := intervalos[possible_range]
		i.n += 1
		i.group = append(i.group, d)
		intervalos[possible_range] = i
	}

	N := len(data) / n
	fmt.Println("Cantidad esperada por sub grupos:", N)
	writer.WriteString(fmt.Sprintf("Cantidad esperada por subgrupos: %d\n", N))
	fmt.Println("Subgrupo\tN° Elementos")
	writer.WriteString("Subgrupo\tN° Elementos\n")
	for k, v := range intervalos {
		show := fmt.Sprintf("%f\t%d\n", k, v.n)
		fmt.Print(show)
		writer.WriteString(show)
	}

	fmt.Println("\nFormula:")
	writer.WriteString("\nFormula:\n")
	show := fmt.Sprintf("X²o = (1/%d)*{", N)
	fmt.Print(show)
	writer.WriteString(show)
	var sum float64
	for _, v := range intervalos {
		show := fmt.Sprintf("(%d - %d)² + ", v.n, N)
		fmt.Print(show)
		writer.WriteString(show)
		sum += math.Pow(float64(v.n-N), 2)
	}
	X2o := 1 / float64(N) * sum
	show = fmt.Sprintf(" } = %f\n", X2o)
	fmt.Print(show)
	writer.WriteString(show)
	X2 := distuv.ChiSquared{K: float64(n - 1)}.Quantile(0.95)
	show = fmt.Sprintf("X²[0.05,%d] = %f\n", n-1, X2)
	fmt.Print(show)
	writer.WriteString(show)

	if X2o < X2 {
		fmt.Println("\nEL CONJUNTO DE NÚMEROS ESTÁ DISTRIBUIDO UNIFORMEMENTE")
		writer.WriteString("\nEL CONJUNTO DE NÚMEROS ESTÁ DISTRIBUIDO UNIFORMEMENTE")
	} else {
		fmt.Println("\nEL CONJUNTO DE NÚMEROS NO ESTÁ DISTRIBUIDO UNIFORMEMENTE")
		writer.WriteString("\nEL CONJUNTO DE NÚMEROS NO ESTÁ DISTRIBUIDO UNIFORMEMENTE")
	}

	writer.Flush()
	fmt.Println("La salida fue almacenada en el archivo:", filename)
}
