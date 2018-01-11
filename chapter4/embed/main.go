package main

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Radius int
}

func main() {
	w := Wheel{Circle{Point{8, 8}, 5}, 20}

}
