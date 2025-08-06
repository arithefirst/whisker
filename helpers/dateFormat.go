package helpers

import (
	"fmt"
	"time"
)

func RFC3339toDateString(inputDateString string) (string, error) {

	t, err := time.Parse(time.RFC3339, inputDateString)
	if err != nil {
		fmt.Printf("Error parsing date: %v\n", err)
		return "", err
	}

	day := t.Day()
	suffix := "th"

	switch day % 10 {
	case 1:
		if day != 11 {
			suffix = "st"
		}
	case 2:
		if day != 12 {
			suffix = "nd"
		}
	case 3:
		if day != 13 {
			suffix = "rd"
		}
	}

	return fmt.Sprintf("%s %d%s, %d",
		t.Format("January"),
		day,
		suffix,
		t.Year(),
	), nil
}
