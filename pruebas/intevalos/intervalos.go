package intervalos

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/stat/distuv"
)

type intervalo struct {
	n     int
	group []float64
}

func IntervalosTest(data []float64, n int) {
	intervalo_value := 1 / float64(n)
	intervalos := make(map[float64]intervalo)
	for i := intervalo_value; i <= 1; i += intervalo_value {
		intervalos[i] = intervalo{n: 0, group: make([]float64, 0, 10)}
	}

	if len(intervalos) == n-1 {
		intervalos[1] = intervalo{n: 0, group: make([]float64, 0, 10)}
	}

	if len(intervalos) != n {
		fmt.Println("error: el número de intervalos no parece cuadrar")
		return
	}

	for _, d := range data {
		var possible_range float64
		possible_range = 1.1
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
	fmt.Println("Subgrupo\tN° Elementos")
	for k, v := range intervalos {
		fmt.Printf("%f\t%d\n", k, v.n)
	}

	fmt.Println("\nFormula:")
	fmt.Printf("X²o = (1/%d)*{", N)
	var sum float64
	for _, v := range intervalos {
		fmt.Printf("(%d - %d)² + ", v.n, N)
		sum += math.Pow(float64(v.n-N), 2)
	}
	X2o := 1 / float64(N) * sum
	fmt.Printf(" } = %f\n", X2o)
	X2 := distuv.ChiSquared{K: float64(n - 1)}.Quantile(0.05)
	fmt.Printf("X²[0.05,%d] = %f\n", n-1, X2)

	if X2o < X2 {
		fmt.Println("\nEL CONJUNTO DE NÚMEROS ESTÁ DISTRIBUIDO UNIFORMEMENTE")
	} else {
		fmt.Println("\nEL CONJUNTO DE NÚMEROS NO ESTÁ DISTRIBUIDO UNIFORMEMENTE")
	}
}
