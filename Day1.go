package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readDirections() ([]int, error) {
	file, err := os.Open("inputs/day1.txt")
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var directions []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 2 {
			continue
		}

		numStr := strings.TrimSpace(line[1:])
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, fmt.Errorf("error converting number: %v", err)
		}

		if line[0] == 'L' {
			num = -num
		}

		directions = append(directions, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return directions, nil
}

func findPassword() int {
	nums, err := readDirections()
	if err != nil {
		panic(err)
	}

	cur := 50
	password := 0

	for i := 0; i < len(nums); i++ {
		cur = (cur + nums[i]) % 100
		if cur == 0 {
			password++
		}
	}

	return password
}
