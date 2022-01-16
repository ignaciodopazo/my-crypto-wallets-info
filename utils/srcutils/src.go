package srcutils

import (
	"log"
	"strconv"
)

func ParseFloat(toParse string) float64 {
	num, err := strconv.ParseFloat(toParse, 64)
	if err != nil {
		log.Fatal(err)
	}
	return num
}
