package handlers

import (
	"StockMatchingEngine/model"
	"StockMatchingEngine/service"
	"encoding/json"
	"fmt"
	"log"

	"github.com/kataras/iris/v12"
)

type OrderRouter struct {
	// Dependencies that OrdeRrouter needs.
	OrderService *service.OrderService
	TradeService *service.ServiceTrade
}

func (r *OrderRouter) List(ctx iris.Context) {
	orders, err := r.OrderService.GetAll()
	if err != nil {
		ctx.StopWithText(iris.StatusInternalServerError, "List Failed")
		return
	}

	ctx.JSON(orders)
}

func (r *OrderRouter) Create(ctx iris.Context) {
	log.Println("=================================== - (1)")
	order := new(model.Order)

	if err := ctx.ReadJSON(order); err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}
	log.Println("=================================== - (2)")

	if err := r.OrderService.Create(order, r.TradeService); err != nil {
		log.Println("=================================== - (3)")
		ctx.Application().Logger().Error(err)
		// fire generic 500 error, client should not be aware the database internals.
		ctx.StopWithText(iris.StatusInternalServerError, "Create Failed")
		return
	}

	ctx.StatusCode(iris.StatusCreated)
}

func (r *OrderRouter) ListTickerValues(ctx iris.Context) {
	ticker := ctx.Params().Get("ticker")

	lowestBuy, highestSell, err := r.OrderService.GetLowerBuyHigherSell(ticker)

	if err != nil {
		ctx.Application().Logger().Error(err)
		ctx.StopWithText(iris.StatusInternalServerError, "List Failed")
		return
	}
	lowestBuyJSON, err := json.Marshal(lowestBuy)
	if err != nil {
		fmt.Println(err)
		return
	}
	highestSeJSONl, err := json.Marshal(highestSell)
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx.JSON(iris.Map{
		"lowest_buy":   lowestBuyJSON,
		"highest_sell": highestSeJSONl,
	})
}
