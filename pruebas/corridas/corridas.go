package corridas

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"pruebas/archivos"
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

	filename := archivos.Filename() + "_corridas_output.txt"
	file, _ := os.Create(filename)
	writer := bufio.NewWriter(file)

	fmt.Println("Ri\tX\t¿Corrida?")
	writer.WriteString("Ri\tX\t¿Corrida?\n")
	corrida := true
	Co := 0
	for i := range data {
		corrida = cmpArray[i] != iprev
		show := fmt.Sprintf("%.5f\t%d\t%v\n", data[i], cmpArray[i], corrida)
		fmt.Print(show)
		writer.WriteString(show)
		if corrida {
			Co++
		}
		iprev = cmpArray[i]
	}
	Cuo := (2*n - 1) / 3
	Vco := (16*n - 29) / 9
	Zo := math.Abs((float64(Co) - Cuo) / math.Sqrt(Vco))
	fmt.Println("\nNúmero de corridas:", Co)
	writer.WriteString(fmt.Sprintf("\nNúmero de corridas: %v", Co))

	fmt.Println()
	show := fmt.Sprintf("Zo = |(Co - Cuo) / Vco| = |(%d - %f)/%f| = %f", Co, Cuo, math.Sqrt(Vco), Zo)
	fmt.Print(show)
	writer.WriteString(show)
	if Zo < 1.96 {
		fmt.Println("\nEL CONJUNTO DE NÚMEROS ES INDEPENDIENTE")
		writer.WriteString("\nEL CONJUNTO DE NÚMEROS ES INDEPENDIENTE\n")
	} else {
		fmt.Println("\nEL CONJUNTO DE NÚMERO NO ES INDEPENDIENTE")
		writer.WriteString("\nEL CONJUNTO DE NÚMEROS NO ES INDEPENDIENTE\n")
	}

	writer.Flush()
	fmt.Println("\nLa salida fue almacenada en", filename)
}
