package handlers

import (
	"StockMatchingEngine/model"
	"StockMatchingEngine/service"

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
	order := new(model.Order)

	if err := ctx.ReadJSON(order); err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	if err := r.OrderService.Create(order, r.TradeService); err != nil {
		ctx.Application().Logger().Error(err)
		// fire generic 500 error, client should not be aware the database internals.
		ctx.StopWithText(iris.StatusInternalServerError, "Create Failed")
		return
	}

	ctx.StatusCode(iris.StatusCreated)
}

func (r *OrderRouter) ListTickerValues(ctx iris.Context) {
	ticker := ctx.Params().Get("ticker")

	orders, err := r.OrderService.GetAllActiveOrders(ticker)
	if err != nil {
		ctx.Application().Logger().Error(err)
		ctx.StopWithText(iris.StatusInternalServerError, "List Failed")
		return
	}

	ctx.JSON(orders)
}
