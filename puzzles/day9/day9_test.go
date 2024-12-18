package day9

import (
	"testing"
)


func TestDay9(t *testing.T) {

  ex1, ex2 := Day9("../../input/day9_test.txt")
  // fmt.Println(ex1)

  t.Run("test 1", func(t *testing.T) {
    expected := 1928
    if ex1 != expected {
      t.Errorf("Day9() = %v, want %v", ex1, expected)
    }
  })
  t.Run("test 2", func(t *testing.T) {
    expected := 2858
    if ex2 != expected {
      t.Errorf("Day9() = %v, want %v", ex2, expected)
    }
  })
   

}

