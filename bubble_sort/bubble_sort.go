package bubble_sort

func Sort(vec []int) {
	for i := 0; i < len(vec)-1; i++ {
		for j := i + 1; j < len(vec); j++ {
			if vec[j] < vec[i] {
				vec[i], vec[j] = vec[j], vec[i]
			}
		}
	}
}
