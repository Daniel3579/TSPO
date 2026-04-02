package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"tspo/db"
	"tspo/dtos"

	"github.com/joho/godotenv"
)

func middleware(url string, model any) error {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Ошибка при запросе: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Ошибка: статус код %d\n", resp.StatusCode)
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Ошибка при чтении ответа: %v\n", err)
		return err
	}

	err = json.Unmarshal(body, model)
	if err != nil {
		fmt.Printf("Ошибка при парсинге JSON: %v\n", err)
		return err
	}

	return nil
}

func getData[T any](url string, model *T) (*T, error) {
	err := middleware(url, model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func LoadEnv() error {
	err := godotenv.Load("../.env")
	if err != nil {
		return fmt.Errorf("Ошибка загрузки файла .env: %w", err)
	}
	return nil
}

func main() {
	err := LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	err = db.ConnectDB("DATABASE_URL")
	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseDB()

	overview, err := getData("https://www.alphavantage.co/query?function=OVERVIEW&symbol=IBM&apikey=demo", &dtos.Overview{})
	if err != nil {
		return
	}

	incomeStatement, err := getData("https://www.alphavantage.co/query?function=INCOME_STATEMENT&symbol=IBM&apikey=demo", &dtos.IncomeStatement{})
	if err != nil {
		return
	}

	balanceSheet, err := getData("https://www.alphavantage.co/query?function=BALANCE_SHEET&symbol=IBM&apikey=demo", &dtos.BalanceSheet{})
	if err != nil {
		return
	}

	cashFlow, err := getData("https://www.alphavantage.co/query?function=CASH_FLOW&symbol=IBM&apikey=demo", &dtos.CashFlow{})
	if err != nil {
		return
	}

	_, _ = db.Insert(overview.ToDatabaseableSlice())
	_, _ = db.Insert(incomeStatement.ToDatabaseableSlice())
	_, _ = db.Insert(balanceSheet.ToDatabaseableSlice())
	_, _ = db.Insert(cashFlow.ToDatabaseableSlice())

	fmt.Println(overview)
	fmt.Println(incomeStatement)
	fmt.Println(balanceSheet)
	fmt.Println(cashFlow)
}
