package main

import "fmt"

type User struct {
	name string
	age Age
}
// конструктор:
func NewUser(name string, age int) User {
	return User {
		name: name,
		age: Age(age),
	}
}
// метод
// принимает копию объекта
func (u User) printInfo() {
	println(u.name, u.age)
}
// принимает ссылку на объект
func (u *User) changeName(name string) {
	u.name = name
}


func (u User) getName() string {
	return u.name
}

func (u *User) setName(name string) {
	u.name = name
}


type Age int

func (a Age) isAdult() bool {
	return a >= 18
}



func main() {
	user := NewUser("Khasbulat", 22)

	// fmt.Println(user)
	fmt.Printf("user: %+v\n", user)

	user.changeName("Bayras")

	user.printInfo()

	fmt.Print(user.getName())

	user.setName("Khasbulat")

	fmt.Print(user.getName())
}