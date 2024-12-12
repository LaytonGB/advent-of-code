package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	data string
	res1 int
	res2 int
)

// get data
func init() {
	file, err := os.Open("./3/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = strings.Join([]string{data, scanner.Text()}, "\n")
	}
}

// q1
func init() {
	r, err := regexp.Compile(`mul\((\d+),(\d+)\)`)
	if err != nil {
		log.Fatal(err)
	}

	matches := r.FindAllStringSubmatch(data, -1)
	for _, m := range matches {
		a, _ := strconv.Atoi(m[1])
		b, _ := strconv.Atoi(m[2])
		res1 += a * b
	}
}

// q2
func init() {
	r, err := regexp.Compile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	if err != nil {
		log.Fatal(err)
	}

	matches := r.FindAllStringSubmatch(data, -1)
	mul_enabled := true
	for _, m := range matches {
		switch m[0][0:3] {
		case "mul":
			if mul_enabled {
				a, _ := strconv.Atoi(m[1])
				b, _ := strconv.Atoi(m[2])
				res2 += a * b
			}
		case "do(":
			mul_enabled = true
		case "don":
			mul_enabled = false
		}
	}
}

func main() {
	log.Println(res1)
	log.Println(res2)
}
