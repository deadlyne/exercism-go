package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		return len(layers) * (time + 2)
	} else {
		return len(layers) * time
	}
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		} else if layer == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	result := make([]float64, len(quantities))
	for i := 0; i < len(result); i++ {
		result[i] = quantities[i] * float64(portions) / 2
	}
	return result
}

/*
quantities*portions/2


	coast1 := 0
	coast2 := 0
	for _, ingredients := range layers {
		if ingredients == la {
			coast1 ++
		} else if sauc == ingredients {
			coast2 ++
		}
	}
	return (noodles * coast1), (sauce * float64(coast2))
// TODO: define the 'Quantities()' function
	result1 := 0
	result2 := 0
	for _, layer := range layers {
		if layer == "noodles" {
			result1++
		} else if layer == "sauce" {
			result2++
		}
	}
	return (noodles * result1), (sauce * float64(result2))


elem1 := 0
	elem2 := 0
	for i, layers = "noodles"{
		i++
		elem1 += i
		} for i, layers = "sauce"{
			i++
			elem2 += i
		}
		return (noodles * result1), (sauce * float64(result2))


	elem1 := 0
	elem2 := 0
	for layers == "noodles" {
		elem1++
		fmt.printf(elem1)
	}
	return elem1
}

func Quantities(layers []string) (int, float64) {

}
*/
