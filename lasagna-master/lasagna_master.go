package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timePerLayer int) int {
	totalPreparationTime := 0

	for i := 0; i < len(layers); i++ {
		totalPreparationTime += 1
	}

	if timePerLayer != 0 {
		totalPreparationTime *= timePerLayer
	} else {
		totalPreparationTime *= 2
	}

	return totalPreparationTime
}

// TODO: define the 'Quantities()' function
func Quantities(ingredients []string) (noddlesGrams int, sauceLiters float64) {
	for i := 0; i < len(ingredients); i++ {
		if ingredients[i] == "noodles" {
			noddlesGrams += 50
		} else if ingredients[i] == "sauce" {
			sauceLiters += 0.2
		}
	}

	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) []string {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
	return myList
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, numberOfPortions int) []float64 {
	var scaledQuantities []float64

	for i := 0; i < len(quantities); i++ {
		scaledQuantities = append(scaledQuantities, quantities[i]*(float64(numberOfPortions)/2))
	}

	return scaledQuantities
}
