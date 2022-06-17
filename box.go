package golang_united_school_homework

import (
	"fmt"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if b.shapesCapacity < 1 {
		return fmt.Errorf("insufficient capacity")
	}
	if len(b.shapes) < b.shapesCapacity {
		b.shapes = append(b.shapes, shape)
		return nil
	} else {
		return fmt.Errorf("Shape is not added, there is not enough capacity")
	}
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if len(b.shapes) != 0 && i < len(b.shapes) {
		return b.shapes[i], nil
	} else {
		return nil, fmt.Errorf("Shape does not exist")
	}

}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if len(b.shapes) != 0 && i < len(b.shapes) {
		v := b.shapes[i]
		copy(b.shapes[i:], b.shapes[i+1:])
		b.shapes[len(b.shapes)-1] = nil
		b.shapes = b.shapes[:len(b.shapes)-1]
		return v, nil
	} else {
		return nil, fmt.Errorf("Shape does not exist")
	}

}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i < len(b.shapes) {
		v := b.shapes[i]
		b.shapes[i] = shape
		return v, nil
	} else {
		return nil, fmt.Errorf("index does not exist")
	}

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	sum := 0.0
	for i := 0; i < len(b.shapes); i++ {
		sum += Shape.CalcPerimeter(b.shapes[i])
	}
	if len(b.shapes) == 0 {
		return 0
	}
	return sum

}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	sum := 0.0
	for i := 0; i < len(b.shapes); i++ {
		sum += Shape.CalcArea(b.shapes[i])
	}
	if len(b.shapes) == 0 {
		return 0
	}
	return sum

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	if len(b.shapes) < 1 {
		return fmt.Errorf("no shapes")
	}
	notCircle := make([]Shape, 0)
	for i := range b.shapes {
		_, ok := b.shapes[i].(*Circle)
		if !ok {
			notCircle = append(notCircle, b.shapes[i])
		}
	}
	if len(b.shapes) == len(notCircle) {
		return fmt.Errorf("Circle not found")
	}
	b.shapes = notCircle
	return nil

}
