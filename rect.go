package main


// Rect - "class" 
type Rect struct {
	Width, Height float64
	Qux string // Obriga o uso explicitamente "nomeado" {Width: 3, Height: 4} não permite {3, 4}
}

// Area of rectangle
func (r Rect) Area() float64 {
	return r.Width * r.Height
}

// Perim of rectangle
func (r Rect) Perim() float64 {
	return 2 * r.Width + 2 * r.Height
}