package main

import (
	"fmt"
	"log"
	"main/logics"
	"main/models"
	"main/slot"
)

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

	skip := func(item string, next string) bool {
		mapping, ok := model.Substitution[next]

		return ok && Contain(item, mapping)
	}

	game := logics.LineGame{
		PayLines: model.Paylines.ThreeX5,
		Paytable: model.Paytable,
		Check:    logics.LtoR,
		Skip:     skip,
	}

	results := game.Exec(display)

	for _, result := range results {
		fmt.Printf("%+v\n", result)
	}
}
