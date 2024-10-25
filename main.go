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
	for _, task := range taskItems { //blank to explicit unused variable as in range loop need to provide index,task
		fmt.Println(task)
	}

}
