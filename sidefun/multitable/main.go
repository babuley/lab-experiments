package main

import "fmt"

func main() {
	t := table(4)
	fmt.Println(t)
}

func table(size int) [][]int {
	r := make([][]int, size)

	for i := range r {
		table := make([]int, 0)
		for j := 1; j <= size; j++ {
			table = append(table, (i+1)*j)
		}
		r[i] = table

	}
	return r
}
