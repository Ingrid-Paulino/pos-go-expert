package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"gorm.io/gorm"
)

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
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Request failed:", err)
	}
	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading response body: %v\n", err)
	}
	fmt.Println("Response body bid:", string(res))

	//var exchange Exchange
	//err = json.Unmarshal(res, &exchange) //transforma em struct
	//
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
	//}

	//create a file
	file, err := os.Create("cotaçao.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
	}

	//write the response body to the file
	_, err = file.WriteString(fmt.Sprintf("Dólar: %s \n", res))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to file: %v\n", err)
	}
	fmt.Println("File created successfully!")
}
