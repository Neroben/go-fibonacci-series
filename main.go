package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var current, next, path = 1, 1, 1
	var max = inputMax()

	fmt.Print(strconv.Itoa(current) + " ")
	for path < max {
		fmt.Print(strconv.Itoa(path) + " ")
		path = current + next
		current = next
		next = path
	}
}

func inputMax() int {
	var max = 0
	fmt.Print("Введите предельное значение: ")
	_, err := fmt.Scanf("%d", &max)
	if err != nil {
		os.Exit(1)
	}
	return max
}
