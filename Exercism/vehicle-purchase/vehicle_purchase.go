package purchase

import (
	"fmt"
	"sort"
	"strings"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return strings.EqualFold(kind, "car") || strings.EqualFold(kind, "truck")
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	options := []string{option1, option2}
	sort.Strings(options)
	return fmt.Sprintf("%s is clearly the better choice.", options[0])
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	switch {
	case age < 3.0:
		return 0.8 * originalPrice
	case age >= 3.0 && age < 10.0:
		return 0.7 * originalPrice
	default:
		return 0.5 * originalPrice
	}
}
