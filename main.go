package main

import (
	"fmt"

	"github.com/RichSvK/Stock_Holder_Composition_Go/config"
	"github.com/RichSvK/Stock_Holder_Composition_Go/helper"
	"github.com/RichSvK/Stock_Holder_Composition_Go/utilities"
)

func init() {
	config.MakeFolder("output")
	utilities.LoginMenu()
}

func main() {
	// Close database when the main function end
	defer config.PoolDB.Close()

	var choice int = 0
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
