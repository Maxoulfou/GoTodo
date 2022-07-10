/*
Copyright © 2022 Brochier Maxence maxence@brochier.xyz

*/
package execution

import (
	"fmt"
	"github.com/Maxoulfou/GoTodo/utilities"
	"os"
)

const TodoFile = ".TodoFile.json"

func AddTodo() {
	todos := &utilities.Todos{}

	if err := todos.Load(TodoFile); err != nil {
		fmt.Println("load file")
		fmt.Fprintln(os.Stderr, err.Error())

		os.Exit(1)
	}

	//---

	task := utilities.GetInput(os.Args[2:])

	todos.Add(task)
	err := todos.Store(TodoFile)

	if err != nil {
		_, err := fmt.Fprintln(os.Stderr, err.Error())
		if err != nil {

			return
		}

		os.Exit(1)
	}
}

func CompleteTodo() {
	todos := &utilities.Todos{}

	if err := todos.Load(TodoFile); err != nil {
		fmt.Println("load file")
		fmt.Fprintln(os.Stderr, err.Error())

		os.Exit(1)
	}

	// ---

	err := todos.Complete(utilities.GetIntInput(os.Args[2:]))
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())

		os.Exit(1)
	}
	err = todos.Store(TodoFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())

		os.Exit(1)
	}
}

func DeleteTodo() {
	todos := &utilities.Todos{}

	if err := todos.Load(TodoFile); err != nil {
		fmt.Println("load file")
		fmt.Fprintln(os.Stderr, err.Error())

		os.Exit(1)
	}

	// ---

	err := todos.Delete(utilities.GetIntInput(os.Args[2:]))
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())

		os.Exit(1)
	}
	err = todos.Store(TodoFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())

		os.Exit(1)
	}
}

func ListTodo() {
	todos := &utilities.Todos{}

	if err := todos.Load(TodoFile); err != nil {
		fmt.Println("load file")
		fmt.Fprintln(os.Stderr, err.Error())

		os.Exit(1)
	}

	// ---

	todos.Print()
}
