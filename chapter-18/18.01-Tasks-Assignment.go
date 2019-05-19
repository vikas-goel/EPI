// We consider the problem of assigning tasks to workers. Each worker must be
// assigned exactly two tasks. Each task takes a fixed amount of time. Tasks
// are independent, i.e., there are no constraints of the form "Task 4 cannot
// start before Task 3 is completed." Any task can be assigned to any worker.
// We want to assign tasks to workers so as to minimize how long it takes
// before all tasks are completed.
// Design an algorithm that takes as input a set of tasks and returns an
// optimum assignment.

package main

import (
	"fmt"
	"sort"
)

type Task = int

func sortTasks(tasks []Task) {
	sort.Ints(tasks)
}

func Assignment (tasks []Task) [][]Task {
	// Sort the tasks in ascending hours.
	sortTasks(tasks)

	numTasks := len(tasks)

	// For odd number of tasks, let the last task be lone assignment in
	// the pair.
	oddFactor := 0
	if numTasks % 2 == 1 {
		oddFactor = 1
	}

	numAssgn := (numTasks+oddFactor)/2
	pair := make([][]Task, numAssgn)

	// Pair the two extremes and keep moving in.
	for i := 0; i < numTasks/2; i++ {
		pair[i] = make([]Task, 2)
		pair[i][0] = tasks[i]
		pair[i][1] = tasks[numTasks-i-1-oddFactor]
	}

	// If the number of given tasks is odd, then the last assignment gets
	// only the last task in the sorted one.
	if oddFactor == 1 {
		pair[numAssgn-1] = make([]Task, 1)
		pair[numAssgn-1][0] = tasks[numTasks-1]
	}

	return pair
}

func main() {
	tasks := []Task{5, 2, 1, 6, 4, 4}
	fmt.Print("Tasks", tasks, " = Assignment")
	fmt.Println(Assignment(tasks))
}
