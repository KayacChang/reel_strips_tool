package main

import (
	"fmt"
	"log"
	"main/logics"
	"main/models"
	"main/slot"
)

func LtoR(substitution map[string][]string) logics.CheckRule {

	match := func(item, next string) bool {

		if check, ok := substitution[item]; ok {
			return Contain(next, check)
		}

		if _, ok := substitution[next]; ok {
			return true
		}

		return item == next
	}

	return func(line []string) (string, int) {
		fmt.Printf("%+v\n", line)

		item := line[0]
		count := 0

		for count+1 < len(line) {

			if !match(item, line[count+1]) {
				item = line[count]

				break
			}

			count += 1
		}

		return item, count
	}
}

func Contain(val string, slice []string) bool {

	for _, item := range slice {

		if item == val {
			return true
		}
	}
	return false
}

func main() {
	model, err := models.GetGameModelFrom("./input.json")
	if err != nil {
		log.Panicln(err)
	}

	display := slot.Spin(model.Reels.ThreeX5, model.Display)

	game := logics.LineGame{
		PayLines:  model.Paylines.ThreeX5,
		Paytable:  model.Paytable,
		CheckRule: LtoR(model.Substitution),
	}

	results := game.Check(display)

	for _, result := range results {
		fmt.Printf("%+v\n", result)
	}
}
