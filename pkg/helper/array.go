package helper

import (
	"golang.org/x/exp/rand"
	"time"
)

func RandomArray[T any](array []T, count int) []T {
	rand.Seed(uint64(time.Now().UnixNano()))

	shuffled := make([]T, len(array))

	rand.Shuffle(len(array), func(i, j int) {
		shuffled[i], shuffled[j] = array[j], array[i]
	})

	if count >= len(shuffled) {
		count = len(shuffled)
	}

	return shuffled[:count]
}
