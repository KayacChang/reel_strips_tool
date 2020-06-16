package logics

type CheckRule func(line []string) (string, int)

type LineGame struct {
	PayLines  [][]int
	Paytable  map[string][]int
	CheckRule CheckRule
}

type Result struct {
	LineID     int
	Line       []int
	Symbols    []string
	Item       string
	Count      int
	Multiplier int
}

func mapping(payline []int, display [][]string) []string {

	line := make([]string, 0, len(payline))

	for col, row := range payline {

		symbol := display[col][row]

		line = append(line, symbol)
	}

	return line
}

func (it LineGame) Check(display [][]string) []Result {

	results := []Result{}

	for lineID, payline := range it.PayLines {

		line := mapping(payline, display)

		item, hits := it.CheckRule(line)

		mapBy, ok := it.Paytable[item]
		if !ok {
			continue
		}

		mul := mapBy[hits]
		if mul == 0 {
			continue
		}

		results = append(results, Result{
			LineID:     lineID,
			Line:       payline,
			Symbols:    line,
			Item:       item,
			Count:      hits + 1,
			Multiplier: mul,
		})
	}

	return results
}
