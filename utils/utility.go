package utils

import (
	"fmt"
	"math/rand"
	"strings"
)

func SplitBabString(babString string) []string {
	babSlice := strings.Split(babString, ";")

	var cleanedBabSlice []string
	for _, bab := range babSlice {
		if bab != "" {
			cleanedBabSlice = append(cleanedBabSlice, bab)
		}
	}

	return cleanedBabSlice
}

func GenerateId() string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letter := make([]byte, 3)
	for i := range letter {
		letter[i] = letters[rand.Intn(len(letters))]
	}
	num := rand.Intn(10000)
	orderString := fmt.Sprintf("%s-%04d", letter, num)

	return orderString
}
