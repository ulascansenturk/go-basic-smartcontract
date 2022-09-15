package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ulascansenturk/go-basic-smartcontract/api"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		panic(err)
	}
	contractAddress := fmt.Sprint("0x1b9C451c4005A3e0bcd78F22ECFC5Ac6aF61a803")
	conn, err := api.NewApi(common.HexToAddress(contractAddress), client)
	if err != nil {
		panic(err)
	}
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/sort", func(c echo.Context) error {

		arr := make([]*big.Int, 0)
		for i := 0; i < 100; i++ {
			arr = append(arr, big.NewInt(int64(rand.Int())))
		}

		reply, err := conn.SortArray(&bind.CallOpts{}, arr)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, reply)
	})

	// Start server
	e.Logger.Fatal(e.Start(":9997"))
}
