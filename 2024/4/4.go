package main

import (
	"bufio"
	"log"
	"os"
)

var (
	data []string
	res1 int
	res2 int
)

// get data
func init() {
	file, err := os.Open("./4/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		data = append(data, scanner.Text())
		i += 1
	}
}

// q1
func addIfWordFound(target string, char, row, col, dy, dx int) {
	if char >= len(target) {
		res1 += 1
		return
	}
	if row+dy < 0 || row+dy >= len(data) || col+dx < 0 || col+dx >= len(data[row+dy]) {
		return
	}
	if data[row+dy][col+dx] == target[char] {
		addIfWordFound(target, char+1, row+dy, col+dx, dy, dx)
	}
}

func init() {
	target := "XMAS"
	for i := range len(data) {
		for j := range len(data[i]) {
			if data[i][j] == target[0] {
				for k := range 3 {
					for l := range 3 {
						dy, dx := k-1, l-1
						if dy == 0 && dx == 0 || i+dy < 0 || i+dy >= len(data) || j+dx < 0 || j+dx >= len(data[i+dy]) {
							continue
						}
						if data[i+dy][j+dx] == target[1] {
							addIfWordFound(target, 2, i+dy, j+dx, dy, dx)
						}
					}
				}
			}
		}
	}
}

// q2
func doDiagonalsMakeMas(row, col int) bool {
	if row == 0 || row == len(data)-1 || col == 0 || col == len(data[row])-1 {
		return false
	}
	if data[row-1][col-1] == 'M' && data[row+1][col-1] == 'S' && data[row+1][col+1] == 'S' && data[row-1][col+1] == 'M' ||
		data[row-1][col-1] == 'M' && data[row+1][col-1] == 'M' && data[row+1][col+1] == 'S' && data[row-1][col+1] == 'S' ||
		data[row-1][col-1] == 'S' && data[row+1][col-1] == 'M' && data[row+1][col+1] == 'M' && data[row-1][col+1] == 'S' ||
		data[row-1][col-1] == 'S' && data[row+1][col-1] == 'S' && data[row+1][col+1] == 'M' && data[row-1][col+1] == 'M' {
		return true
	}
	return false
}

func init() {
	for i := range len(data) {
		for j := range len(data[i]) {
			if data[i][j] == 'A' && doDiagonalsMakeMas(i, j) {
				res2 += 1
			}
		}
	}
}

func main() {
	log.Println(res1)
	log.Println(res2)
}
