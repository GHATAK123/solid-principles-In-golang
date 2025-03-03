package principles

import "fmt"

// OCP - Open/Closed Principle
// - A class should be open for extension but closed for modification
// - Once a class is written, it should be closed for modification
// - We should be able to extend a class's behavior, without modifying it
// - We should be able to add new features to a class without changing its existing code

type Shape interface {
	Volume() float64
}

type Cubiod struct {
	Length, Breadth, Height float64
}

func (c Cubiod) Volume() float64 {
	return c.Length * c.Breadth * c.Height
}

type Cone struct {
	Radius, Height float64
}

func (c Cone) Volume() float64 {
	return (1.0 / 3.0) * 3.14 * c.Radius * c.Radius * c.Height
}

type Cylinder struct {
	Radius, Height float64
}

func (c Cylinder) Volume() float64 {
	return 3.14 * c.Radius * c.Radius * c.Height
}

type VolumeCalculator struct{}

func (vc VolumeCalculator) SumVolume(shapes []Shape) float64 {
	totalVolume := 0.0
	for _, shape := range shapes {
		totalVolume += shape.Volume()
	}
	return totalVolume
}

func OCP() {
	fmt.Println("From OCP Package!!!")
	cubiod := Cubiod{Length: 10, Breadth: 5, Height: 2}
	cone := Cone{Radius: 3, Height: 4}
	cylinder := Cylinder{Radius: 3, Height: 4}
	shapes := []Shape{cubiod, cone, cylinder}
	volumeCalculator := VolumeCalculator{}
	totalVolume := volumeCalculator.SumVolume(shapes)
	fmt.Printf("Total Volume: %.2f\n", totalVolume)
}

/*
✅ Benefits of This Approach
Extensible: Adding a new shape (e.g., Sphere) only requires defining a new struct implementing Shape, without modifying VolumeCalculator.
Maintains OCP: The VolumeCalculator class is closed for modification but open for extension.
Follows SRP (Single Responsibility Principle): Each class does only one job.
*/
