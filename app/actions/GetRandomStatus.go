package actions

import (
	"math/rand"
	"time"
)

func randInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

func GetNewOrErrorStatus() int {
	if randInt(0, 100) > 50 {
		return 1
	}
	return 4
}
