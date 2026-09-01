// package lasagnamaster is an upgradet version of the lasagna package with
// extendet capabillitys
package lasagnamaster

// PreparationTime takes the layers as a slice and the time needet per Layer as an int
// to calculate the total time needed for preparattion. if provided time is 0, 2 is used
func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		timePerLayer = 2
	}
	return len(layers) * timePerLayer
}

// Quantities calculates how much noodles and sauce are needed, based on a given slice
// For every layer of noodles, we use 50g, for every layer of sauce 0.2L
func Quantities(layers []string) (int, float64) {
	var noodles int = 0
	var sauce float64 = 0

	for i := 0; i < len(layers); i++ {
		switch layers[i] {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}

	return noodles, sauce
}

func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	newQuantities := []float64{}
	for i := 0; i < len(quantities); i++ {
		newQuantities = append(newQuantities, (quantities[i]/2)*float64(portions))
	}
	return newQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
