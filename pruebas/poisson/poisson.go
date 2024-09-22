package poisson

import (
	"fmt"

	"gonum.org/v1/gonum/stat/distuv"
)

type Limit struct {
	li float64
	ls float64
	i  int
	n  int
}

func searchBetweenLimits(value float64, limits []Limit) int {
	for _, v := range limits {
		if v.li <= value && value <= v.ls {
			return v.i
		}
	}
	return -1
}

func PoissonVariable(data []float64, x float64, n int) {
	poisson_data := make([]float64, 0, 10)
	poisson := distuv.Poisson{Lambda: x}
	limits := make([]Limit, 0, n)
	fmt.Println("i\tP(x)\tLi\tLs")
	last := 0.0
	for i := range n {
		poisson_data = append(poisson_data, poisson.Prob(float64(i)))
		fmt.Printf("%d\t%.5f\t%.5f\t%.5f\n", i, poisson_data[i], last, last+poisson_data[i])
		limits = append(limits, Limit{i: i, li: last, ls: last + poisson_data[i], n: 0})
		last += poisson_data[i]
	}

	limits[len(limits)-1].ls = 1.0

	fmt.Println("\nRandom\tGrupo pertenencia")
	for _, d := range data {
		pos := searchBetweenLimits(d, limits)
		if pos == -1 {
			fmt.Println("Se determinó que", d, "no se encuentra en ningun limite")
		}
		fmt.Printf("%.5f\t%d\n", d, limits[pos].i)
		limit := limits[pos]
		limit.n += 1
		limits[pos] = limit
	}
	fmt.Println("i\tP(x)\tLi\tLs\tCount\tPorcentaje")
	total := float64(len(data))
	for i, limit := range limits {
		fmt.Printf("%d\t%.5f\t%.5f\t%.5f\t%d\t%.2f%s\n", i, poisson_data[limit.i], limit.li, limit.ls, limit.n, (float64(limit.n)/total)*100, "%")
	}
}