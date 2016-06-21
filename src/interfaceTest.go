package main

import "fmt"
import "math"

type geometry interface {
	area() float64
}
type rect struct {
	width, height float64
}

type circle struct {
	raidus float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.raidus * c.raidus
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
}

func measureMyData() {
	r := rect{width: 3, height: 4}
	c := circle{raidus: 5}
	measure(r)
	measure(c)
}
