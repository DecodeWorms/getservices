package pkg

import (
	"math/rand"
	"strings"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz1234567890"
const stringLength = 10

func GenerateUserId(userFirstTwoChar string) string {
	rand.Seed(time.Now().UnixNano())

	//strings.Builder{} is used to concatenate character and avoid copy
	//and also optimize memo allocatio
	sb := strings.Builder{}

	//Grow grows the string builder
	sb.Grow(stringLength)
	for i := 0; i < stringLength; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return userFirstTwoChar + sb.String()
}

func GetUserFirstTwoChar(firstName string) string {
	var result string
	for index, character := range firstName {
		if index == 3 {
			break
		}
		result += string(character)
	}
	return strings.ToLower(result) + "-"
}

func ParseCategory(cat string)bool{
	for _ , category := range ServiceCategory{
		if cat == category{
			return true
		}
	}
	return false
}
