package main

import "fmt"

func main() {

	// name <-> grade
	studentGrades := make(map[string]int)

	studentGrades["Prince"] = 100
	studentGrades["Alice"] = 90
	studentGrades["Bob"] = 85
	studentGrades["Charlie"] = 95

	fmt.Println("Marks of Bob : ", studentGrades["Bob"])
	studentGrades["Bob"] = 100
	fmt.Println("new Marks of Bob : ", studentGrades["Bob"])

	delete(studentGrades, "Bob")
	fmt.Println("Marks of Bob : ", studentGrades["Bob"])

	// Checking if a key exists
	grades, exists := studentGrades["David"]
	fmt.Println("Grades of David : ", grades)
	fmt.Println("Davis exists : ", exists)

	// fmt.Println("Marks of David : ", studentGrades["David"])

	// Checking if a key exists
	Grades, Exists := studentGrades["Prince"]
	fmt.Println("Grades of Prince : ", Grades)
	fmt.Println("Prince exists : ", Exists)

	for index, value := range studentGrades {
		fmt.Printf("Key is %s and marks is %d\n", index, value)
	}

	person := map[string]int{
		"Alice":   90,
		"Bob":     85,
		"Charlie": 95,
	}

	for index, value := range person {
		fmt.Printf("---------Key is %s and marks is %d\n", index, value)
	}
}
