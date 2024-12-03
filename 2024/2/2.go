package main

import (
	"bufio"
	"log"
	"os"
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
					log.Printf("BAD %v\n", data[i])
					continue Outer1
				}
			case INCREASING:
				if last <= curr || last > curr+3 {
					log.Printf("BAD %v\n", data[i])
					continue Outer1
				}
			}
			last = curr
		}

		log.Printf("GOOD %v\n", data[i])
		res1 += 1
	}

	// q2
Outer2:
	for i := range 1000 {
		var (
			bar      int
			problems int
		)

		last := data[i][0]
		for _, curr := range data[i][1:] {
			var new_bar = bar
			if last < curr && last >= curr-3 {
				new_bar += 1
			} else if last > curr && last <= curr+3 {
				new_bar -= 1
			}

			var is_new_bar_bad bool
			if bar < 0 {
				is_new_bar_bad = bar <= new_bar
			} else if bar > 0 {
				is_new_bar_bad = bar >= new_bar
			} else {
				is_new_bar_bad = bar == new_bar
			}

			if is_new_bar_bad {
				problems += 1
			}

			if problems >= 2 {
				log.Printf("BAD %v\n", data[i])
				continue Outer2
			}

			bar = new_bar
		}

		log.Printf("GOOD %v\n", data[i])
		res2 += 1
	}
}

func main() {
	log.Println(res1)
	log.Println(res2)
}
