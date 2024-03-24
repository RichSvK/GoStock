package utility

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/RichSvK/Stock_Holder_Composition_Go/database"
)

func LoginMenu() *sql.DB {
	var (
		username string
		password string
		dbName   string
	)

	fmt.Print("Insert username: ")
	fmt.Scanln(&username)
	fmt.Print("Insert password: ")
	fmt.Scanln(&password)
	fmt.Print("Insert Database name: ")
	fmt.Scanln(&dbName)
	return database.GetConnection(username, password, dbName)
}

func MainMenu() int {
	var userInput string = ""
	var choice int = 0
	var err error = nil
	ClearScreen()
	fmt.Println("Main Menu")
	fmt.Println("1. Insert Data to Database")
	fmt.Println("2. Export Data to Database")
	fmt.Println("3. Exit")
	for {
		fmt.Print("Input[1 - 3]: ")
		fmt.Scanln(&userInput)
		choice, err = strconv.Atoi(userInput)
		if err == nil {
			if choice >= 1 && choice <= 3 {
				break
			}
		}
	}
	return choice
}

func ExportMenu(db *sql.DB) {
	var code string = ""
	for len(code) != 4 {
		fmt.Print("Input stock name: ")
		fmt.Scanln(&code)
	}
	database.Export(code, db)
}

func InsertMenu(db *sql.DB) {
	ClearScreen()
	fmt.Println("Menu Insert")

	directory := "Data/"  // Change this to the directory you want to list files from
	var fileList []string // Initialize an empty slice to hold file paths
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			fileList = append(fileList, path) // Add the file path to the slice
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var size int = len(fileList)
	var tempName []string = nil
	for i := 0; i < size; i++ {
		tempName = strings.Split(fileList[i], "Data\\")
		fmt.Printf("%d. %s from %s\n", (i + 1), tempName[1], fileList[i])
	}

	var choice int = 0
	var userInput string = ""
	for {
		fmt.Printf("Input [1 - %d]: ", size)
		fmt.Scanln(&userInput)
		choice, err = strconv.Atoi(userInput)
		if err == nil {
			if choice > 0 && choice < 4 {
				break
			}
		} else {
			fmt.Println("Invalid Input")
		}
	}
	database.InsertData(db, fileList[choice-1])
}

func ClearScreen() {
	// Clear screen for Unix systems
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		// Clear screen for Windows
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
