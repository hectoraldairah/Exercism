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
	return Car{
		battery:      100,
		batteryDrain: batteryDrain,
		speed:        speed,
	}
}

// TODO: define the 'Track' type struct

type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	newDistance := car.speed + car.distance
	newBaterry := car.battery - car.batteryDrain

	if newBaterry < 0 {
		return car
	}

  car.distance = newDistance
  car.battery = newBaterry

	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
  requiredBattery := track.distance / car.speed * car.batteryDrain
  return car.battery >= requiredBattery 
}
