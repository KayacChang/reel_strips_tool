package slot

import (
	"math/rand"
	"time"

	"github.com/seehuhn/mt19937"
)

var generator *rand.Rand

func init() {
	generator = rand.New(mt19937.New())

	generator.Seed(time.Now().UnixNano())
}

func random() int {
	return int(generator.Int63())
}

func slice(reel []string, from, length int) []string {

	res := make([]string, 0, length)

	for offset := 0; offset < length; offset += 1 {

		pos := (from + offset) % len(reel)

		res = append(res, reel[pos])
	}

	return res
}

func Spin(reels [][]string, display []int) [][]string {

	res := make([][]string, 0, len(reels))

	for idx, reel := range reels {

		pos := random() % len(reel)

		height := display[idx]

		symbols := slice(reel, pos, height)

		res = append(res, symbols)
	}

	return res
}
