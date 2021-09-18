package main

import (
	"fmt"
	"lts-demo/leetcode_test/unit"
	"strconv"
)

func main() {
	print(unit.RandomString())
}

func increment(numbers []string) bool {
	isOverflow := false
	length := len(numbers)
	for i := length - 1; i >= 0; i-- {
		sum, _ := strconv.Atoi(numbers[i])
		if sum == 9 {
			if i == 0 {
				return true
			}
			numbers[i] = strconv.Itoa(0)
			continue
		}
		sum += 1
		numbers[i] = strconv.Itoa(sum)
		break
	}

	return isOverflow
}

func printToMaxOfDigitsRecursively(numbers []string, length int, index int) {
	if index == length-1 {
		fmt.Println(numbers)
		return
	}
	for i := 0; i < 10; i++ {
		numbers[index+1] = strconv.Itoa(i)
		printToMaxOfDigitsRecursively(numbers, length, index+1)
	}
}
