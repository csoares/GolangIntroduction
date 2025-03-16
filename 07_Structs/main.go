package main

import (
	"fmt"
	"math"
)

// Person demonstrates basic struct declaration
type Person struct {
	Name    string
	Age     int
	Address string
}

// Student demonstrates struct composition
type Student struct {
	Person    // Embedded struct
	ID        string
	Grades    []float64
	GradeMean float64
}

// Circle demonstrates struct with methods
type Circle struct {
	Radius float64
}

// Rectangle demonstrates another shape struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Method for Student to calculate grade average
func (s *Student) CalculateGradeMean() {
	sum := 0.0
	for _, grade := range s.Grades {
		sum += grade
	}
	s.GradeMean = sum / float64(len(s.Grades))
}

// Method for Circle to calculate area
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Method for Rectangle to calculate area
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	// Basic struct usage
	person := Person{
		Name:    "John Doe",
		Age:     20,
		Address: "123 Main St",
	}
	fmt.Printf("Person: %+v\n", person)

	// Student management example
	student := Student{
		Person: Person{
			Name:    "Jane Smith",
			Age:     19,
			Address: "456 College Ave",
		},
		ID:     "CS101",
		Grades: []float64{85.5, 92.0, 88.5, 95.0},
	}
	student.CalculateGradeMean()
	fmt.Printf("\nStudent: %+v\n", student)
	fmt.Printf("Grade Mean: %.2f\n", student.GradeMean)

	// Geometry calculator example
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}

	fmt.Printf("\nCircle Area: %.2f\n", circle.Area())
	fmt.Printf("Rectangle Area: %.2f\n", rectangle.Area())
}
