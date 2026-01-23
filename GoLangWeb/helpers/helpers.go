package helpers

import "math/rand"

func RandNum(n int) int {

	
	value := rand.Intn(n)
	return value
}
