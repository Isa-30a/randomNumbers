package archivos

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var filenameAbs string

func Filename() string {
	return filenameAbs
}

func loadCSV(filename string) []string {
	file, err := os.Open(filename)

	if err != nil {
		panic(fmt.Errorf("error abriendo el archivo: %v", err))
	}

	reader := csv.NewReader(file)

	data := make([]string, 0, 10)

	for {
		fileData, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(fmt.Errorf("error leyendo el archivo: %v", err))
		}
		data = append(data, fileData[0])
	}

	return data
}

func LoadDataFromCSV(filename string) []float64 {
	finalFileName := filename
	filenameAbs = filename
	if !strings.HasSuffix(filename, ".csv") {
		finalFileName += ".csv"
	}

	dataAsString := loadCSV(finalFileName)

	data := make([]float64, len(dataAsString))

	for i, n := range dataAsString {
		n = strings.ReplaceAll(n, "\"", "")
		n = strings.ReplaceAll(n, ",", ".")
		value, err := strconv.ParseFloat(n, 64)
		if err != nil {
			panic(fmt.Errorf("error el linea %d del archivo: %s no es valor decimal.\n%v", i+1, n, err))
		}
		data[i] = value
	}

	return data
}
