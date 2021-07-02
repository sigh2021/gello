package main

import "testing"

//call methods
func TestPerimeter(t *testing.T){
    t.Run("rectangles", func (t *testing.T) {
        a := rectangles{10.0, 10.0}
        got := rectangles.Area(a)
        want := 50.0
        if got != want {
            t.Errorf("got '%2f' but want '%2f' ",got ,want)
        }
    })

    t.Run("circle", func (t *testing.T) {
        a := circle{10.0}
        got := circle.Area(a)
        want := 314.0
        if got != want {
            t.Errorf("got '%2f' but want '%2f' ",got ,want)
        }
    })
}

//call interface
func TestShapes(t *testing.T) {
    var shape shapes

    check_Area := func(t *testing.T, shape shapes, want float64) {
        t.Helper()
        got := shape.Area()
        if got != want {
            t.Errorf("got '%2f' but want '%2f' ",got ,want)
        }
    }
    t.Run("rectangles", func (t *testing.T) {
        shape = rectangles{10.0, 10.0}
        want := 50.0
        check_Area(t, shape, want)
    })

    t.Run("circle", func (t *testing.T) {
        shape = circle{10.0}
        want := 314.0
        check_Area(t, shape, want)
    })
}

//table driven tests
func TestTable(t *testing.T) {
    areaTest := []struct  {
        shape shapes
        want float64
    }{
        {rectangles{10.0, 10.0}, 50.0},
        {circle{10.0}, 314.0},
    }

    for _,tt := range areaTest {
        got := tt.shape.Area()
        if got != tt.want {
            t.Errorf("got '%2f' but want '%2f' ",got ,tt.want)
        }
    }
}
