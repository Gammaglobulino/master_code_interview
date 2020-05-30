package selection_sort

func Sort(vec []int) {
	for i := 0; i < len(vec)-1; i++ {
		min := vec[i]
		index := i
		for j := i + 1; j < len(vec); j++ {
			if vec[j] < min {
				min = vec[j]
				index = j
			}
		}
		vec[i], vec[index] = vec[index], vec[i]
	}
}
