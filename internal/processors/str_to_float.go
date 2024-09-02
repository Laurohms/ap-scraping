package processors

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func StrToFloat(value string) (float64, error) {
	validChars := "1234567890+-,. "

	if value == "" {
		log.Printf("Value is empty")
		return 0, errors.New("value is empty")
	}

	for _, char := range value {
		if !strings.ContainsRune(validChars, char) {
			log.Printf("Invalid character found: %c", char)
			return 0, fmt.Errorf("invalid character: %c", char)
		}
	}
	value = strings.ReplaceAll(value, "+", "")
	value = strings.ReplaceAll(value, ".", "")
	value = strings.ReplaceAll(value, ",", ".")
	value = strings.ReplaceAll(value, " ", "")
	value = strings.TrimSpace(value)
	converted, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Printf("Erro parsin string to float: %v", err)
	}
	return converted, nil
}
