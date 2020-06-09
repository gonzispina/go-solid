package user

// New creates a user
func New(name string, age int) User {
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

