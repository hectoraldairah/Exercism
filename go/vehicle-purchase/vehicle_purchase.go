package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
  return kind == "car" || kind == "truck" 
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
  sentence := " is clearly the better choice."
  if option1 > option2 {
    return option2 + sentence
  } else {
    return option1 + sentence
  }
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
  var percentage float64
  if age < 3 {
    percentage = 0.80
  } else if age >= 10 {
    percentage = 0.50
  } else {
    percentage = 0.70
  }

  return (originalPrice * percentage)
}
