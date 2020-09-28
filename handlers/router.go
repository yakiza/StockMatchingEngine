package handlers

import (
	"StockMatchingEngine/service"

	"github.com/kataras/iris/v12"
)

// Router is responsible to make calls to functions based on the url accessed
func Router(db *service.DatabaseService) func(iris.Party) {
	return func(r iris.Party) {
		userRouter := &UserRouter{
			UserService: service.NewUserService(db),
		}

		orderRouter := &OrderRouter{
			OrderService: service.NewOrderService(db),
			TradeService: service.NewTradeService(db),
		}

		r.Get("/users", userRouter.List)
		r.Post("/users", userRouter.Create)

		r.Get("/orders", orderRouter.List)
		r.Post("/orders", orderRouter.Create)
		r.Post("/orders/{ticker}", orderRouter.ListTickerValues)
	}
}
