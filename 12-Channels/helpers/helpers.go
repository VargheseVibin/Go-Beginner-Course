package helpers

import (
	"math/rand"
	"time"
)

func GetRandomNumber(n int) int {
	rand.Seed(time.Now().UnixNano())
	rNum := rand.Intn(n)
	return rNum
}
