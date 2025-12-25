package service

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/RichSvK/Stock_Holder_Composition_Go/model"
	"github.com/RichSvK/Stock_Holder_Composition_Go/repository"
)

func Export(code string) {
	listStock, err := repository.FindDataByCode(code)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(listStock) == 0 {
		fmt.Println("No stock with code:", code)
		return
	}

	file, err := os.OpenFile("output/"+code+".csv", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Fail to open file because", err.Error())
		return
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Failed to close file")
		}
	}()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Date", "Code", "Local IS", "Local CP", "Local PF", "Local IB", "Local ID", "Local MF", "Local SC", "Local FD", "Local OT", "Foreign IS", "Foreign CP", "Foreign PF", "Foreign IB", "Foreign ID", "Foreign MF", "Foreign SC", "Foreign FD", "Foreign OT"}
	// Write header
	if err := writer.Write(header); err != nil {
		return
	}

	for _, stock := range listStock {
		record := []string{
			stock.Date.Format("02-01-2006"),
			stock.Code,
			strconv.FormatUint(stock.LocalIS, 10),
			strconv.FormatUint(stock.LocalCP, 10),
			strconv.FormatUint(stock.LocalPF, 10),
			strconv.FormatUint(stock.LocalIB, 10),
			strconv.FormatUint(stock.LocalID, 10),
			strconv.FormatUint(stock.LocalMF, 10),
			strconv.FormatUint(stock.LocalSC, 10),
			strconv.FormatUint(stock.LocalFD, 10),
			strconv.FormatUint(stock.LocalOT, 10),

			strconv.FormatUint(stock.ForeignIS, 10),
			strconv.FormatUint(stock.ForeignCP, 10),
			strconv.FormatUint(stock.ForeignPF, 10),
			strconv.FormatUint(stock.ForeignIB, 10),
			strconv.FormatUint(stock.ForeignID, 10),
			strconv.FormatUint(stock.ForeignMF, 10),
			strconv.FormatUint(stock.ForeignSC, 10),
			strconv.FormatUint(stock.ForeignFD, 10),
			strconv.FormatUint(stock.ForeignOT, 10),
		}

		if err := writer.Write(record); err != nil {
			return
		}
	}

	if err := writer.Error(); err != nil {
		return
	}

	fmt.Printf("File %s.csv exported\n", code)
}

func InsertData(fileName string) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0444)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Failed to close file")
		}
	}()

	reader := bufio.NewReader(file)

	// Remove header
	_, _, err = reader.ReadLine()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var rowsData []byte
	var stock = model.Stock{}
	dateFormatter := "02-Jan-2006"

	for {
		rowsData, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		}

		stockData := strings.Split(string(rowsData), "|")

		// Data "Type" from KSEI are "EQUITY", "CORPORATE BOND", and etc
		// If the data type is equal then "CORPORATE BOND" then the "EQUITY" type is already read
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
			fmt.Println(err.Error())
			return
		}

		// Format the date
		stock.Date, err = time.Parse("02-01-2006", stock.Date.Format("02-01-2006"))
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		stock.Code = stockData[1]
		fields := []struct {
			ptr   *uint64
			index int
			name  string
		}{
			{&stock.LocalIS, 5, "Local IS"},
			{&stock.LocalCP, 6, "Local CP"},
			{&stock.LocalPF, 7, "Local PF"},
			{&stock.LocalIB, 8, "Local IB"},
			{&stock.LocalID, 9, "Local ID"},
			{&stock.LocalMF, 10, "Local MF"},
			{&stock.LocalSC, 11, "Local SC"},
			{&stock.LocalFD, 12, "Local FD"},
			{&stock.LocalOT, 13, "Local OT"},
			{&stock.ForeignIS, 15, "Foreign IS"},
			{&stock.ForeignCP, 16, "Foreign CP"},
			{&stock.ForeignPF, 17, "Foreign PF"},
			{&stock.ForeignIB, 18, "Foreign IB"},
			{&stock.ForeignID, 19, "Foreign ID"},
			{&stock.ForeignMF, 20, "Foreign MF"},
			{&stock.ForeignSC, 21, "Foreign SC"},
			{&stock.ForeignFD, 22, "Foreign FD"},
			{&stock.ForeignOT, 23, "Foreign OT"},
			{&stock.ListedShare, 24, "Listed Share"},
		}

		for _, f := range fields {
			val, err := strconv.ParseUint(string(stockData[f.index]), 10, 64)
			if err != nil {
				fmt.Printf("Failed to parse %s", f.name)
				return
			}
			*f.ptr = val
		}

		if err := repository.InsertData(stock); err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	fmt.Println("Success Insert Data")
}
