package bernoulli

import (
	"bufio"
	"fmt"
	"os"
	"pruebas/archivos"
)

func BernoulliVariable(data []float64, pe float64) {
	pf := 1 - pe
	exitoFunc := func(data float64) int {
		if data <= pe {
			return 1
		}
		return 0
	}
	sumExitosos := 0
	filename := archivos.Filename() + "_bernoulli_output.txt"
	file, _ := os.Create(filename)
	writer := bufio.NewWriter(file)

	fmt.Println("Probabilidad de éxito:", pe)
	writer.WriteString(fmt.Sprintf("Probabilidad de éxito: %.2f\n", pe))

	fmt.Println("Probabilidad de falla:", pf)
	writer.WriteString(fmt.Sprintf("Probabilidad de fracaso: %.2f\n", pf))

	fmt.Println("i\tRi    \tXi\t¿éxtoso?")
	writer.WriteString("i\tRi    \tXi\t¿éxtoso?\n")
	for i, d := range data {
		exito := exitoFunc(d)
		show := fmt.Sprintf("%d\t%.5f\t%d\t%v\n", i+1, d, exito, exito == 1)
		writer.WriteString(show)
		fmt.Print(show)
		if exito == 1 {
			sumExitosos++
		}
	}
	pde := (float64(sumExitosos) / float64(len(data))) * 100
	fmt.Printf("Porcentaje de éxitos: %.2f%s \n", pde, "%")
	writer.WriteString(fmt.Sprintf("Porcentaje de éxitos: %.2f%s \n", pde, "%"))

	fmt.Printf("Porcentaje de fracasos: %.2f%s \n", 100-pde, "%")
	writer.WriteString(fmt.Sprintf("Porcentaje de fracasos: %.2f%s \n", 100-pde, "%"))

	writer.Flush()
	fmt.Println("\nLa salida fue almacenada en", filename)
}
