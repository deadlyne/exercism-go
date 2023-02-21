package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var sum int
	for i := 0; i < len(birdsPerDay); i++ {
		sum += birdsPerDay[i]
	}
	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
//week: 1 = 0*7
//week: 2 = 1*7
//week: 3 = 2*7
//week: 4 = 3*7
// i = (week-1)*7

// (0 1 2 3 4 5 6) (7 8 9 10 11 12 13) 14 15...

func BirdsInWeek(birdsPerDay []int, week int) int {
	var sum int
	for i := 0; i < 7; i++ {
		sum += birdsPerDay[i+(week-1)*7]
	}
	return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.

func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ { // для каждого четного индекса увеличивать значение элемента массива на 1
		if i%2 == 0 {
			birdsPerDay[i] = birdsPerDay[i] + 1
		}
	}
	return birdsPerDay
}

/*
	birdsPerDay[i]++

for _, element := range birdsPerDay {
		sum += element
	}
	return sum
	if week == 1 {
			sum += birdsPerWeek1
		} else {
			for _, birdsPerWeek1 := range birdsPerDay[14:] {
				sum += birdsPerWeek1
			}
		}
	}
*/
