package main

import "testing"

func TestHello(t *testing.T) {

    assertCorrectMessage:= func(t *testing.T, got, want string) {
        //声明辅助函数
        t.Helper()
        if got != want {
            t.Errorf("got '%q' want '%q'", got, want)
        }
    }

    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Crist", "English")
        want := "Hello Crist"
        assertCorrectMessage(t, got, want)
    })

    t.Run("saying hello world when empty", func(t *testing.T) {
        got := Hello("", "")
        want := "Hello World"
        assertCorrectMessage(t, got, want)
    })

    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Crist", "Spanish")
        want := "Hola Crist"
        assertCorrectMessage(t, got, want)
    })

}
