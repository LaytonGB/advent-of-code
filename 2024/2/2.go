package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Change int

const (
	INCREASING Change = iota
	DECREASING
)

var (
	data [1000][]int
	res1 int
	res2 int
)

func init() {
	// q1
	file, err := os.Open("./2/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		txt := scanner.Text()
		parts := strings.Fields(txt)

		for _, p := range parts {
			n, _ := strconv.Atoi(p)
			data[i] = append(data[i], n)
		}

		i += 1
	}

Outer1:
	for i := range 1000 {
		var change Change

		switch {
		case data[i][0] < data[i][1]:
			if data[i][0] < data[i][1]-3 {
				continue
			}
			change = DECREASING
		case data[i][0] > data[i][1]:
			if data[i][0] > data[i][1]+3 {
				continue
			}
			change = INCREASING
		default:
			continue
		}

		last := data[i][1]
		for _, curr := range data[i][2:] {
			switch change {
			case DECREASING:
				if last >= curr || last < curr-3 {
					continue Outer1
				}
			case INCREASING:
				if last <= curr || last > curr+3 {
					continue Outer1
				}
			}
			last = curr
		}

		res1 += 1
	}
}

// q2
func abs(a int) int {
	return max(a, a*-1)
}

func isSameSign(a, b int) bool {
	return a/abs(a) == b/abs(b)
}

func init() {
	var diffs [1000][]int

	for i, level := range data {
		diffs[i] = make([]int, len(level)-1)
		for j := range len(level) - 1 {
			diffs[i][j] = level[j+1] - level[j]
		}
	}

	for _, level := range diffs {
		tolerance := 2

	Outer2:
		for tolerance > 0 {
			if level[0] == 0 || abs(level[0]) > 3 {
				level = level[1:]
				tolerance -= 1
				continue
			}

			j := 0
			for j < len(level)-1 {
				if level[j+1] == 0 || abs(level[j+1]) > 3 || !isSameSign(level[j], level[j+1]) {
					level = slices.Concat(level[:j+1], level[j+2:])
					tolerance -= 1
					continue Outer2
				}
				j += 1
			}

			break
		}

		if tolerance > 0 {
			res2 += 1
		}
	}
}

func main() {
	log.Println(res1)
	log.Println(res2)
}
