package two_sum

import "errors"

//this is not an optimal solution, even if really elegant and idiomatic
//the best is using the incremental complement shown in the Google code example

type valueIndex struct {
	value int
	index int
}
type valueIndexArray struct {
	valueIndexContainer []valueIndex
}

func (ai *valueIndexArray) IndexArray() []int {
	var out []int
	for _, v := range ai.valueIndexContainer {
		out = append(out, v.index)
	}
	return out
}
func TwoSum(a []int, sum int) ([]int, error) {
	out := valueIndexArray{}

	for i, v := range a {
		if len(out.valueIndexContainer) != 0 {
			if v < sum {
				if v+out.valueIndexContainer[0].value == sum {
					out.valueIndexContainer = append(out.valueIndexContainer, valueIndex{
						value: v,
						index: i,
					})
					return out.IndexArray(), nil
				}
			}
		} else {
			if v < sum {
				out.valueIndexContainer = append(out.valueIndexContainer, valueIndex{
					value: v,
					index: i,
				})
			}
		}

	}
	return nil, errors.New("No indexes found")
}
