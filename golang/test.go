package main
import "fmt"
func main()  {
    // 动态类型
    var b Bird
    // b 是main.Sparrow类型
    b = Sparrow{}
    fmt.Printf("%T\n", b)
    // b 是main.Parrot类型
    b = Parrot{}
    fmt.Printf("%T\n", b)
}
type Bird interface {
    fly()
    sing()
}
// Goose implement Bird interface
type Sparrow struct {
    age  int
    name string
}
func (s Sparrow) fly()  {
    fmt.Println("I am flying.")
}
func (s Sparrow) sing()  {
    fmt.Println("I can sing.")
}
type Parrot struct {
    age  int
    kind int
    name string
}
func (p Parrot) fly()  {
    fmt.Println("I am flying.")
}
func (p Parrot) sing()  {
    fmt.Println("I can sing.")
}
