package fileops

import (
	"errors"
	"fmt"
	"os"
)

func WriteToFile(fileName string, value float64) error {
	valueInStirng := fmt.Sprint(value)
	error := os.WriteFile(fileName, []byte(valueInStirng), 0644)
	if error != nil {
		panic("can't write to file, Fuck!!!!!")
	}
	return nil
}

func ReadFromFile(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return "", errors.New("Failed to fetch balace")
	} else {
		value := string(data)
		return value, err
	}

}
