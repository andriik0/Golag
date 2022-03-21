package elon

import "fmt"

// TODO: define the 'Drive()' method
func (car *Car) Drive() {
	car.battery -= car.batteryDrain
	car.distance += car.speed
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (car Car) CanFinish(distance int) bool {
	if car.battery == 0 {
		return false
	}

	if distance <= 0 {
		return true		
	}

	return distance / car.speed * car.batteryDrain <= car.battery
}

// TODO: define the 'DisplayDistance() string' method
func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// TODO: define the 'DisplayBattery() string' method
func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}
