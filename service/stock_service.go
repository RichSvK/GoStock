package service

import (
	"bufio"
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
	defer file.Close()

	file.WriteString("Date,Code,Local IS,Local CP,Local PF,Local IB,Local ID,Local MF,Local SC,Local FD,Local OT,Foreign IS,Foreign CP,Foreign PF,Foreign IB,Foreign ID,Foreign MF,Foreign SC,Foreign FD,Foreign OT\n")
	for _, stock := range listStock {
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
	}
	fmt.Printf("File %s.csv exported\n", code)
}

func InsertData(fileName string) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0444)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	// Remove header
	_, _, err = reader.ReadLine()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var rowsData []byte = nil
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

		if err := repository.InsertData(stock); err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	fmt.Println("Success Insert Data")
}
