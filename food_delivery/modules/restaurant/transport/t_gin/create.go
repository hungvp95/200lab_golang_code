package tgin

import (
	"net/http"

	"github.com/definev/200lab_golang/food_delivery/common"
	"github.com/definev/200lab_golang/food_delivery/component"
	"github.com/definev/200lab_golang/food_delivery/modules/restaurant/handler"
	"github.com/definev/200lab_golang/food_delivery/modules/restaurant/model"
	"github.com/definev/200lab_golang/food_delivery/modules/restaurant/storage"
	"github.com/gin-gonic/gin"
)

func CreateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data model.RestaurantCreate

		if err := ctx.Bind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSQLStore(appCtx.GetDBConnection())
		handler := handler.NewCreateRestaurantHandler(store)

		if err := handler.CreateRestaurant(ctx, &data); err != nil {
			ctx.JSON(http.StatusBadGateway, err)
			return
		}

		ctx.JSON(http.StatusOK, common.NewDataSuccessRes(data))
	}
}
