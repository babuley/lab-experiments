package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	n := 2071
	s := strconv.Itoa(n)

	t := make([]int, 0)
	for _, c := range strings.Split(s, "") {
		cc, _ := strconv.ParseInt(c, 10, 0)
		t = append(t, int(cc))
	}
	fmt.Println(t)
	bucket := permutation(t)

	result := make([]int, 0)
	for _, chunk := range bucket {
		ss := ""

		for idx, v := range chunk {

			if idx == 0 && v == 0 {
				continue
			}
			ss = ss + strconv.Itoa(v)
		}
		sss, _ := strconv.Atoi(ss)
		result = append(result, sss)
	}
	sort.Ints(result)

	rr := -1
	for idx, no := range result {

		if n == no {
			rr = result[idx-1]
		}
	}
	fmt.Println(rr)

}

func permutation(xs []int) (permuts [][]int) {
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			permuts = append(permuts, append([]int{}, a...))
		} else {
			for i := k; i < len(xs); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(xs, 0)

	return permuts
}
