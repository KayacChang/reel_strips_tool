package logics

func LtoR(line []string, skip SkipRule) (string, int) {

	count := 0
	item := line[count]

	for count+1 < len(line) {

		next := line[count+1]

		if skip(item, next) {
			count++

			continue
		}

		if next == item {
			count++

			item = next

			continue
		}

		break
	}

	return item, count
}
