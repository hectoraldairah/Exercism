package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
  return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
  rate := successRate / 100.0;
  carsPerMinute := (rate * float64(productionRate)) / 60.0
  return int(carsPerMinute) 
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
  if carsCount == 0 {
    return 0
  }

  groupOfCars, restOfCars := carsCount / 10, carsCount % 10

  valueOfGroup := uint(groupOfCars * 95000)
  valueOfRest := uint(restOfCars * 10000)

  return valueOfGroup + valueOfRest
}
