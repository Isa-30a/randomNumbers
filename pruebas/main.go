package main

import (
	"fmt"
	"pruebas/archivos"
	"pruebas/estadisticas"
	"pruebas/intervalos"
	"pruebas/kolgomorov"
	"pruebas/poker"
)

func main() {
	opc := 0
	var data []float64
	loaded := false
	filename := ""
	for opc != 6 {
		fmt.Println("\nBienvenido al módulo de pruebas de números pseudo aleatorios")
		fmt.Println("\n[1]. Cargar datos desde un archivo")
		fmt.Println("[2]. Ejecutar prueba de Poker")
		fmt.Println("[3]. Ejecutar prueba de intervalos")
		fmt.Println("[4]. Ejecutar prueba de Kolgomorov Smirnov")
		fmt.Println("[5]. Ejecutar prueba de la media")
		fmt.Println()
		fmt.Println("[6]. Salir")
		fmt.Println()
		fmt.Print("Seleccionar: ")
		fmt.Scan(&opc)
		switch opc {
		case 1:
			fmt.Print("Ingrese el nombre del archivo .csv donde se encuentran los números pseudo aleatorios: ")
			fmt.Scan(&filename)
			data = archivos.LoadDataFromCSV(filename)
			loaded = true
			fmt.Println("\nDatos cargados correctamente")
		case 2:
			if !loaded {
				fmt.Println("\nNo se han cargado datos")
				continue
			}
			poker.PokerTest(data)
		case 3:
			if !loaded {
				fmt.Println("\nNo se han cargado datos")
				continue
			}
			n := 0
			for n <= 1 {
				fmt.Print("Ingrese el número de intervalos (recomendado raiz de n): ")
				fmt.Scan(&n)
			}
			intervalos.IntervalosTest(data, n)
		case 4:
			if !loaded {
				fmt.Println("\nNo se han cargado datos")
				continue
			}
			kolgomorov.KolgomorovTest(data)
		case 5:
			if !loaded {
				fmt.Println("\nNo se han cargado datos")
				continue
			}
			estadisticas.MediaTest(data)
		}
	}
}
