package main

import "fmt"


func main() {

	users := map[string]int {
		"Khasbulat": 22,
		"Masha": 18,
	}

	// fmt.Println(users)

	age, exists := users["Khasbulat"]
	if exists {
		fmt.Println(age)
	}

	age, exists = users["fasdf"]
	if exists {
		fmt.Println(age)
	} else {
		fmt.Println("user doesn't exist", age)
	}
	// fmt.Println(users["Khasbulat"])
	// fmt.Printf("users: %v\n", users)

	users["Larisa"] = 47
	users["Noname"] = 999
	for key, value := range users {
		println(key, value)
	}

	// удаление
	delete(users, "Noname")
	for key, value := range users {
		println(key, value)
	}

}