package main

import (
	"fmt"
	"pruebas/archivos"
	"pruebas/poker"
)

func main() {
	filename := ""
	fmt.Print("Ingrese el nombre del archivo .csv donde se encuentran los números pseudo aleatorios: ")
	fmt.Scan(&filename)

	data := archivos.LoadDataFromCSV(filename)

	poker.PokerTest(data)
}
