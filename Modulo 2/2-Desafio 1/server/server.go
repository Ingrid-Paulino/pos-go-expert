package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	gorm.Model
}

func main() {
	//rodar a conecção do banco aqui?

	http.HandleFunc("/cotacao", GetExchangeHandler)
	http.ListenAndServe(":8080", nil)
}

func GetExchangeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	//select { //switch case assincrono
	////case <-ctx.Done():
	////	log.Println("Request of api canceled")
	////	w.Write([]byte("Request of api canceled\n"))
	////	return
	////case <-ctx2.Done():
	////	log.Println("Request of database canceled")
	////	w.Write([]byte("Request of database canceled\n"))
	////	return
	//case <-ctx.Done():
	//	log.Println("Request canceled by client")
	//}

	db, err := connectionDataBase()
	if err != nil {
		w.Write([]byte("Error connecting to database\n"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	exchange, err := getExchange(ctx)
	if err != nil {
		w.Write([]byte("Error getting exchange\n"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = saveExchangeInDB(db, exchange)
	if err != nil {
		w.Write([]byte("Error saving exchange in database\n"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(exchange.Data.Bid)
}

func connectionDataBase() (*gorm.DB, error) {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local" //parseTime=True e loc=Local são necessários para o gorm funcionar corretamente com o mysql e o go. O charset=utf8mb4 é necessário para suportar emojis. É necessário reiniciar o docker depois de adicionar esses parâmetros. Ajuda a evitar erros de timezone e de tipo de dado.
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, errors.New(fmt.Sprintf("fail to connect to database: %s", err))
	}
	fmt.Println("Successfully connected to database")
	return db, nil
}

func getExchange(ctx context.Context) (*ExchangeData, error) {
	//ctx, cancel := context.WithTimeout(ctx, 200*time.Second)
	//defer cancel()

	url := fmt.Sprintf(`https://economia.awesomeapi.com.br/json/all/USD-BRL`)
	req, err := http.NewRequest(http.MethodGet, url, nil)
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

func saveExchangeInDB(db *gorm.DB, exchange *ExchangeData) error {
	//ctx, cancel := context.WithTimeout(ctx, 500*time.Second)
	//db.WithContext(ctx)
	//defer cancel()

	db.AutoMigrate(&Exchange{}) //Create the bank if not exist
	res := db.Create(&exchange.Data)
	if res.Error != nil {
		log.Println("Error: ", res.Error)
		return errors.New(fmt.Sprintf("fail save in the database: %s", res.Error))
	}
	return nil
}
