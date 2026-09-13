package standardlibrary

import (
	"fmt"
	"errors"
)

var(
	ValidationError = errors.New("validation error")
	NotFoundError = errors.New("not found error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "Won-Young" {
		return NotFoundError
	}

	// sukses
	return nil
}

func ErrorFunctions() {
	err := GetById("ahyeon")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Not found error")
		} else {
			fmt.Println("Unknown error")
		}
	}
}
	

