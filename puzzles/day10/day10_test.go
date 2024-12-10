package day10
import (
  "testing"
)

func TestDay10(t *testing.T) {
  input :=  "../../input/day10_test.txt"

  t.Run ("Day10", func(t *testing.T) {
    Day10(input)
  })
}
