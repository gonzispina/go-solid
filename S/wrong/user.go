package wrong

import (
	"fmt"
	"os"
	"strconv"
)

// NewUser creates a user
func NewUser(name string, age int) User {
	return User{
		name: name,
		age:  age,
	}
}

// User model
type User struct {
	name string
	age  int
}

// Name returns the user's name
func (u *User) Name() string {
	return u.name
}

// Age returns the user's age
func (u *User) Age() int {
	return u.age
}

// Save saves the implementation to the os
func (u *User) Save() bool {
	f, err := os.Create(u.name)
	if err != nil {
		return false
	}

	defer f.Close()

	l, err := f.WriteString(fmt.Sprintf("%s%s", u.name, strconv.Itoa(u.age)))
	if err != nil {
		return false
	}

	fmt.Println(l, "bytes written successfully")
	return true
}