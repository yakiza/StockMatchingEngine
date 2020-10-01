package handlers

import (
	"StockMatchingEngine/service"
	"StockMatchingEngine/storage"

	"github.com/kataras/iris/v12"
)

// Router is responsible to make calls to functions based on the url accessed
func Router(db *service.DatabaseService) func(iris.Party) {
	return func(r iris.Party) {
		userRouter := &UserRouter{
			OrderRepo: storage.NewPostgresOrderRepository(db),
		}

		orderRouter := &OrderRouter{
			OrderRepo: storage.NewPostgresOrderRepository(db),
		}

		// r.Get("/users", userRouter.List)
		r.Post("/users", userRouter.Create)

		// r.Get("/orders", orderRouter.List)
		r.Post("/orders", orderRouter.CreateAndMatchOrder)
		r.Post("/orders/{ticker}", orderRouter.ListTickerValues)
	}
}
