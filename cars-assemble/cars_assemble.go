package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	x := float64(productionRate)     // изменяем значение int у productionRate на float64 для совместимости
	return ((x * successRate) / 100) // вычисляет, сколько машин успешно производится в час
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(((successRate * float64(productionRate) / 100) / 60))
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	z := uint(carsCount)
	if z <= 9 {
		return (z * 10000)
	} else if z == 10 {
		return (z * 9500)
	} else if z >= 11 {
		return ((z / 10) * 95000) + ((z % 10) * 10000)
	}
	return (0)
}

/*
x := int(successRate)
	return ((z%10)*10000 + (z/10)*9500)

	z := uint(carsCount)
	if z <= 9 {
		return z * 10000
	}
	if z >= 10 {
		return (z%10)*10000 + ((z/10)%10)*10
	}
}
*/
/* z := uint(carsCount)
result1:= z
fmt.Printf("Result of z = %d", result1)

result2:= z * 10000
fmt.Printf("Result of z * 10000 = %d", result2)

fmt.Println()
*/
