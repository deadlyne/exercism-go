package speed

// TODO: define the 'Car' type struct

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{battery: 100, speed: speed, batteryDrain: batteryDrain, distance: 0}
}

// TODO: define the 'Track' type struct

type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	var distanceup int = (car.speed + car.distance)
	if car.battery <= 0 || car.battery < car.batteryDrain {
		car = Car{car.battery, car.batteryDrain, car.speed, car.distance}
		return car
	} else {
		car = Car{car.battery - car.batteryDrain, car.batteryDrain, car.speed, distanceup}
		return car
	}
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	var cardistance int = (car.battery / car.batteryDrain) * car.speed
	track = Track{track.distance}
	car = Car{car.battery, car.batteryDrain, car.speed, car.distance}
	return track.distance <= cardistance
}

/*
:= Car{car.battery, car.batteryDrain, car.speed, car.distance}
	battery:      100,
	speed:        5,
	batteryDrain: 2,
	distance:     0,
*/
