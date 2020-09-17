package main

func main() {
	m := Mixbonacci([]string{"fib"}, 10)
	//m := Mixbonacci([]string{"fib", "tet"}, 10)
	//m := Mixbonacci([]string{"jac", "jac", "pel"}, 10)
	//m := Mixbonacci([]string{
	//	"tet", "fib", "pad", "tri", "fib", "jac", "tri", "jac", "tet", "fib", "tet", "fib", "pel", "fib", "pel", "pel"}, 13)
	for i, v := range m {
		println(i, v)
	}
}

func getFactors(patterns []string) []int {
	r := make([]int, 0)
	m := map[string]int{}
	for _, v := range patterns {
		if val, ok := m[v]; ok {
			val++
			m[v] = val
		} else {
			m[v] = 0
		}
		r = append(r, m[v])
	}
	return r
}

func flatPattern(pattern []string, length int) []string {
	flat := make([]string, 0)
	if length < len(pattern) {
		for idx, p := range pattern {
			if idx < length {
				flat = append(flat, p)
			}
		}
		return flat
	}
	r := length / len(pattern)
	rr := length % len(pattern)
	n := 0
	for n < r {
		for _, p := range pattern {
			flat = append(flat, p)
		}
		n++
	}
	m := 0
	for m < rr {
		for idx, p := range pattern {
			if idx == m {
				flat = append(flat, p)
			}
		}
		m++
	}
	return flat
}

func Mixbonacci(pattern []string, length int) []int64 {
	r := make([]int64, 0)
	if len(pattern) == 0 || length == 0 {
		return r
	}
	flatMix := flatPattern(pattern, length)
	factors := getFactors(flatMix)
	for idx, p := range flatMix {
		m := MixMap[p](factors[idx])
		r = append(r, m)
	}
	return r
}

type Mnacci func(i int) int64

var MixMap = map[string]Mnacci{
	"fib": fib,
	"pad": padovan,
	"pel": pell,
	"jac": jacobsthal,
	"tri": tribonacci,
	"tet": tetranacci,
}

func fib(n int) int64 {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func pell(n int) int64 {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return 2*pell(n-1) + pell(n-2)
}

func jacobsthal(n int) int64 {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return jacobsthal(n-1) + 2*jacobsthal(n-2)
}

func padovan(n int) int64 {
	if n == 0 {
		return 1
	}
	if n == 1 || n == 2 {
		return 0
	}
	return padovan(n-2) + padovan(n-3)
}

func tribonacci(n int) int64 {
	if n == 0 || n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	return tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3)
}

func tetranacci(n int) int64 {
	if n == 0 || n == 1 || n == 2 {
		return 0
	}
	if n == 3 {
		return 1
	}
	return tetranacci(n-1) + tetranacci(n-2) + tetranacci(n-3) + tetranacci(n-4)
}
