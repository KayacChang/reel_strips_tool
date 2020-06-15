package main

import (
	"fmt"
	"log"
	"main/models"
	"main/slot"
)

func main() {
	model, err := models.GetGameModelFrom("./input.json")
	if err != nil {
		log.Panicln(err)
	}

	for i := 0; i < 1e3; i += 1 {

		res := slot.Spin(model.Reels.ThreeX5, model.Display)

		fmt.Printf("%+v\n", res)
	}
}
