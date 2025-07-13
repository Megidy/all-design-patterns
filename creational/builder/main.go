package main

import "log"

func main() {
	u := NewUser().WithName("hello").WithEmail("hello email").WithAge(31)
	log.Printf("user: %+v", u)
}

type User struct {
	Name  string
	Email string
	Age   int
}

func NewUser() *User {
	return &User{}
}

func (u *User) WithName(name string) *User {
	u.Name = name
	return u
}

func (u *User) WithEmail(email string) *User {
	u.Email = email
	return u
}

func (u *User) WithAge(age int) *User {
	u.Age = age
	return u
}
