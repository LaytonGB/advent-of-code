package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"os/signal"
	"slices"
	"syscall"
)

type Coord struct {
	row int
	col int
}

func (pos Coord) IsInBounds(rowBound, colBound int) bool {
	return pos.row >= 0 && pos.row < rowBound && pos.col >= 0 && pos.col < colBound
}

func (pos Coord) AddCoord(pos2 Coord) Coord {
	return Coord{row: pos.row + pos2.row, col: pos.col + pos2.col}
}

func (pos Coord) StepInDirection(d D) Coord {
	switch d {
	case UP:
		pos.row -= 1
	case RIGHT:
		pos.col += 1
	case DOWN:
		pos.row += 1
	case LEFT:
		pos.col -= 1
	}
	return pos
}

type D int

const (
	UP D = iota
	DOWN
	LEFT
	RIGHT
)

func (d D) Next() D {
	switch d {
	case UP:
		return RIGHT
	case RIGHT:
		return DOWN
	case DOWN:
		return LEFT
	case LEFT:
		return UP
	default:
		return d
	}
}

func (d D) AsRune() rune {
	switch d {
	case UP:
		return '^'
	case RIGHT:
		return '>'
	case DOWN:
		return 'v'
	case LEFT:
		return '<'
	default:
		return '_'
	}
}

func (d D) AsCoord() Coord {
	switch d {
	case UP:
		return Coord{row: -1, col: 0}
	case RIGHT:
		return Coord{row: 0, col: 1}
	case DOWN:
		return Coord{row: 1, col: 0}
	case LEFT:
		return Coord{row: 0, col: -1}
	default:
		return Coord{row: 0, col: 0}
	}
}

func getNextRune(data [][]rune, pos Coord, d D) *rune {
	switch d {
	case UP:
		if pos.row == 0 {
			return nil
		}
		return &data[pos.row-1][pos.col]
	case RIGHT:
		if pos.col == len(data[pos.row])-1 {
			return nil
		}
		return &data[pos.row][pos.col+1]
	case DOWN:
		if pos.row == len(data)-1 {
			return nil
		}
		return &data[pos.row+1][pos.col]
	case LEFT:
		if pos.col == 0 {
			return nil
		}
		return &data[pos.row][pos.col-1]
	default:
		return nil
	}
}

func directionFromRune(r rune) (d D, err error) {
	switch r {
	case '^':
		return UP, nil
	case '>':
		return RIGHT, nil
	case 'v':
		return DOWN, nil
	case '<':
		return LEFT, nil
	default:
		return UP, errors.New("invalid direction rune")
	}
}

var (
	err  error
	data [][]rune
	res1 int
	res2 int
)

func init() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
}

// get data
func init() {
	file, err := os.Open("./6/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		t := scanner.Text()
		data = append(data, []rune(t))
		i += 1
	}
}

// q1
func init() {
	var (
		pos Coord
		dir D
	)

	startPositionCharacters := []rune{'v', '<', '^', '>'}
	for i, row := range data {
		for j, val := range row {
			if slices.Contains(startPositionCharacters, val) {
				pos = Coord{
					row: i,
					col: j,
				}
				dir, err = directionFromRune(val)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}

	vis := make([][]bool, len(data))
	for i := range vis {
		vis[i] = make([]bool, len(data[i]))
	}

	m, n := len(data), len(data[0])
	for pos.IsInBounds(m, n) {
		vis[pos.row][pos.col] = true

		r := getNextRune(data, pos, dir)
		if r == nil {
			break
		}

		switch *r {
		case '#':
			dir = dir.Next()
		default:
			newPos := pos.StepInDirection(dir)
			pos = newPos
		}
	}

	for _, r := range vis {
		for _, v := range r {
			if v {
				res1 += 1
			}
		}
	}
}

// q2
func copySlice(a [][]rune) [][]rune {
	duplicate := make([][]rune, len(a))
	for i := range a {
		duplicate[i] = make([]rune, len(a[i]))
		copy(duplicate[i], a[i])
	}
	return duplicate
}

func printMatrix(m [][]rune) {
	for i := range m {
		a := ""
		for j := range m {
			a += string(m[i][j])
		}
		log.Println(a)
	}
	log.Println()
}

func init() {
	var (
		pos      Coord
		initialP Coord
		dir      D
		initialD D
	)

	startPositionCharacters := []rune{'v', '<', '^', '>'}
	for i, row := range data {
		for j, val := range row {
			if slices.Contains(startPositionCharacters, val) {
				pos = Coord{
					row: i,
					col: j,
				}
				dir, err = directionFromRune(val)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
	initialP = pos
	initialD = dir

	vis := make([][]bool, len(data))
	for i := range vis {
		vis[i] = make([]bool, len(data[i]))
	}

	m, n := len(data), len(data[0])
	for pos.IsInBounds(m, n) {
		vis[pos.row][pos.col] = true

		r := getNextRune(data, pos, dir)
		if r == nil {
			break
		}

		switch *r {
		case '#':
			dir = dir.Next()
		default:
			newPos := pos.StepInDirection(dir)
			pos = newPos
		}
	}

	for i, r := range vis {
	Outer:
		for j, v := range r {
			if !v {
				continue
			}

			dc := copySlice(data)
			if dc[i][j] == '#' {
				continue
			}
			dc[i][j] = '#'

			pos = initialP
			dir = initialD
			i := 0
			for pos.IsInBounds(m, n) {
				r := getNextRune(dc, pos, dir)
				if r == nil {
					continue Outer
				}

				switch *r {
				case '#':
					dir = dir.Next()
				default:
					newPos := pos.StepInDirection(dir)
					pos = newPos
				}

				if i == 10_000 {
					break
				}

				i += 1
			}

			res2 += 1
		}
	}
}

func main() {
	log.Println(res1)
	log.Println(res2)
}
