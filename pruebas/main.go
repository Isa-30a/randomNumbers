package main

import (
	"fmt"
	"pruebas/archivos"
	"pruebas/bernoulli"
	"pruebas/corridas"
	"pruebas/estadisticas"
	"pruebas/intervalos"
	"pruebas/kolgomorov"
	"pruebas/poisson"
	"pruebas/poker"
	"strconv"
	"strings"
)

var functions = map[string]func(data []float64, inputs []string){
	"BERNOULLI":  Bernoulli,
	"CORRIDAS":   Corridas,
	"MEDIA":      Media,
	"VARIANZA":   Varianza,
	"INTERVALOS": Intervalos,
	"KOLGOMOROV": Kolgomorov,
	"POISSON":    Poisson,
	"POKER":      Poker,
}

func Bernoulli(data []float64, inputs []string) {
	pe, err := strconv.ParseFloat(inputs[2], 64)
	if err != nil {
		panic(err)
	}
	bernoulli.BernoulliVariable(data, pe)
}

func Corridas(data []float64, inputs []string) {
	corridas.CorridasTest(data)
}

func Media(data []float64, inputs []string) {
	estadisticas.MediaTest(data)
}

func Varianza(data []float64, inputs []string) {
	estadisticas.VarianzaTest(data)
}

func Intervalos(data []float64, inputs []string) {
	n, err := strconv.Atoi(inputs[2])
	if err != nil {
		panic(err)
	}
	intervalos.IntervalosTest(data, n)
}

func Kolgomorov(data []float64, inputs []string) {
	kolgomorov.KolgomorovTest(data)
}

func Poisson(data []float64, inputs []string) {
	x, err := strconv.ParseFloat(inputs[2], 64)
	if err != nil {
		panic(err)
	}
	n, err := strconv.Atoi(inputs[3])
	if err != nil {
		panic(err)
	}
	poisson.PoissonVariable(data, x, n)
}

func Poker(data []float64, inputs []string) {
	poker.PokerTest(data)
}

func main() {
	var stdin string
	fmt.Scan(&stdin)
	inputs := strings.Split(stdin, ";")
	if len(inputs) < 2 {
		panic("la entrada no fue valida")
	}
	inputs[0] = strings.ToUpper(inputs[0])
	data := archivos.LoadDataFromCSV(inputs[1])
	function, ok := functions[inputs[0]]
	if !ok {
		panic("unknowed function")
	}
	function(data, inputs)
}
