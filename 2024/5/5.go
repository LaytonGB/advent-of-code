package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type O struct {
	a int
	b int
}

var (
	tree             map[int][]int
	updates          [][]int
	incorrectUpdates [][]int
	res1             int
	res2             int
)

// get data
func init() {
	file, err := os.Open("./5/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var orders []O
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			break
		}
		parts := strings.Split(t, "|")
		numParts := O{}
		numParts.a, _ = strconv.Atoi(parts[0])
		numParts.b, _ = strconv.Atoi(parts[1])
		orders = append(orders, numParts)
	}

	tree = make(map[int][]int)
	for _, o := range orders {
		if _, ok := tree[o.a]; !ok {
			tree[o.a] = make([]int, 0)
		}
		tree[o.a] = append(tree[o.a], o.b)
	}

	for scanner.Scan() {
		t := scanner.Text()
		parts := strings.Split(t, ",")
		var numParts []int
		for _, p := range parts {
			n, _ := strconv.Atoi(p)
			numParts = append(numParts, n)
		}
		updates = append(updates, numParts)
	}
}

// q1
func init() {
Outer:
	for _, up := range updates {
		vis := make(map[int]struct{})
		for _, u := range up {
			vis[u] = struct{}{}
			for _, c := range tree[u] {
				if _, ok := vis[c]; ok {
					incorrectUpdates = append(incorrectUpdates, up)
					continue Outer
				}
			}
		}
		res1 += up[len(up)/2]
	}
}

// q2
func sortByTree(a, b int) int {
	if c, ok := tree[a]; ok && slices.Contains(c, b) {
		return 1
	}
	return -1
}

func init() {
	for _, up := range incorrectUpdates {
		slices.SortFunc(up, sortByTree)
		res2 += up[len(up)/2]
	}
}

func main() {
	log.Println(res1)
	log.Println(res2)
}
