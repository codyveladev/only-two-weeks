package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

// Interface Def
type Shape interface {
	Area() float64
	Perimeter() float64
}

type InvalidDimensionError struct {
	Shape string
	Value float64
}

// We satisify the internal error interface, not needed to declare a new one
func (e InvalidDimensionError) Error() string {
	return e.Shape + " cannot have negative dimension " + strconv.FormatFloat(e.Value, 'f', -1, 64)
}

var ErrNegativeDimension = errors.New("negative dimension")

// Any type that implements interface methods satisfies the interface
// implicit implementation
type Circle struct {
	Radius float64
}

// Interface methods there for is a shape?
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func NewCircle(radius float64) (Circle, error) {
	if radius < 0 {
		return Circle{}, fmt.Errorf("circle: %w %w", ErrNegativeDimension, InvalidDimensionError{Shape: "Circle", Value: radius})
	}
	return Circle{Radius: radius}, nil
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return (2 * r.Width) + (2 * r.Height)
}

func NewRectangle(width float64, height float64) (Rectangle, error) {
	if width < 0 {
		return Rectangle{}, fmt.Errorf("rectangle: %w %w", ErrNegativeDimension, InvalidDimensionError{Shape: "Rectangle", Value: width})
	}
	if height < 0 {
		return Rectangle{}, fmt.Errorf("rectangle: %w %w", ErrNegativeDimension, InvalidDimensionError{Shape: "Rectangle", Value: height})
	}
	return Rectangle{
		Width:  width,
		Height: height,
	}, nil

}

type Triangle struct {
	A float64
	B float64
	C float64
}

func (t Triangle) Area() float64 {
	s := (t.A + t.B + t.C) / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

func NewTriangle(a float64, b float64, c float64) (Triangle, error) {
	if a < 0 {
		return Triangle{}, fmt.Errorf("triangle: %w %w", ErrNegativeDimension, InvalidDimensionError{Shape: "Triangle", Value: a})
	}
	if b < 0 {
		return Triangle{}, fmt.Errorf("triangle: %w %w", ErrNegativeDimension, InvalidDimensionError{Shape: "Triangle", Value: b})
	}
	if c < 0 {
		return Triangle{}, fmt.Errorf("triangle: %w %w", ErrNegativeDimension, InvalidDimensionError{Shape: "Triangle", Value: c})
	}
	return Triangle{
		A: a,
		B: b,
		C: c,
	}, nil
}

func printArea(s Shape) {
	fmt.Println(s.Area())
}

func printInfo(shapes []Shape) {
	for _, s := range shapes {
		switch v := s.(type) {
		case Circle:
			if v.Area() == 0 {
				continue
			}
			fmt.Printf("circle — radius: %.2f, area: %.2f\n", v.Radius, v.Area())
		case Rectangle:
			if v.Area() == 0 {
				continue
			}
			fmt.Printf("rectangle — width: %.2f height: %.2f, area: %.2f, perimeter:%.2f\n", v.Width, v.Height, v.Area(), v.Perimeter())
		case Triangle:
			if v.Area() == 0 {
				continue
			}
			fmt.Printf("triangle — sides: %.2f %.2f %.2f, area: %.2f, perimeter:%.2f\n", v.A, v.B, v.C, v.Area(), v.Perimeter())
		default:
			fmt.Println("unknown shape")
		}
	}
}

func main() {
	r1, r1err := NewRectangle(4.0, 4.0)
	c1, c1err := NewCircle(-5.5)
	t1, t1err := NewTriangle(3.0, 4.0, 5.0)
	var dimErr InvalidDimensionError

	if errors.Is(r1err, ErrNegativeDimension) {
		fmt.Println("caught with is", r1err)
	}

	if errors.As(c1err, &dimErr) {
		fmt.Println(c1err)
	}
	if errors.As(t1err, &dimErr) {
		fmt.Println(t1err)
	}

	// How does Go know a type satififies an interface
	// Checks at compile time that a type has the methods the interface it requires
	// Again, this is implicit implementation
	// printArea(r1)
	// printArea(c1)

	// Type Assertions & type switches
	shapes := []Shape{}
	if c1err == nil {
		shapes = append(shapes, c1)
	}
	if r1err == nil {
		shapes = append(shapes, r1)
	}
	if t1err == nil {
		shapes = append(shapes, t1)
	}
	printInfo(shapes)

}
