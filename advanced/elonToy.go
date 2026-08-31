package advanced

import "strconv"

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func (car *Car) Drive() {
	if car.battery >= car.batteryDrain {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
}

func (car Car) DisplayDistance() string {
	return "Driven " + strconv.Itoa(car.distance) + " meters"
}

func (car Car) DisplayBattery() string {
	return "Battery at " + strconv.Itoa(car.battery) + "%"
}

// CanFinish checks if a car is able to finish a certain track.
func (car Car) CanFinish(track int) bool {
	dist := car.distance + car.speed*(car.battery/car.batteryDrain)
	return dist >= track
}
