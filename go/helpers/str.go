package helpers

import (
	"fmt"
	"strconv"
)

func ParseInt(num string) int {
	i, err := strconv.Atoi(num)
	if err != nil {
		fmt.Print("Error converting %s to int\n", num)
	}
	return i
}