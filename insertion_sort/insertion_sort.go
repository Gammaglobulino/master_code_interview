package insertion_sort

func Sort(vec []int) {
	for i := 0; i < len(vec)-1; i++ {
		if vec[i] > vec[i+1] {
			vec[i], vec[i+1] = vec[i+1], vec[i]
			for j := i - 1; j <= 0; j-- {
				if vec[j] > vec[i] {
					vec[j], vec[i] = vec[i], vec[j]
				}
			}
		}
	}
}
