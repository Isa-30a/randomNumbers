package kolgomorov

import (
	"fmt"
	"math"
	"slices"
)

func KolgomorovTest(data []float64) {
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
	n := len(clonedData)
	divsArrays := make([]float64, n)
	for i := range n {
		divsArrays[i] = float64(i+1) / float64(n)
	}
	max := -1.0
	for i := range n {
		fmt.Printf("[%.1f - %.5f], ", divsArrays[i], clonedData[i])
		divsArrays[i] = math.Abs(divsArrays[i] - clonedData[i])
		if divsArrays[i] > max {
			max = divsArrays[i]
		}
	}
	fmt.Println()
	for i := range n {
		fmt.Printf("[%.2f], ", divsArrays[i])
	}
	fmt.Println()
	fmt.Println("Máximo valor:", max)
	ks := math.Sqrt(-math.Log(0.05/2) / (2 * float64(n)))
	fmt.Println("KS:", ks)
	fmt.Printf("Dmax < Dks = %v", max < ks)
	if max < ks {
		fmt.Println("\nLOS NÚMEROS TIENEN UNA DISTRIBUCIÓN UNIFORME")
	} else {
		fmt.Println("\nLOS NÚMEROS NO TIENEN UNA DISTRIBUCIÓN UNIFORME")
	}
}
