package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timePerLayer int)int{
    if timePerLayer == 0{
        return len(layers) * 2
    }else{
        return len(layers) * timePerLayer
    }
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string)(int, float64){
    var noodlesNeeded int
    var sauceNeeded float64
    for i := 0; i < len(layers); i++{
        switch layers[i]{
        case "noodles": 
            noodlesNeeded += 50
        case "sauce": 
            sauceNeeded += 0.2
    	}
    }
    return noodlesNeeded, sauceNeeded
}
// TODO: define the 'AddSecretIngredient()' function    
func AddSecretIngredient(friendsList, myList []string) {
    myList[len(myList)-1] = friendsList[len(friendsList)-1]
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantitiesForTwo []float64, numberOfPortions int) []float64{
    var quantitiesNeeded = []float64 {}
    for i := 0; i < len(quantitiesForTwo); i++{
        quantitiesNeeded = append(quantitiesNeeded, quantitiesForTwo[i]*float64(numberOfPortions)/2.0)
    }
    return quantitiesNeeded
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
