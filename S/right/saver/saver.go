package saver

import (
	"fmt"
	"os"
)

type Encoder interface {
	Encode() string
}

// Save saves the implementation to the os
func SaveToFile(fileName string, data Encoder) bool {
	f, err := os.Create(fileName)
	if err != nil {
		return false
	}

	defer f.Close()

	l, err := f.WriteString(data)
	if err != nil {
		return false
	}

	fmt.Println(l, "bytes written successfully")
	return true
}
