package lasagna


// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timeToLayer int) int  {
	if timeToLayer < 1 {
		timeToLayer = 2
	}
	return len(layers)*timeToLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	sauceCount := 0
	noodleCount := 0
	for _, layer := range layers{
		switch layer {
		case "sauce":{
			sauceCount++
		}
	case "noodles":{
		noodleCount++
	}
			
		} 
			}
	return noodleCount * 50, float64(sauceCount) * 0.2
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) []string {
	return append(append(make([]string, 0), myList...), friendsList[len(friendsList)-1])
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portionQuantities int) []float64  {
	result := make([]float64, 0)
	for _, val := range quantities {
		result = append(result, val * float64(portionQuantities)/2.0)
	}
return result
}