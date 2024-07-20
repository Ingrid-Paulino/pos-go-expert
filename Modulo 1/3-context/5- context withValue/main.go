package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "senha")
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) { //context sempre vem como o primeiro valor do parametro
	token := ctx.Value("token")
	fmt.Println(token)
}
