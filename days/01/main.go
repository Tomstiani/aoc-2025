package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: go run main.go <input>")
	}

	inputPath := os.Args[1]

	input, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatalf("failed to read input: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	part1Answer := part1(lines)
	part2Answer := part2(lines)

	fmt.Println("Part 1:", part1Answer)
	fmt.Println("Part 2:", part2Answer)
}

func part1(lines []string) int {
	currentPos := 50
	zeroCount := 0

	for _, line := range lines {
		direction := line[0]
		distance, _ := strconv.Atoi(line[1:])

		if direction == 'L' {
			currentPos -= distance
		} else {
			currentPos += distance
		}

		currentPos = ((currentPos % 100) + 100) % 100

		if currentPos == 0 {
			zeroCount++
		}
	}

	return zeroCount
}

func part2(lines []string) int {
	currentPos := 50
	zeroCount := 0

	// Not very efficient, but it works
	// Could be optimized with math probably
	// If ever in doubt, brute force it
	for _, line := range lines {
		direction := line[0]
		distance, _ := strconv.Atoi(line[1:])

		for click := 0; click < distance; click++ {
			if direction == 'L' {
				currentPos--
				if currentPos < 0 {
					currentPos = 99
				}
			} else {
				currentPos++
				if currentPos > 99 {
					currentPos = 0
				}
			}

			if currentPos == 0 {
				zeroCount++
			}
		}
	}

	return zeroCount
}
