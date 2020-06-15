package models

import (
	"encoding/json"
	"io/ioutil"
)

type GameModel struct {
	Display  []int `json:"display"`
	Paytable struct {
		Wild []int `json:"Wild"`
		H1   []int `json:"H1"`
		H2   []int `json:"H2"`
		H3   []int `json:"H3"`
		H4   []int `json:"H4"`
		H5   []int `json:"H5"`
		L1   []int `json:"L1"`
		L2   []int `json:"L2"`
		L3   []int `json:"L3"`
		L4   []int `json:"L4"`
	} `json:"paytable"`
	Substitution struct {
		Wild []string `json:"Wild"`
	} `json:"substitution"`
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
