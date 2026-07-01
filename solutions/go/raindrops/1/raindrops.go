package raindrops

import (
	"strconv"
	"strings"
)

func Convert(number int) string {
	var builder strings.Builder

	if number%3 == 0 {
		builder.WriteString("Pling")
	}

	if number%5 == 0 {
		builder.WriteString("Plang")
	}

	if number%7 == 0 {
		builder.WriteString("Plong")
	}

	// if number%3 != 0 && number%5 != 0 && number%7 != 0 {
	// 	builder.WriteString(strconv.Itoa(number))
	// }

	if builder.Len() == 0 {
		builder.WriteString(strconv.Itoa(number))
	}

	return builder.String()
}
