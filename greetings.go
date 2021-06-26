// this represents the package name. In go a module may have multiple packages
// the module may also have packages that depend on other packages
// in this file, we are creating a simple greeting package that can be used in another
// go program to display a greeting
package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Here we are declaring a function with a string return type.
// The function also takes in a string value as a parameter.
// NB: In go, variable types are declared after variable declaration
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("missing name parameter")
	}

	// In go, using the := symbol is a short cut for declaring and instantiatin a variable
	// the longer war is to do this: var message string.
	// Notice in the long hand we declare the variable type. with the shorthand approach,
	// the variable will inherit the type of the expression to it's right

	// message := fmt.Sprintf("Hello there, %v. Welcome to the Go program", name)
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// function to greet multiple people. A new function was created for backward compatibility
func GreetAll(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

// init() functions are always run when the go program starts. after global variables
// have been initialised.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Notice this function starts with a small letter? In Go, any function that starts
// with a small letter is not exported and hence, is only accessible within it's own
// package. Kinda like protected and private functions in other languages.
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
