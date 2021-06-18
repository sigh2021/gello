package iteration

import "testing"
import "fmt"

func TestItera(t *testing.T) {
    want := "aaaaa"
    got := itera("a",5)
    if want != got {
        t.Errorf("expected %s but got %s", want, got)
    }
}

func BenchmarkItera(b *testing.B) {
    for i:= 0; i < b.N; i++ {
        itera("a", 5)
    }
}

func ExampleItera() {
    result := itera("b", 5)
    fmt.Println(result)
    //output: bbbbb
}
