package corridas

import (
	"fmt"
	"math"
)

func CorridasTest(data []float64) {
	n := float64(len(data))
	prev := -1.0
	cmpArray := make([]int, len(data))
	for i := range data {
		if data[i] > prev {
			cmpArray[i] = 1
		} else {
			cmpArray[i] = 0
		}
		prev = data[i]
	}
	iprev := -1
	fmt.Println("Ri\tX\t¿Corrida?")
	corrida := true
	Co := 0
	for i := range data {
		corrida = cmpArray[i] != iprev
		fmt.Printf("%.5f\t%d\t%v\n", data[i], cmpArray[i], corrida)
		if corrida {
			Co++
		}
		iprev = cmpArray[i]
	}
	Cuo := (2*n - 1) / 3
	Vco := (16*n - 29) / 9
	Zo := math.Abs((float64(Co) - Cuo) / math.Sqrt(Vco))
	fmt.Println("\nNúmero de corridas:", Co)
	fmt.Println("")
	fmt.Printf("Zo = |(Co - Cuo) / Vco| = |(%d - %f)/%f| = %f", Co, Cuo, math.Sqrt(Vco), Zo)
	if Zo < 1.96 {
		fmt.Println("\nEL CONJUNTO DE NÚMEROS ES INDEPENDIENTE")
	} else {
		fmt.Println("\nEL CONJUNTO DE NÚMERO NO ES INDEPENDIENTE")
	}
}
