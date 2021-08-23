package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var current, next, path = 0, 1, 0
	var max = inputMax()

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
