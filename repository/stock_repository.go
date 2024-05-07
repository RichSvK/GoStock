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
	defer statement.Close()

	rows, err := statement.QueryContext(ctx, code)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stock model.Stock
	var listStock []model.Stock = nil
	for rows.Next() {
		err = rows.Scan(&stock.Date, &stock.Kode, &stock.LocalIS, &stock.LocalCP, &stock.LocalPF,
			&stock.LocalIB, &stock.LocalID, &stock.LocalMF, &stock.LocalSC, &stock.LocalFD, &stock.LocalOT,
			&stock.ForeignIS, &stock.ForeignCP, &stock.ForeignPF, &stock.ForeignIB, &stock.ForeignID,
			&stock.ForeignMF, &stock.ForeignSC, &stock.ForeignFD, &stock.ForeignOT)

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

	sql_query := "INSERT INTO Stock VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	statement, err := config.PoolDB.PrepareContext(ctx, sql_query)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer statement.Close()

	_, err = statement.ExecContext(ctx, stock.Date, stock.Kode, stock.LocalIS, stock.LocalCP, stock.LocalPF, stock.LocalIB, stock.LocalID, stock.LocalMF, stock.LocalSC, stock.LocalFD, stock.LocalOT,
		stock.ForeignIS, stock.ForeignCP, stock.ForeignPF, stock.ForeignIB, stock.ForeignID, stock.ForeignMF, stock.ForeignSC, stock.ForeignFD, stock.ForeignOT)
	return err
}
