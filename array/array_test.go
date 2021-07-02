package myarray


import "testing"
import "reflect"

func TestArraySum(t *testing.T) {
    numbers := []int{1,2,3,4,5}
    want := 15
    got := ArraySum(numbers)
    if got != want {
        t.Errorf("expected %d but got %d numbers %v", want, got, numbers)
    }
}

func TestSumAll(t *testing.T) {
    numbers := []int{1,2,3}
    numbers1 := []int {6,9,}
    want := []int{6,15}
    got := SumAll(numbers,numbers1)
    if !reflect.DeepEqual(want, got) {
        t.Errorf("expected %v but got %v numbers %v numbers1 %v", want, got, numbers, numbers1)
    }
}
