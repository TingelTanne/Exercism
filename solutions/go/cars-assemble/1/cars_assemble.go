package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var workingCarsPerHour float64 = (float64(productionRate) * successRate) / 100
	return workingCarsPerHour
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var workingCarsPerMinute int = int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
	return workingCarsPerMinute
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var totalCost uint = uint(carsCount/10*95000 + carsCount%10*10000)
	return totalCost
}
