package logics

import "fmt"

func hit(line []string) {

	for count := 0; count < len(line); count += 1 {

		// window := line[count]

	}

}

func mapToLine(payline []int, display [][]string) []string {

	res := make([]string, 0, len(payline))

	for col, row := range payline {

		symbol := display[col][row]

		res = append(res, symbol)
	}

	return res
}

func Check(paylines [][]int, display [][]string) {

	for _, payline := range paylines {

		line := mapToLine(payline, display)

		fmt.Printf("%+v\n", line)
	}
}
