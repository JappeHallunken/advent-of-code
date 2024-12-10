package day10
import (
  "testing"
)

func TestDay10(t *testing.T) {
  input :=  "../../input/day10_test.txt"

  t.Run ("Day10", func(t *testing.T) {
    expecected := 36
    score1 := Day10(input)
    if score1 != expecected {
      t.Errorf("Day10() = %v, want %v", score1, expecected)
    }
  })
}
