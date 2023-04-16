package main

import (
	"context"
	"fmt"
	"log"

	"github.com/adshao/go-binance/v2"
)

func main() {
	// Define API key and secret
	apiKey := "acv"
	apiSecret := "qwe"

	// Set to true if you want to use testnet
	binance.UseTestnet = true

	// Use API key and secret to connect to Binance API
	client := binance.NewClient(apiKey, apiSecret)

	// Get the server time
	serverTime, err := client.NewServerTimeService().Do(context.Background())
	if err != nil {
		panic(err)
	}

	// Print the server time
	fmt.Printf("Server time: %d\n", serverTime)

	// Retrieve the latest price of BTCUSDT
	price, err := client.NewListPricesService().
		Symbol("BTCUSDT").
		Do(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Latest price of BTCUSDT: %s", price[0].Price)

	// Place a limit buy order for 1 BTC at 40000 USDT
	order, err := client.NewCreateOrderService().
		Symbol("BTCUSDT").
		Side(binance.SideTypeBuy).
		Type(binance.OrderTypeLimit).
		TimeInForce(binance.TimeInForceTypeGTC).
		Quantity("1").
		Price("2").
		Do(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Order placed: %+v", order)
}
