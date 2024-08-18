package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type ExchangeData struct { //Exchange == Cambio
	Data Exchange `json:"USD"`
}

type Exchange struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

//go run server.go
//Banco de dados:
//sqlite3 data.db
//.tables
//create table exchanges (id string, code string, codein string, name string, high string, low string, varBid string, pctChange string, bid string, ask string, timestamp string, create_date string);
//.tables
//go run client.go
//select * from exchanges;

func main() {
	http.HandleFunc("/cotacao", GetExchangeHandler)
	http.ListenAndServe(":8080", nil)
}

func GetExchangeHandler(w http.ResponseWriter, r *http.Request) {
	//func GetExchangeHandler() {
	ctxCli := r.Context()
	ctx := context.Background()
	select { //switch case assincrono
	case <-ctx.Done():
		log.Println("Request of api canceled")
		w.Write([]byte("Request of api canceled\n"))
		return
	case <-ctxCli.Done():
		log.Println("Request of client canceled")
		w.Write([]byte("Request of client canceled\n"))
		return
	case <-ctx.Done():
		log.Println("Request canceled by client")
	}

	db, err := connectionDataBase()
	log.Println(err)
	defer db.Close()

	if err != nil {
		w.Write([]byte("Error connecting to database: " + err.Error() + "\n"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	exchange, err := getExchange(ctx)
	if err != nil {
		w.Write([]byte("Error getting exchange\n"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = saveExchangeInDB(ctx, db, exchange)
	if err != nil {
		w.Write([]byte("Error saving exchange in database: " + err.Error() + "\n"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(exchange.Data.Bid)
}

func connectionDataBase() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", "./data.db")

	if err != nil {
		return db, errors.New(fmt.Sprintf("fail to connect to database: %s", err))
	}
	fmt.Println("Successfully connected to database")
	return db, nil
}

func getExchange(ctx context.Context) (*ExchangeData, error) {
	ctx, cancel := context.WithTimeout(ctx, 200*time.Second)
	defer cancel()

	url := fmt.Sprintf(`https://economia.awesomeapi.com.br/json/all/USD-BRL`)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("fail to create request: %s", err))
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("fail to get response: %s", err))
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("fail to read body: %s", err))
	}

	var exchange ExchangeData
	err = json.Unmarshal(body, &exchange)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("fail to unmarshal body: %s", err))
	}
	return &exchange, nil
}

func saveExchangeInDB(ctx context.Context, db *sql.DB, exchange *ExchangeData) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	id := uuid.New().String()
	_, err := db.ExecContext(ctx, "INSERT INTO exchanges (id, code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, create_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		id, exchange.Data.Code, exchange.Data.Codein, exchange.Data.Name, exchange.Data.High, exchange.Data.Low, exchange.Data.VarBid, exchange.Data.PctChange, exchange.Data.Bid, exchange.Data.Ask, exchange.Data.Timestamp, exchange.Data.CreateDate)
	if err != nil {
		log.Println("Error: ", err)
		return errors.New(fmt.Sprintf("fail to save in the database: %s", err))
	}
	return nil
}
