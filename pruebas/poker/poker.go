package poker

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type pokerHand struct {
	n     int
	group []string
	ei    float64
}

func PokerTest(data []float64) {
	values := make([][]string, len(data))
	possibles_results := map[string]pokerHand{
		"TD": {n: 0, group: make([]string, 0, 10), ei: 0.3024},
		"1P": {n: 0, group: make([]string, 0, 10), ei: 0.5040},
		"2P": {n: 0, group: make([]string, 0, 10), ei: 0.1080},
		"T ": {n: 0, group: make([]string, 0, 10), ei: 0.0090},
		"FH": {n: 0, group: make([]string, 0, 10), ei: 0.0720},
		"PK": {n: 0, group: make([]string, 0, 10), ei: 0.0045},
		"Q ": {n: 0, group: make([]string, 0, 10), ei: 0.0001},
	}
	for i, n := range data {
		n_string := strconv.FormatFloat(n, 'f', 5, 64)
		str, _ := strings.CutPrefix(n_string, "0.")
		str, _ = strings.CutPrefix(str, "1.")
		values[i] = strings.Split(str, "")

		result := verifyNumber(values[i])
		ph := possibles_results[result]
		ph.n += 1
		ph.group = append(ph.group, n_string)
		possibles_results[result] = ph
	}

	for hand, hand_data := range possibles_results {
		fmt.Println("Mano:", hand)
		fmt.Println("Número de elementos con esta mano:", hand_data.n)
		fmt.Println("Elementos en esta mano:", hand_data.group)
		fmt.Println("\n********************************************************")
	}

	fmt.Println("Cat\tOi\tEi\t\t\t(Ei - Oi)² / Ei")
	var sum float64
	for hand, hand_data := range possibles_results {
		ei := hand_data.ei * float64(len(data))
		r := math.Pow(ei-float64(hand_data.n), 2) / ei
		fmt.Printf("%s \t%d\t(%.5f)(%d)=%.5f   | \t%.5f\n", hand, hand_data.n, hand_data.ei, len(data), ei, r)
		sum += r
	}
	fmt.Println("X²o =", sum)
	if 12.56 < sum {
		fmt.Println("\nEL CONJUNTO DE NÚMEROS NO ES INDEPENDIENTE")
	} else {
		fmt.Println("\nEL CONJUNTO DE NÚMEROS ES INDEPENDIENTE")
	}
}

func verifyNumber(digits []string) string {
	reps := make(map[string]bool)
	for _, d := range digits {
		reps[d] = true
	}

	if len(reps) == 5 {
		return "TD"
	}

	if len(reps) == 4 {
		return "1P"
	}

	if len(reps) == 3 || len(reps) == 2 {
		return reverify(digits)
	}

	if len(reps) == 1 {
		return "Q "
	}

	return "FAILURE"
}

func reverify(digits []string) string {
	data := make(map[string]int)

	for i := range digits {
		n := data[digits[i]]
		data[digits[i]] = n + 1
	}

	n2 := 0
	for _, v := range data {
		if v == 4 {
			return "PK"
		}

		if v == 2 {
			n2++
		}
	}

	if n2 == 2 {
		return "2P"
	}

	if n2 == 1 {
		return "FH"
	}

	return "T "
}
