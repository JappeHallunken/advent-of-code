package day10
import (
  "testing"
)

func TestDay10(t *testing.T) {
  input :=  "../../input/day10_test.txt"

    score1, score2 := Day10(input)
  t.Run ("Day10 puzzle 2", func(t *testing.T) {
    expecected := 36
    if score1 != expecected {
      t.Errorf("Day10() = %v, want %v", score1, expecected)
    }
  })
  t.Run ("Day10 puzzle 1", func(t *testing.T) {
    expecected := 81
    if score2 != expecected {
      t.Errorf("Day10() = %v, want %v", score1, expecected)
    }
  })
}
