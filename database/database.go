package database

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/RichSvK/Stock_Holder_Composition_Go/models"
	_ "github.com/go-sql-driver/mysql"
)

func GetConnection(username string, password string, dbName string) *sql.DB {
	poolDB, err := sql.Open("mysql", username+":"+password+"@tcp(localhost:3306)/"+dbName+"?parseTime=true")
	if err != nil {
		return nil
	}

	// Check if the connection to database is alive
	// If user inserted wrong password, username, or database name, err != nil
	err = poolDB.Ping()
	if err != nil {
		fmt.Println("Failed to connect")
		return nil
	}

	poolDB.SetMaxIdleConns(10)
	poolDB.SetMaxOpenConns(100)
	poolDB.SetConnMaxIdleTime(3 * time.Minute)
	poolDB.SetConnMaxLifetime(60 * time.Minute)
	fmt.Println("Success make connection")
	return poolDB
}

func Export(code string, poolDB *sql.DB) {
	ctx := context.Background()
	sql_query := "SELECT * FROM Stocks WHERE `Code` = ? ORDER BY `Date`"
	statement, err := poolDB.PrepareContext(ctx, sql_query)
	if err != nil {
		fmt.Println("Fail to export because", err.Error())
		return
	}
	defer statement.Close()

	rows, err := statement.QueryContext(ctx, code)
	if err != nil {
		fmt.Println("Fail to export because", err.Error())
		return
	}
	defer rows.Close()

	if !rows.Next() {
		fmt.Println("No stock with code:", code)
		return
	}

	// Check if there is a "Output" directory in current directory
	_, checkFolder := os.Stat("./Output")

	// If checkFolder != nil means there is no "Output" directory in current directory
	if checkFolder != nil {
		// Make "Output" directory
		for checkFolder != nil {
			err := os.Mkdir("./Output", 0755)
			if err != nil {
				fmt.Println("Error creating directory")
			}
			_, checkFolder = os.Stat("./Output")
		}
	}

	file, err := os.OpenFile("Output/"+code+".csv", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Fail to open file because", err.Error())
		return
	}
	defer file.Close()

	stock := models.Stock{}
	file.WriteString("Date,Code,Local IS,Local CP,Local PF,Local IB,Local ID,Local MF,Local SC,Local FD,Local OT,Foreign IS,Foreign CP,Foreign PF,Foreign IB,Foreign ID,Foreign MF,Foreign SC,Foreign FD,Foreign OT\n")
	for {
		err = rows.Scan(&stock.Date, &stock.Kode, &stock.LocalIS, &stock.LocalCP, &stock.LocalPF,
			&stock.LocalIB, &stock.LocalID, &stock.LocalMF, &stock.LocalSC, &stock.LocalFD, &stock.LocalOT,
			&stock.ForeignIS, &stock.ForeignCP, &stock.ForeignPF, &stock.ForeignIB, &stock.ForeignID,
			&stock.ForeignMF, &stock.ForeignSC, &stock.ForeignFD, &stock.ForeignOT)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		formattedDate := stock.Date.Format("02-01-2006")
		file.WriteString(formattedDate + ",")
		file.WriteString(stock.Kode + ",")
		file.WriteString(strconv.Itoa(int(stock.LocalIS)) + ",")
		file.WriteString(strconv.Itoa(int(stock.LocalCP)) + ",")
		file.WriteString(strconv.Itoa(int(stock.LocalPF)) + ",")
		file.WriteString(strconv.Itoa(int(stock.LocalIB)) + ",")
		file.WriteString(strconv.Itoa(int(stock.LocalID)) + ",")
		file.WriteString(strconv.Itoa(int(stock.LocalMF)) + ",")
		file.WriteString(strconv.Itoa(int(stock.LocalSC)) + ",")
		file.WriteString(strconv.Itoa(int(stock.LocalFD)) + ",")
		file.WriteString(strconv.Itoa(int(stock.LocalOT)) + ",")

		file.WriteString(strconv.Itoa(int(stock.ForeignIS)) + ",")
		file.WriteString(strconv.Itoa(int(stock.ForeignCP)) + ",")
		file.WriteString(strconv.Itoa(int(stock.ForeignPF)) + ",")
		file.WriteString(strconv.Itoa(int(stock.ForeignIB)) + ",")
		file.WriteString(strconv.Itoa(int(stock.ForeignID)) + ",")
		file.WriteString(strconv.Itoa(int(stock.ForeignMF)) + ",")
		file.WriteString(strconv.Itoa(int(stock.ForeignSC)) + ",")
		file.WriteString(strconv.Itoa(int(stock.ForeignFD)) + ",")
		file.WriteString(strconv.Itoa(int(stock.ForeignOT)) + "\n")

		// If there is not next data then break out of loop
		if !rows.Next() {
			break
		}
	}
	fmt.Printf("File %s.csv exported\n", code)
}

