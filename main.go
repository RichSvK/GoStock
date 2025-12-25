package main

import (
	"fmt"

	"github.com/RichSvK/Stock_Holder_Composition_Go/config"
	"github.com/RichSvK/Stock_Holder_Composition_Go/helper"
	"github.com/RichSvK/Stock_Holder_Composition_Go/utilities"
)

func main() {
	config.MakeFolder("output")

	db := utilities.LoginMenu()

	// Close database when the main function end
	defer func() {
		if err := db.Close(); err != nil {
			fmt.Println("Failed to close database connection:", err)
		}
	}()

	choice := 0
	for choice != 3 {
		choice = utilities.MainMenu()
		switch choice {
		case 1:
			utilities.InsertMenu()
		case 2:
			utilities.ExportMenu()
		default:
			fmt.Println("Program finished")
			return
		}
		helper.PressEnter()
		helper.ClearScreen()
	}
}
