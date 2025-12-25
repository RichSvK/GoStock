package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/RichSvK/Stock_Holder_Composition_Go/config"
	"github.com/RichSvK/Stock_Holder_Composition_Go/model"
)

func FindDataByCode(code string) ([]model.Stock, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	sql_query := "SELECT * FROM Stock WHERE Code = ? ORDER BY Date LIMIT 6"
	statement, err := config.PoolDB.PrepareContext(ctx, sql_query)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := statement.Close(); err != nil {
			fmt.Println("Failed to close statement")
		}
	}()

	rows, err := statement.QueryContext(ctx, code)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			fmt.Println("Failed to close rows")
		}
	}()

	var stock model.Stock
	var listStock []model.Stock = nil
	for rows.Next() {
		err = rows.Scan(&stock.Date, &stock.Code, &stock.LocalIS, &stock.LocalCP, &stock.LocalPF,
			&stock.LocalIB, &stock.LocalID, &stock.LocalMF, &stock.LocalSC, &stock.LocalFD, &stock.LocalOT,
			&stock.ForeignIS, &stock.ForeignCP, &stock.ForeignPF, &stock.ForeignIB, &stock.ForeignID,
			&stock.ForeignMF, &stock.ForeignSC, &stock.ForeignFD, &stock.ForeignOT, &stock.ListedShare)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		listStock = append(listStock, stock)
	}
	return listStock, nil
}

func InsertData(stock model.Stock) error {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	sql_query := "INSERT INTO Stock VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	statement, err := config.PoolDB.PrepareContext(ctx, sql_query)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	defer func() {
		if err := statement.Close(); err != nil {
			fmt.Println("Failed to close statement")
		}
	}()

	_, err = statement.ExecContext(ctx, stock.Date, stock.Code, stock.LocalIS, stock.LocalCP, stock.LocalPF, stock.LocalIB, stock.LocalID, stock.LocalMF, stock.LocalSC, stock.LocalFD, stock.LocalOT,
		stock.ForeignIS, stock.ForeignCP, stock.ForeignPF, stock.ForeignIB, stock.ForeignID, stock.ForeignMF, stock.ForeignSC, stock.ForeignFD, stock.ForeignOT, stock.ListedShare)
	return err
}
