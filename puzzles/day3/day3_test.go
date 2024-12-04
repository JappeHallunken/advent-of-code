package day3

import (
	"testing"

)
func MockReadFile() []byte {
	return []byte("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")
}

func MockReadFile2() []byte {
	return []byte("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))")
}

func TestCalculateSum(t *testing.T) {
	t.Run("TestCalculateSum", func(t *testing.T) {
		expected := 161

		body := MockReadFile()

		actual, _ := calculateSum(body)

		if actual != expected {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	})
}

func TestMakeString(t *testing.T) {
	t.Run("TestMakeString", func(t *testing.T) {
		body := MockReadFile2()

		newString := makeString(body)

		actual, _ := calculateSum([]byte(newString))

		expected := 48

		if actual != expected {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	})
}
