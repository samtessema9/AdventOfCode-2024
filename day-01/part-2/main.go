package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/samtessema9/AdventOfCode-2024/helpers"
)

func main() {
	// Configure a logger
	textHandler := slog.NewTextHandler(os.Stdout, nil)
	logger := slog.New(textHandler)

	var leftSlice []int
	var rightSlice []int

	file, err := os.Open("../input.txt")
	helpers.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		// Each line will contain 2 ints separated by 3 spaces
		//
		// ex. "123   456"
		s := strings.Split(line, "   ")

		// Convert from str to int and append to respective list
		// Program will exit if conversion fails
		leftSlice = append(leftSlice, helpers.ConvertStrToInt(s[0]))
		rightSlice = append(rightSlice, helpers.ConvertStrToInt(s[1]))
	}

	count := make(map[int]int)
	for _, val := range rightSlice {
		_, ok := count[val]
		if !ok {
			count[val] = 1
		} else {
			count[val] += 1
		}
	}

	similarityScore := 0
	for _, val := range leftSlice {
		multiplier := 0
		count, ok := count[val]
		if ok {
			multiplier = count
		}
		similarityScore += val * multiplier
	}

	logger.Info(fmt.Sprintf("Answer: %v", similarityScore))
}
