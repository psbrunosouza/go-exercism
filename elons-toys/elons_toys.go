package elon

import "fmt"

// Drive allow car be driven, reduce the battery and sum the distance performed
func (car *Car) Drive() Car {
	if car.battery > 0 && car.battery >= car.batteryDrain {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
	return *car
}

// DisplayDistance display the performed distance by car
func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// DisplayBattery display the amout of battery remaining
func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

// CanFinish define if a car can finish the race
func (car Car) CanFinish(trackDistance int) bool {
	carDrivenLimit := car.battery / car.batteryDrain
	distanceCarCanCover := car.speed * carDrivenLimit

	return distanceCarCanCover >= trackDistance
}
