package user

import "fmt"

// NewEncoder creates an encoder that can be passed to a saver
func NewEncoder(u *User) *Encoder {
	return &Encoder{u}
}

type Encoder struct {
	*User
}

func (e *Encoder) Encode() string {
	return fmt.Sprintf("%s - %v", e.name, e.age)
}
