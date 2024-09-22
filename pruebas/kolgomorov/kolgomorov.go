package kolgomorov

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"pruebas/archivos"
	"slices"
)

func KolgomorovTest(data []float64) {
	filename := archivos.Filename() + "_intervalos_output.txt"
	file, _ := os.Create(filename)
	writer := bufio.NewWriter(file)

	clonedData := slices.Clone(data)
	slices.SortFunc(clonedData, func(a, b float64) int {
		if a < b {
			return -1
		} else if b > a {
			return 1
		}
		return 0
	})
	fmt.Println("Datos ordenados:", clonedData)
	writer.WriteString(fmt.Sprintf("Datos ordenados: %v\n", clonedData))
	n := len(clonedData)
	divsArrays := make([]float64, n)
	for i := range n {
		divsArrays[i] = float64(i+1) / float64(n)
	}
	max := -1.0
	for i := range n {
		show := fmt.Sprintf("[%.5f - %.5f] = %.5f\n", divsArrays[i], clonedData[i], math.Abs(divsArrays[i]-clonedData[i]))
		fmt.Print(show)
		writer.WriteString(show)
		divsArrays[i] = math.Abs(divsArrays[i] - clonedData[i])
		if divsArrays[i] > max {
			max = divsArrays[i]
		}
	}
	fmt.Println()
	writer.WriteString("\n")

	fmt.Println("Máximo valor:", max)
	writer.WriteString(fmt.Sprintf("Máximo valor: %v\n", max))
	ks := math.Sqrt(-math.Log(0.05/2) / (2 * float64(n)))
	fmt.Println("KS:", ks)
	writer.WriteString(fmt.Sprintf("KS: %v\n", ks))
	show := fmt.Sprintf("Dmax < Dks = %v", max < ks)
	fmt.Print(show)
	writer.WriteString(show)
	if max < ks {
		fmt.Println("\nLOS NÚMEROS TIENEN UNA DISTRIBUCIÓN UNIFORME")
		writer.WriteString("\nLOS NÚMEROS TIENEN UNA DISTRIBUCIÓN UNIFORME")
	} else {
		fmt.Println("\nLOS NÚMEROS NO TIENEN UNA DISTRIBUCIÓN UNIFORME")
		writer.WriteString("\nLOS NÚMEROS NO TIENEN UNA DISTRIBUCIÓN UNIFORME")
	}

	writer.Flush()
	fmt.Println("\nLa salida fue almacenada en el archivo:", filename)
}
