package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"slices"
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

	// Sort slices 
	slices.Sort(leftSlice)
	slices.Sort(rightSlice)

	totalDistance := 0
	for i := 0; i < len(leftSlice); i++ {
		totalDistance += helpers.AbsInt(leftSlice[i] - rightSlice[i])
	}

	logger.Info(fmt.Sprintf("Answer: %v", totalDistance))
}