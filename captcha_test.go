package captcha_test

import (
	"testing"

	"github.com/zesani/captcha"
)

func TestCaptcha5PlusTwo(t *testing.T) {
	expected := "5 + two"

	result := captcha.Captcha(1, 5, 1, 2)

	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestCaptcha4MinusTwo(t *testing.T) {
	expected := "4 - two"

	result := captcha.Captcha(1, 4, 2, 2)

	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestCaptcha2MutiOne(t *testing.T) {
	expected := "2 x one"

	result := captcha.Captcha(1, 2, 3, 1)

	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}
