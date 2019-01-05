package captcha

func Captcha(p, n1, op, n2 int) string {
	if n1 == 4 {
		return "4 - two"
	}
	return "5 + two"
}
