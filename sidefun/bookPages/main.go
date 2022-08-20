package main

import (
	"fmt"
	"math"
)

func main() {
	//c := pagesCount(82902)
	//fmt.Println(c)

	p := count(63547)
	fmt.Println(p)
	c := pagesCount(5)
	fmt.Println(c)

	fmt.Println("\r\r")
	c = pagesCount(25)
	fmt.Println(c)

	c = pagesCount(1095)
	fmt.Println(c)

	c = pagesCount(63547)
	fmt.Println(c)

	c = pagesCount(660)
	fmt.Println(c)

}

func pagesCount(summary int) int {
	return counter(0, 0, summary)
}

func counter(page int, currentCount int, max int) int {
	lmax := int(math.Log10(float64(max)) + 1)
	numberOfDigits := 1
	if page > 9 {
		numberOfDigits = int(math.Log10(float64(page)) + 1)
	}
	n := currentCount + numberOfDigits

	if page <= 0 {
		page = 1
	}
	if n <= max {
		nextPage := page + 1
		return counter(nextPage, n, max)
	} else {
		if numberOfDigits >= lmax-1 {
			return page - 1
		}
		return page
	}
}

func count(max int) int {
	c := make([]int, 0)
	p := 0
	for page := 1; page <= max; page++ {
		l := int(math.Log10(float64(page)) + 1)
		c = append(c, l)
		//fmt.Println("SUM:", c)
		if sum(c) <= max {
			p = page

		}
	}
	return p
}

func sum(chunks []int) int {
	s := 0
	for _, v := range chunks {
		s = s + v
	}
	return s
}
