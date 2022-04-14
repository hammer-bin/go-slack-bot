package app

import (
	"bufio"
	"os"
)

const fileName = "candidates.txt"

func readCandidates() ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var candidates []string
	for scanner.Scan() {
		candidates = append(candidates, scanner.Text())
	}
	return candidates, nil
}