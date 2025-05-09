package main

import (
	"fmt"
	"os"
	"strings"

	"todo-tool"
)

const todoFileName = ".todo.json"

func main() {
	l := &todo.List{}
    
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case len(os.Args) == 1:
		for _,item := range *l {
			fmt.Println(item.Task)
		}
	default:
		item := strings.Join(os.Args[1:]," ")
		l.Add(item)

		// Save new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr,err)
			os.Exit(1)
		}
	}
	
	
}