func InsertData(poolDB *sql.DB, fileName string) {
	ctx := context.Background()
	sql_query := "INSERT INTO Stocks VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	statement, err := poolDB.PrepareContext(ctx, sql_query)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer statement.Close()

	file, err := os.OpenFile(fileName, os.O_RDONLY, 0444)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var stock = models.Stock{}
	dateFormatter := "02-Jan-2006"

	// Remove header
	_, _, err = reader.ReadLine()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var rowsData []byte = nil
	for {
		rowsData, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		}

		stockData := strings.Split(string(rowsData), "|")

		// Data "Type" from KSEI are "EQUITY", "CORPORATE BOND", and etc
		// If the data type is equal then "CORPORATE BOND" then the "EQUITY" type is already read
		// "EQUITY" is the type of the stock
		if stockData[2] == "CORPORATE BOND" {
			break
		}

		// Skip Preferred stock and other who has more than 4 character
		if len(stockData[1]) != 4 {
			continue
		}

		// Change the string to date
		stock.Date, err = time.Parse(dateFormatter, string(stockData[0]))
		if err != nil {
			fmt.Println(err.Error(), "1")
			return
		}

		// Format the date
		stock.Date, err = time.Parse("02-01-2006", stock.Date.Format("02-01-2006"))
		if err != nil {
			fmt.Println(err.Error(), "2")
			return
		}

		stock.Kode = string(stockData[1])
		stock.LocalIS, _ = strconv.ParseUint(string(stockData[5]), 10, 64)
		stock.LocalCP, _ = strconv.ParseUint(string(stockData[6]), 10, 64)
		stock.LocalPF, _ = strconv.ParseUint(string(stockData[7]), 10, 64)
		stock.LocalIB, _ = strconv.ParseUint(string(stockData[8]), 10, 64)
		stock.LocalID, _ = strconv.ParseUint(string(stockData[9]), 10, 64)
		stock.LocalMF, _ = strconv.ParseUint(string(stockData[10]), 10, 64)
		stock.LocalSC, _ = strconv.ParseUint(string(stockData[11]), 10, 64)
		stock.LocalFD, _ = strconv.ParseUint(string(stockData[12]), 10, 64)
		stock.LocalOT, _ = strconv.ParseUint(string(stockData[13]), 10, 64)

		stock.ForeignIS, _ = strconv.ParseUint(string(stockData[15]), 10, 64)
		stock.ForeignCP, _ = strconv.ParseUint(string(stockData[16]), 10, 64)
		stock.ForeignPF, _ = strconv.ParseUint(string(stockData[17]), 10, 64)
		stock.ForeignIB, _ = strconv.ParseUint(string(stockData[18]), 10, 64)
		stock.ForeignID, _ = strconv.ParseUint(string(stockData[19]), 10, 64)
		stock.ForeignMF, _ = strconv.ParseUint(string(stockData[20]), 10, 64)
		stock.ForeignSC, _ = strconv.ParseUint(string(stockData[21]), 10, 64)
		stock.ForeignFD, _ = strconv.ParseUint(string(stockData[22]), 10, 64)
		stock.ForeignOT, _ = strconv.ParseUint(string(stockData[23]), 10, 64)

		fmt.Println(stock.Kode)
		_, err = statement.ExecContext(ctx, stock.Date, stock.Kode, stock.LocalIS, stock.LocalCP, stock.LocalPF, stock.LocalIB, stock.LocalID, stock.LocalMF, stock.LocalSC, stock.LocalFD, stock.LocalOT,
			stock.ForeignIS, stock.ForeignCP, stock.ForeignPF, stock.ForeignIB, stock.ForeignID, stock.ForeignMF, stock.ForeignSC, stock.ForeignFD, stock.ForeignOT)

		if err != nil {
			fmt.Println(err.Error(), "3")
			return
		}
	}
	fmt.Println("Success Insert Data")
}
