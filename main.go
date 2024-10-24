package main

import "fmt"

func main() {
	// variables
	codingTask := "Practice Coding Problems"
	behavioralTask := "Prepare Answers for Behavioral Questions"
	researchTask := "Research the Company and Role"

	//array and slices
	var taskItems = [20]string{codingTask, behavioralTask, researchTask}

	fmt.Println("##### Welcome to our Todo List App! #####")

	fmt.Println("List of my Todos")
	fmt.Println(taskItems[0])
	fmt.Println(taskItems[1])
	fmt.Println(taskItems[2])

	fmt.Println()
	fmt.Println("Technical Interviews")
	fmt.Println(taskItems[0])
	fmt.Println(taskItems[1])

	fmt.Println()
	fmt.Println("Final Interview")
	fmt.Println(taskItems[2])

}
