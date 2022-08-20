package main

func main() {
	var t []int
	//t = append(t, 3, 4, 7, 4, 6, 5, 8, 2, 1, 4, 5, 11, 55, 44, 34, 1, 2, 3, 4)
	//Bubble sort - Worst case performance O(n2)
	//printT(bubbleSort(t))

	//Insertion sort- as bad as the bubble O(n2) - best case O(n)
	//t = append(t, 3, 4, 7, 4, 6, 5, 8, 2, 1, 4, 5, 11, 55, 44, 34, 1, 2, 3, 4)
	//printT(insertSort(t))

	//t = append(t, 3, 4, 7, 4, 6, 5, 8)

}

func printT(t []int) {
	for _, i := range t {
		print(i, ",")
	}
}

func bubbleSort(t []int) []int {
	t, s := bubble(t)
	for s {
		t, s = bubble(t)
	}
	return t
}

func bubble(t []int) ([]int, bool) {
	s := 0
	for idx := range t {
		if idx >= len(t)-1 {
			return t, s != 0
		}
		if t[idx+1] < t[idx] {
			swap(t, idx)
			s++
		}
	}
	return t, s != 0
}

func swap(t []int, idx int) {
	var n = t[idx]
	var m = t[idx+1]
	t[idx] = m
	t[idx+1] = n
}

func insertSort(t []int) []int {
	for idx, i := range t {
		newIdx := findInsertIndex(t, i, idx)

		var m = t[newIdx]
		t[newIdx] = i
		t[idx] = m

	}
	return t
}

func findInsertIndex(t []int, cand int, candIdx int) int {
	max := 0
	for idx, i := range t {
		if (idx < len(t)-1) && idx < candIdx && i > cand && cand < t[idx+1] {
			max = idx

		}
	}

	return max
}

func mergeSort(t []int) (t1 []int, t2 []int) {
	return nil, nil
}

func halve(t []int) (t1 []int, t2 []int) {
	hp := len(t) / 2
	return t[0:hp], t[hp:(len(t))]
}