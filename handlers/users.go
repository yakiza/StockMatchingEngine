package handlers

import (
	"StockMatchingEngine/model"

	"github.com/kataras/iris/v12"
)

type UserRouter struct {
	// Dependencies that UserRouter needs.
	OrderRepo model.Repository
}

// Create is responsible to read the POST json data, and for to call the CreateUser
// method that executes the query that adds the user to the database
func (r *UserRouter) Create(ctx iris.Context) {

	user := new(model.User)
	if err := ctx.ReadJSON(user); err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	if err := r.OrderRepo.CreateUser(user); err != nil {
		ctx.Application().Logger().Error(err)
		ctx.StopWithText(iris.StatusInternalServerError, "Create failed")
		return
	}

	ctx.StatusCode(iris.StatusCreated)
}

// //	swagger:route GET /Users user userList
// //	Returns a list of all the registered users saved in the database
// //	responses:
// //		200: User
// func (r *UserRouter) List(ctx iris.Context) {
// 	userBasket, err := r.UserService.GetAll()
// 	if err != nil {
// 		ctx.Application().Logger().Error(err)
// 		ctx.StopWithText(iris.StatusInternalServerError, "Retrieving List failed")
// 		return
// 	}

// 	ctx.JSON(userBasket)
// }
