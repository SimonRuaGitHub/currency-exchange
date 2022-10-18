package utils

import (
	"fmt"
	"strconv"
)

func FromStringToFloat(value string) (parsedValue float64, err error) {

	floatValue, err := strconv.ParseFloat(value, 64)

	if err != nil {
		fmt.Errorf("Error when parsing from string to float64", err)
		return 0.0, err
	} else {
		return floatValue, nil
	}
}
