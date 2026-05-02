package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"tspo/db"
	"tspo/dtos"
	"tspo/handlers"
	pb "tspo/proto/gen"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
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
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("Ошибка загрузки файла .env: %w", err)
	}
	return nil
}

func GetLinks(apikey string, symbol string) []string {
	return []string{
		fmt.Sprintf("https://www.alphavantage.co/query?function=OVERVIEW&symbol=%s&apikey=%s", symbol, apikey),
		fmt.Sprintf("https://www.alphavantage.co/query?function=INCOME_STATEMENT&symbol=%s&apikey=%s", symbol, apikey),
		fmt.Sprintf("https://www.alphavantage.co/query?function=BALANCE_SHEET&symbol=%s&apikey=%s", symbol, apikey),
		fmt.Sprintf("https://www.alphavantage.co/query?function=CASH_FLOW&symbol=%s&apikey=%s", symbol, apikey),
	}
}

func Process(links []string) {
	overview, err := getData(links[0], &dtos.Overview{})
	if err != nil {
		return
	}

	incomeStatement, err := getData(links[1], &dtos.IncomeStatement{})
	if err != nil {
		return
	}

	balanceSheet, err := getData(links[2], &dtos.BalanceSheet{})
	if err != nil {
		return
	}

	cashFlow, err := getData(links[3], &dtos.CashFlow{})
	if err != nil {
		return
	}

	_, err = db.Insert(overview.ToDatabaseableSlice())
	if err != nil {
		print("1")
	}
	_, err = db.Insert(incomeStatement.ToDatabaseableSlice())
	if err != nil {
		print("2")
	}
	_, err = db.Insert(balanceSheet.ToDatabaseableSlice())
	if err != nil {
		print("3")
	}
	_, err = db.Insert(cashFlow.ToDatabaseableSlice())
	if err != nil {
		print("4")
	}
}

// func main() {
// 	err := LoadEnv()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = db.ConnectDB("DATABASE_URL")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.CloseDB()

// 	// apikey := "YR4YSSIHBCLS3OJM"
// 	// Process(GetLinks(apikey, "AAPL"))
// 	// Process(GetLinks(apikey, "AMZN"))
// 	// Process(GetLinks(apikey, "ADBE"))

// 	// Создаем новый gRPC сервер
// 	grpcServer := grpc.NewServer()

// 	// Регистрация сервиса
// 	pb.RegisterDataServiceServer(grpcServer, &handlers.Server{})

// 	// Создаем сетевой слушатель
// 	var port string = os.Getenv("AUTH_PORT")
// 	if port == "" {
// 		log.Fatal("PORT environment variable is not set")
// 	}

// 	lis, err := net.Listen("tcp", port)
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}

// 	log.Println("Auth gRPC server is running at " + port)
// 	if err := grpcServer.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }

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

	var wg *sync.WaitGroup = &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		var port string = os.Getenv("GRPC_PORT")
		if port == "" {
			log.Fatal("GRPC_PORT environment variable is not set")
		}

		grpcSrv := grpc.NewServer()
		pb.RegisterDataServiceServer(grpcSrv, &handlers.Server{})

		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		defer lis.Close()

		log.Println("gRPC server is running at: " + port)
		if err := grpcSrv.Serve(lis); err != nil {
			log.Fatal("gRPC server error:", err)
		}

		wg.Done()
	}()

	wg.Add(1)
	go func() {
		var port string = os.Getenv("REST_PORT")
		if port == "" {
			log.Fatal("REST_PORT environment variable is not set")
		}

		hs := &handlers.HttpServer{GrpcSrv: &handlers.Server{}}
		http.HandleFunc("/overview", handlers.EnableCORS(hs.OverviewHandler))
		http.HandleFunc("/income", handlers.EnableCORS(hs.IncomeHandler))
		http.HandleFunc("/balance", handlers.EnableCORS(hs.BalanceHandler))
		http.HandleFunc("/cashflow", handlers.EnableCORS(hs.CashFlowHandler))

		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		defer lis.Close()

		log.Println("REST server is running at: " + port)
		if err := http.Serve(lis, nil); err != nil {
			log.Fatal("REST server error:", err)
		}

		wg.Done()
	}()

	wg.Wait()
}
