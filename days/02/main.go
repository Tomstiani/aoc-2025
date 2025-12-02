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

	part1Answer := part1(strings.TrimSpace(string(input)))
	part2Answer := part2(strings.TrimSpace(string(input)))

	fmt.Println("Part 1:", part1Answer)
	fmt.Println("Part 2:", part2Answer)
}

func part1(input string) int {
	var repeats []int

	for part := range strings.SplitSeq(input, ",") {
		start, end := strings.Split(part, "-")[0], strings.Split(part, "-")[1]

		startInt, _ := strconv.Atoi(start)
		endInt, _ := strconv.Atoi(end)

		for i := startInt; i <= endInt; i++ {
			s := strconv.Itoa(i)

			// If the string length is odd, can't be repeat
			if len(s)%2 != 0 {
				continue
			}

			firstHalf := s[:len(s)/2]
			secondHalf := s[len(s)/2:]

			if firstHalf == secondHalf {
				repeats = append(repeats, i)
			}
		}
	}

	// sum of all repeats
	sum := 0
	for _, r := range repeats {
		sum += r
	}

	return sum
}

func part2(input string) int {
	sum := 0

	for rangeStr := range strings.SplitSeq(input, ",") {
		bounds := strings.Split(rangeStr, "-")
		from, _ := strconv.Atoi(bounds[0])
		to, _ := strconv.Atoi(bounds[1])

		for num := from; num <= to; num++ {
			numStr := strconv.Itoa(num)

			// Try each possible pattern length (1 up to half the string)
			for patternLen := 1; patternLen <= len(numStr)/2; patternLen++ {
				// Pattern must divide evenly into string length
				if len(numStr)%patternLen != 0 {
					continue
				}

				pattern := numStr[:patternLen]

				repeatCount := len(numStr) / patternLen
				if strings.Repeat(pattern, repeatCount) == numStr {
					sum += num
					break
				}
			}
		}
	}
	return sum
}
