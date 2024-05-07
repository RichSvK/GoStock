package utilities

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/RichSvK/Stock_Holder_Composition_Go/config"
	"github.com/RichSvK/Stock_Holder_Composition_Go/helper"
	"github.com/RichSvK/Stock_Holder_Composition_Go/service"
)

func LoginMenu() {
	var (
		username string = ""
		password string = ""
		dbName   string = ""
	)

	var db *sql.DB = nil
	for db == nil {
		fmt.Println("Login Menu")
		username = helper.ScanInput("Insert username: ")
		password = helper.ScanInput("Insert password: ")
		dbName = helper.ScanInput("Insert Database name: ")
		db = config.GetConnection(username, password, dbName)
		helper.PressEnter()
		helper.ClearScreen()
	}
}

func MainMenu() int {
	var userInput string = ""
	var choice int = 0
	var err error = nil
	helper.ClearScreen()
	fmt.Println("Main Menu")
	fmt.Println("1. Insert Data to Database")
	fmt.Println("2. Export Data to Database")
	fmt.Println("3. Exit")
	for {
		userInput = helper.ScanInput("Input[1 - 3]: ")
		choice, err = strconv.Atoi(userInput)
		if err == nil && choice >= 1 && choice <= 3 {
			break
		}
	}
	return choice
}

func ExportMenu() {
	var code string = ""

	for len(code) != 4 {
		code = helper.ScanInput("Input stock name [4 Letter]: ")
	}
	service.Export(code)
}

func InsertMenu() {
	helper.ClearScreen()
	fmt.Println("Menu Insert")

	// Change this to the directory you want to list files from
	directory := "data/"

	// Initialize an empty slice to hold file paths
	var fileList []string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			// Add the file path to the slice
			fileList = append(fileList, path)
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
		tempName = strings.Split(fileList[i], "data\\")
		fmt.Printf("%d. %s from %s\n", (i + 1), tempName[1], fileList[i])
	}

	var choice int = 0
	var userInput string = ""
	promptString := fmt.Sprintf("Input [1 - %d]: ", size)
	for {
		userInput = helper.ScanInput(promptString)
		choice, err = strconv.Atoi(userInput)
		if err == nil && choice >= 1 && choice <= size {
			break
		} else {
			fmt.Println("Invalid Input")
		}
	}
	service.InsertData(fileList[choice-1])
}
