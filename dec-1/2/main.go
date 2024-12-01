package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	arr1 := make([]int, 0)
	map2 := make(map[int]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		arr1 = append(arr1, num1)
		map2[num2] += 1
	}

	res := 0

	for _, v := range arr1 {
		res += v * map2[v]
	}

	fmt.Println(res)
}
