package models

import (
	"encoding/json"
	"io/ioutil"
)

type GameModel struct {
	Display []int `json:"display"`

	Paytable map[string][]int `json:"paytable"`

	Substitution map[string][]string `json:"substitution"`

	Reels struct {
		ThreeX5 [][]string `json:"3x5"`
	} `json:"reels"`

	Paylines struct {
		ThreeX5 [][]int `json:"3x5"`
	} `json:"paylines"`
}

func GetGameModelFrom(filePath string) (GameModel, error) {
	result := GameModel{}

	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(bytes, &result); err != nil {
		return result, err
	}

	return result, nil
}
