package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(strs []string) ([]float64, error) {
	var floats []float64

	for _, lineValue := range strs {
		price, err := strconv.ParseFloat(lineValue, 64)
		if err != nil {
			return nil, errors.New("failed convert strings to floats")

			// or just return the error for more spesific error information
			// return nil,err
		}

		floats = append(floats, price)
	}

	return floats, nil
}
