package utils

import (
	"math/rand"
	"strconv"
	"strings"
)

func GenerateCode() string {
	var code strings.Builder
	for i := 0; i < 6; i++ {
		code.WriteString(strconv.Itoa(rand.Intn(10)))
	}
	return code.String()
}
