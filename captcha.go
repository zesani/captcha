package captcha

import (
	"fmt"
	"strconv"
)

var NumberToWord = map[int]string{
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

func Captcha(p, n1, op, n2 int) string {
	str1 := NumberToWord[n1]
	str2 := strconv.Itoa(n2)
	if p == 1 {
		str1 = strconv.Itoa(n1)
		str2 = NumberToWord[n2]
	}

	switch op {
	case 1:
		return fmt.Sprintf("%v + %v", str1, str2)
	case 2:
		return fmt.Sprintf("%v - %v", str1, str2)
	}
	return fmt.Sprintf("%v x %v", str1, str2)
}
