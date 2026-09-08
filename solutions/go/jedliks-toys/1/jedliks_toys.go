package jedlik

import "fmt"

func (car *Car) Drive() {
	if car.battery >= car.batteryDrain {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
}

func (car *Car) DisplayDistance() string {
	return fmt.Sprint("Driven ", car.distance, " meters")
}

func (car *Car) DisplayBattery() string {
	return fmt.Sprint("Battery at ", car.battery, "%")
}

func (car *Car) CanFinish(trackDistance int) bool {
	maxDistance := car.battery / car.batteryDrain * car.speed
	return trackDistance <= maxDistance
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
