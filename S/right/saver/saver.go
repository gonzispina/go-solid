package saver

import (
	"fmt"
	"os"
)

type Saver struct {}

// Save saves the implementation to the os
func (s *Saver) Save(fileName, data string) bool {
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
