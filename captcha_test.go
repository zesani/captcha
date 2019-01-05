package captcha_test

import (
	"testing"

	"github.com/zesani/captcha"
)

func TestCaptchaFivePlusTwo(t *testing.T) {
	expected := "5 + two"

	result := captcha.Captcha(1, 5, 1, 2)

	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestCaptchaFourMinusTwo(t *testing.T) {
	expected := "4 - two"

	result := captcha.Captcha(1, 4, 2, 2)

	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}
