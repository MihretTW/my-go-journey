package main

import "fmt"

var studName string
var numSubjects int
var subjectsGrade = make(map[string]float64)

func askName() {
	fmt.Println("Enter your name: ")
	fmt.Scan(&studName)
	fmt.Println("Hello, " + studName + "!")
}

func askNumberOfSubjects() {
	fmt.Println("Enter the number of subjects: ")

	fmt.Scan(&numSubjects)

	if numSubjects <= 0 {
		fmt.Println("Invalid number of subjects. Please enter a positive integer.")
		askNumberOfSubjects()
	}

}

func averageGrade() float64 {
	if len(subjectsGrade) == 0 {
		return 0
	}

	var total float64
	for _, grade := range subjectsGrade {
		total += grade
	}
	return total / float64(len(subjectsGrade))
}

func main() {
	askName()
	askNumberOfSubjects()

	for i := 1; i <= numSubjects; i++ {
		var subject string
		var grade float64

		fmt.Printf("Enter subject %d name: ", i)
		fmt.Scan(&subject)

		if _, exists := subjectsGrade[subject]; exists {
			fmt.Println("Subject already entered.")
			i--
			continue
		}

		fmt.Println("Enter the grade you received:  ")
		fmt.Scan(&grade)

		if grade < 0 || grade > 100 {
			fmt.Println("Invalid grade. Please enter a grade between 0 and 100.")
			i--
			continue
		}

		subjectsGrade[subject] = grade
	}

	fmt.Println("Student Name: " + studName)
	fmt.Println("Number of Subjects: " + fmt.Sprint(numSubjects))
	fmt.Println("Grades:")

	for subject, grade := range subjectsGrade {
		fmt.Println("- " + subject + ": " + fmt.Sprint(grade))
	}

	fmt.Println("Average Grade: " + fmt.Sprint(averageGrade()))
}
