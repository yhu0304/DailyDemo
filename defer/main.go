package main

import "fmt"

func add() (i int) {
	i = 10
	defer func() { i++ }()
	return
}

func main() {
	fmt.Println("result:", add())
}
