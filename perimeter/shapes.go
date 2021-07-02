package main

type shapes interface {
    Area() float64
    //area_rectangles float64
    //area_circle float64
}

type rectangles struct {
    width float64
    length float64
}
type circle struct {
    radius float64
}
//this kind of define methods is kind of intresting.
func (a rectangles) Area() float64 {
    return a.width * a.length /2
}
func (c circle) Area() float64 {
    return c.radius * c.radius * 3.14
}
