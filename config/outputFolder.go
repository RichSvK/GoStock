package config

import (
	"fmt"
	"os"
)

func MakeFolder(folderName string) {
	// Check if there is a "output" directory in current directory
	_, checkFolder := os.Stat("./" + folderName)

	// If checkFolder != nil means there is no "output" directory in current directory
	for checkFolder != nil {
		err := os.Mkdir("./"+folderName, 0755)
		if err != nil {
			fmt.Println("Error creating directory")
		}
		_, checkFolder = os.Stat("./" + folderName)
	}
}
