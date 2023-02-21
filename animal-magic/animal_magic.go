package chance

import (
	"math/rand"
	"strings"
	"time"
)

// SeedWithTime seeds math/rand with the current computer time.
func SeedWithTime() {
	rand.Seed(time.Now().Unix())
}

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	n := rand.Intn(19) + 1
	return n
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	f := rand.Float64() * 12
	return f
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	animals := strings.Fields("ant beaver cat dog elephant fox giraffe hedgehog")
	rand.Shuffle(len(animals), func(i, j int) {
		animals[i], animals[j] = animals[j], animals[i]
	})
	return animals
}

/*
n := rand.Intn(10) + 1
	f := float64(n) + rand.Float64()
	return f
*/
