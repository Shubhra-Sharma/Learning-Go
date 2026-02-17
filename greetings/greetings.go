package greetings
// greetings module
import (
"fmt"
"errors"
"math/rand"
)

func Hello(name string) (string,error){
	if(name==""){
		return "", errors.New("No name given") // error handling
	}

   message := fmt.Sprintf(randomFormat(), name)
   return message, nil
}

// multiple greetings
func Hellos(names []string) (map[string]string, error) {
    messages := make(map[string]string)
    for _, name := range names {          // _ is the index as range returns an index, value of the current item in slice
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }
        messages[name] = message
    }
    return messages, nil
}
func randomFormat() string{
	 // A slice of message formats.
    formats := []string{
        "Hi, %v. Nice to meet you!",
        "Pleasure to meet you, %v!",
        "Hail, %v! Well met!",
    }
	return formats[rand.Intn(len(formats))]
}