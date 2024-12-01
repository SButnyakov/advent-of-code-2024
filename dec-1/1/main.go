package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

	var arr1, arr2 []int

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
		arr2 = append(arr2, num2)
	}

	sort.Ints(arr1)
	sort.Ints(arr2)

	res := 0

	for i := range arr1 {
		res += int(math.Abs(float64(arr1[i] - arr2[i])))
	}

	fmt.Println(res)
}
