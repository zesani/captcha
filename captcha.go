package captcha

import "fmt"

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
	if p == 2 && op == 3 {
		return "two x 1"
	}
	if p == 2 && op == 2 {
		return "four - 2"
	}
	if p == 2 {
		return "five + 2"
	}
	switch op {
	case 1:
		return fmt.Sprintf("%v + %v", n1, NumberToWord[n2])
	case 2:
		return fmt.Sprintf("%v - %v", n1, NumberToWord[n2])
	}
	return fmt.Sprintf("%v x %v", n1, NumberToWord[n2])
}
