package main

import "fmt"

// int string bool Person Contact Address
// var name Person
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct {
	Email string
	Phone string
	Fax   string
}

type Address struct {
	House int
	Area  string
	State string
}

type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}

func main() {
	var prince Person
	// fmt.Println("Prince Person : ", prince) // " " " " 0 // __0
	prince.FirstName = "Prince"
	prince.LastName = "Agarwal"
	// fmt.Println("Prince Person : ", prince)

	// 2nd method
	person1 := Person{
		FirstName: "Aakash",
		LastName:  "Singh",
		Age:       26,
	}
	fmt.Println("Person 1 : ", person1)

	// new keyword
	var person2 = new(Person)
	person2.FirstName = "Simran"
	person2.LastName = "Agarwal"
	person2.Age = 26

	// fmt.Println("Person 2 : ", person2.FirstName)

	// fmt.Println("Age of Prince is ", prince.Age)

	var employee1 Employee
	employee1.Person_Details = Person{
		FirstName: "Prince",
		LastName:  "Agarwal",
		Age:       26,
	}
	employee1.Person_Contact.Email = "contact@gmail.com"
	employee1.Person_Contact.Phone = "98765432"

	employee1.Person_Address = Address{
		House: 12,
		Area:  "Ranchi",
		State: "Jharkhand",
	}
	employee1.Person_Contact.Fax = "Fax@4567890"
	fmt.Println("Employee 1 : ", employee1.Person_Contact)
}
