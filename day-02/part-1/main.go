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

	file, err := os.Open("../input.txt")
	helpers.Check(err)
	defer file.Close()

	safeCounter := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		levels := strings.Split(line, " ")
		levelsInt := []int{}
		for _, val := range levels {
			levelsInt = append(levelsInt, helpers.ConvertStrToInt(val))
		}

		if safe := isSafe(levelsInt); safe {
			fmt.Printf("arr: %v , isSafe: %v \n", levelsInt, safe)
			safeCounter++
		}
	}

	logger.Info(fmt.Sprintf("Answer: %v", safeCounter))
}

func isSafe(arr []int) bool {
	return isIncreasinglySorted(arr) || isDecreasinglySorted(arr)
}

func isIncreasinglySorted(arr []int) bool {
	for i := 0; i < len(arr) - 1; i++ {
		if arr[i] > arr[i+1] || !checkDifference(arr[i], arr[i+1]) {
			return false
		}
	}
	return true
}

func isDecreasinglySorted(arr []int) bool {
	for i := 0; i < len(arr) - 1; i++ {
		if arr[i] < arr[i+1] || !checkDifference(arr[i], arr[i+1]) {
			return false
		}
	}
	return true
}

func checkDifference(x, y int) bool {
	difference := helpers.AbsInt(x - y) 
	return difference > 0 && difference < 4
}
