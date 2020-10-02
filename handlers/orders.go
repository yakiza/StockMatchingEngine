package handlers

import (
	"StockMatchingEngine/model"
	"StockMatchingEngine/service"
	"encoding/json"
	"fmt"

	"github.com/kataras/iris/v12"
)

type OrderRouter struct {
	// Dependencies that OrdeRrouter needs.
	OrderRepo model.Repository
}

//CreateAndMatchOrder calls the method that creates the order.
//additionally makes a call to matchingOrder to get all matching
//orders if any exists
func (r *OrderRouter) CreateAndMatchOrder(ctx iris.Context) {
	order := new(model.Order)
	fmt.Println("======================================================= 4")

	if err := ctx.ReadJSON(order); err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	if err := r.OrderRepo.CreateOrder(order); err != nil {
		ctx.Application().Logger().Error(err)
		// fire generic 500 error, client should not be aware the database internals.
		ctx.StopWithText(iris.StatusInternalServerError, "Create Failed")
		return
	}
	fmt.Println("======================================================= 4")

	orderMatchingService := service.OrderMatchingService{r.OrderRepo}
	orderMatchingService.MatchingOrderEngine(order)

	ctx.StatusCode(iris.StatusCreated)
}

//------------------------------------------------------------------
//==================================================================
//------------------------------------------------------------------

// func (r *OrderRouter) List(ctx iris.Context) {
// 	orders, err := r.OrderService.GetAll()
// 	if err != nil {
// 		ctx.StopWithText(iris.StatusInternalServerError, "List Failed")
// 		return
// 	}

// 	ctx.JSON(orders)
// }

func (r *OrderRouter) ListTickerValues(ctx iris.Context) {
	ticker := ctx.Params().Get("ticker")

	lowestBuy, err := r.OrderRepo.GetTickerLowerBuy(ticker)
	if err != nil {
		ctx.Application().Logger().Error(err)
		ctx.StopWithText(iris.StatusInternalServerError, "List Failed")
		return
	}

	highestSell, err := r.OrderRepo.GetTickerHigherSell(ticker)

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
