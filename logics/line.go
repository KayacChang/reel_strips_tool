package logics

type CheckRule func(line []string, skip SkipRule) (string, int)
type SkipRule func(next string, item string) bool

type LineGame struct {
	PayLines [][]int
	Paytable map[string][]int
	Check    CheckRule
	Skip     SkipRule
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

func (it LineGame) hit(line []string, skip SkipRule) (string, int, int) {

	item, hits := it.Check(line, skip)

	mapBy, ok := it.Paytable[item]
	if !ok {
		return item, hits, 0
	}

	return item, hits, mapBy[hits]
}

func noSkip(item string, next string) bool {
	return false
}

func (it LineGame) compare(line []string) (string, int, int) {

	item1, hits1, mul1 := it.hit(line, noSkip)
	item2, hits2, mul2 := it.hit(line, it.Skip)

	if mul1 > mul2 {

		return item1, hits1, mul1
	}

	return item2, hits2, mul2
}

func (it LineGame) Exec(display [][]string) []Result {

	results := []Result{}

	for lineID, payline := range it.PayLines {

		line := mapping(payline, display)

		item, hits, mul := it.compare(line)

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
