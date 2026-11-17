package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Returns greeting for the named person
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("Empty Name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil

}

func randomFormat() string {

	formats := []string{
		"hi, %v,Welcome!",
		"Great to see you, %v",
		"Hail, %v!",
	}
	//Return a randomly selected message format by specifying a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
