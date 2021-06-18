package main

import "testing"

func TestHello(t *testing.T) {
    got := Hello("Crist")
    want := "Hello Crist"
    if got != want {
        t.Errorf("got '%q' want '%q'", got, want)
    }
}
