package main

import (
	"fmt"
)

type geometry interface {
	sumer() float64
}

type sumes struct {
	first, second float64
}

func (r sumes) sumer() float64 {
	return r.first + r.second
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.sumer())
}

func main() {
	r := sumes{first: 3, second: 4}
	measure(r)
}
