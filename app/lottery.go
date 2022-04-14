package app

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const candidateCnt = 2

func Lottery() {

	candidates, err := readCandidates()
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot read candidates file!", err)
		return
	}

	rand.Seed(time.Now().UnixNano())

	winners := make([]string, candidateCnt)
	for i := 0; i < candidateCnt; i++ {
		n := rand.Intn(len(candidates))
		winners[i] = candidates[n]
		candidates = append(candidates[:n], candidates[n+1:]...)
	}

	fmt.Println("Winners are !!!")
	for _, winner := range winners {
		fmt.Println(winner)
	}
}
