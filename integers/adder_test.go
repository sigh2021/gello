package integers

import "testing"
import "fmt"

func TestAdder(t *testing.T) {
    sum := Add(2,2)
    expect := 4
    if sum != expect {
        t.Errorf(" sum '%d' but expect '%d'",sum ,expect)
    }
}

func ExampleAdd() {
    sum := Add(1, 5)
    fmt.Println(sum)
    //output: 6
}
