package utils

import (
	"math/rand"
	"time"
)

var runes = []rune("23456789abcdefghijklmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ")

func RandomURL(size int) string {
	str := make([]rune, size)

	rand.Seed(time.Now().UnixNano())
	for i := range str {
		str[i] = runes[rand.Intn(len(runes))]
	}

	return string(str)
}
