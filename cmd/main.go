package main

import (
	"fmt"
	"game-library-management-system/src/app"
)

func main() {
	a, err := app.NewApp()
	if err != nil {
		fmt.Println("Error creating app", err)
	}
	err = a.Run()
	if err != nil {
		fmt.Println("Error running app", err)
	}
}
