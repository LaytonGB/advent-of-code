package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	numbers [2][1000]int
	counts  map[int]int

	res1 int
	res2 int
)

func init() {
	// q1
	file, err := os.Open("./1/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		txt := scanner.Text()
		parts := strings.Fields(txt)

		numbers[0][i], _ = strconv.Atoi(parts[0])
		numbers[1][i], _ = strconv.Atoi(parts[1])

		i += 1
	}

	sort.Slice(numbers[0][:], func(i, j int) bool { return numbers[0][i] < numbers[0][j] })
	sort.Slice(numbers[1][:], func(i, j int) bool { return numbers[1][i] < numbers[1][j] })

	var diffs [1000]int
	for i := range 1000 {
		diffs[i] = max(numbers[0][i], numbers[1][i]) - min(numbers[0][i], numbers[1][i])
	}

	for _, d := range diffs {
		res1 += d
	}

	// q2
	counts = make(map[int]int)

	for i := range 1000 {
		counts[numbers[1][i]] += 1
	}

	for i := range 1000 {
		res2 += numbers[0][i] * counts[numbers[0][i]]
	}
}

func main() {
	log.Println(res1)
	log.Println(res2)
}
