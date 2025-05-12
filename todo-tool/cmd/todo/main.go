package main

import (
	"fmt"
	"os"
	"flag"

	"todo-tool"
)

const todoFileName = ".todo.json"

func main() {

	// parsing command line flags
    task := flag.String("task", "", "Task to be included in the ToDo list")
	list := flag.Bool("list", false, "list all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")
	flag.Parse()

	l := &todo.List{}
    
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		// list current to do items
		for _,item := range *l {
			if !item.Done{
				fmt.Println(item.Task)
			}
		}
	case *complete > 0:
		// Complete the given item
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	    // save new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr,err)
			os.Exit(1)
		}
	case *task != "":
		// Add task
		l.Add(*task)

		// Save new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		// Invalid flag provided
		fmt.Fprintln(os.Stderr, "Invalid Operation")
		os.Exit(1)
	}
	
	
}