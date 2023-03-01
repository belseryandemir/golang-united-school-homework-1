package main

import "fmt"

type User struct {
    firstName string
    lastName  string
}

func New() *User {
    return &User{}
}

func (u *User) SetFirstName(firstName string) {
    u.firstName = firstName
}

func (u *User) SetLastName(lastName string) {
    u.lastName = lastName
}

func (u *User) FullName() string {
    return fmt.Sprintf("%s %s", u.lastName, u.firstName)
}

func ResetUser(input interface{}) {
    if user, ok := input.(*User); ok {
        user.firstName = ""
        user.lastName = ""
    }
}

func IsUser(input interface{}) bool {
    _, ok := input.(*User)
    return ok
}

func ProcessUser(input UserInterface) string {
    input.SetFirstName("Jane")
    input.SetLastName("Doe")
    return input.FullName()
}

type UserInterface interface {
    SetFirstName(string)
    SetLastName(string)
    FullName() string
}

func main() {
    user := New()
    fmt.Println(ProcessUser(user)) // Doe Jane
    ResetUser(user)
    fmt.Println(IsUser(user)) // true
}