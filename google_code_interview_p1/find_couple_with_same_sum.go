package google_code_interview_p1

import "log"

func FindCoupleWithSameSum(a []int, sum int) bool {
	complementOfA := make(map[int]bool, len(a))
	for _, v := range a {
		if complementOfA[v] {
			return true
		}
		complementOfA[sum-v] = true
		log.Println(complementOfA)
	} //N
	return false
}
