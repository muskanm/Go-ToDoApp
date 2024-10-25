package main

import "fmt"

func main() {
	// variables
	codingTask := "Practice Coding Problems"
	behavioralTask := "Prepare Answers for Behavioral Questions"
	researchTask := "Research the Company and Role"

	//array and slices
	var taskItems = []string{codingTask, behavioralTask, researchTask}

	fmt.Println("##### Welcome to our Todo List App! #####")

	fmt.Println("List of my Todos")
	for index, task := range taskItems { //blank to explicit unused variable as in range loop need to provide index,task
		fmt.Printf("%d. %s\n", index+1, task)
	}

}
