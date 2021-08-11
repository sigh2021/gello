package main

import "fmt"

type my_interf interface {
	my_method()
}

type insert_interf interface {
	test_method()
}

type my_struct struct {
	name string
    insertIf insert_interf
}

func (my *my_struct) my_method () {
	fmt.Println(my.name)
}

func newmy_struct(myname string) (*my_struct){
	return &my_struct {
		name: string(myname) }
}
//----------------------------
type test_struct struct {
    name string
}

func (test_struct) test_method () {
    fmt.Println("test test_method")
}

func main() {
    var interface1 my_interf
    struct1 := newmy_struct("hello")

    test := &test_struct {
        name: "test object" }
    struct1.insertIf = test
    struct1.insertIf.test_method()

    interface1 = struct1
    interface1.my_method()
}
