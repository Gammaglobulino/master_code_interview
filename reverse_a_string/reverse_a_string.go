package reverse_a_string

import "strings"

func ReverseString(str string) string {
	reversed := strings.Builder{}
	for i := len(str) - 1; i >= 0; i-- {
		reversed.WriteString(string(str[i]))
	}
	return reversed.String()
}
