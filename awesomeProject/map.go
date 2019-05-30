package main

import (
	"fmt"
	"time"
	"math"
	"math/rand"
)

func main() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["Li"] = "lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"

	if name, ok := elements["B"]; ok{
		fmt.Println(name, ok)
	}
	//fmt.Println(elements["C"])
	//another way of declaring a map
	_ = map[string]string{
		"H": "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
	}

	xs := []float64{43,34,34,54}
	fmt.Println(average(xs))
	f()
	x, y := m()
	fmt.Println(x, y)
	fmt.Println(add(1,2,3,4))

	nextEven := makeEvenNumbers()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	fmt.Println(factorial(4))

	c := Circle{0,0,5}
	fmt.Println(area(&c))
	fmt.Println(c.area())

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())

	fmt.Println(totalArea(&c, &r))

	for i := 0; i < 10 ; i++  {
		go w(0)
	}
	var input string
	fmt.Scanln(&input)
}

func average(xs []float64) float64{
	total := 0.0
	for _, v := range xs  {
		total += v
	}

	return total /float64(len(xs))
}

func f(){
	fmt.Println("cooking")
}

func m() (int, int) {
	return 5, 6
}

func add(args ...int) int {
	total := 0

	for _,v := range args  {
		total += v
	}
	return total
}

func qd() {
	add := func() int {
		return 1 + 1
	}

	fmt.Println(add())
}

func inc() {
	x := 0

	increment := func () int{
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

func makeEvenNumbers() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func factorial(x uint) uint{
	if x == 0{
		return 1
	}

	return x * factorial(x -1)
}

func distance(x1, x2,y1, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)

	return l * w
}



func (m *MultiShapes) area() float64  {
	var area float64
	for _, v := range m.shapes {
		area =+ v.area()
	}
	return area
}

func area(c *Circle) float64{
	return math.Pi * c.r * c.r
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, v := range shapes  {
		area =+ v.area()
	}
	return area
}

func (c *Circle) area() float64 {
	return math.Pi * c.r*c.r
}

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, x2, y1, y2 float64
}

type MultiShapes struct {
	shapes []Shape
}

type Shape interface {
	area() float64
	//perimeter() float64
}


func w(n int) {
	for i :=0; i < 10 ; i++  {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}