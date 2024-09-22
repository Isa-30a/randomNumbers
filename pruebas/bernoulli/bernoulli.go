package bernoulli

import "fmt"

func BernoulliVariable(data []float64, pe float64) {
	pf := 1 - pe
	exitoFunc := func(data float64) int {
		if data <= pe {
			return 1
		}
		return 0
	}
	sumExitosos := 0
	fmt.Println("Probabilidad de éxito:", pe)
	fmt.Println("Probabilidad de falla:", pf)
	fmt.Println("i\tRi    \tXi\t¿éxtoso?")
	for i, d := range data {
		exito := exitoFunc(d)
		fmt.Printf("%d\t%.5f\t%d\t%v\n", i+1, d, exito, exito == 1)
		if exito == 1 {
			sumExitosos++
		}
	}
	pde := (float64(sumExitosos) / float64(len(data))) * 100
	fmt.Printf("Porcentaje de éxitos: %.2f%s \n", pde, "%")
	fmt.Printf("Porcentaje de fracasos: %.2f%s \n", 100-pde, "%")
}
