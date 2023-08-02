package main

import "fmt"

//simple const set global//
const (
	x = iota
	y = iota
	z = iota
)

//you can write different way//

const (
	a = iota
	b
	c
)

// you can break the output

const (
	n = iota
	_
	m
	_
	s
)

func main() {
	// const i = 10 => you can not change const value

	//x//
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	//a//
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//a//
	fmt.Println(n)
	fmt.Println(m)
	fmt.Println(s)
}
