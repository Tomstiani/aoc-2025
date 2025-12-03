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

	part1Answer := part1(string(input))
	part2Answer := part2(string(input))

	fmt.Println("Part 1:", part1Answer)
	fmt.Println("Part 2:", part2Answer)
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0

	for _, line := range lines {
		maxSoFar := 0
		bestAnswer := 0
		for i := len(line) - 1; i >= 0; i-- {
			digit, _ := strconv.Atoi(string(line[i]))
			if i == len(line)-1 {
				maxSoFar = digit
				continue
			}

			if digit*10+maxSoFar > bestAnswer {
				bestAnswer = digit*10 + maxSoFar
			}

			if digit > maxSoFar {
				maxSoFar = digit
			}
		}

		sum += bestAnswer
	}

	return sum
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		result := ""
		digitsLeft := 12 - len(result)
		startPos := 0

		for i := 0; i < digitsLeft; i++ {
			workableLine := line[startPos : len(line)-digitsLeft+i+1]
			biggestDigit := 0
			posOfBiggestDigit := 0

			for j, digit := range workableLine {
				digitInt, _ := strconv.Atoi(string(digit))
				if digitInt > biggestDigit {
					biggestDigit = digitInt
					posOfBiggestDigit = j
				}
			}

			result += strconv.Itoa(biggestDigit)
			startPos = posOfBiggestDigit + startPos + 1
		}

		resultInt, _ := strconv.Atoi(result)
		sum += resultInt
	}
	return sum
}
