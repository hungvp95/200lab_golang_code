package tgin

import (
	"strconv"

	"github.com/definev/200lab_golang/food_delivery/common"
	"github.com/definev/200lab_golang/food_delivery/component"
	"github.com/definev/200lab_golang/food_delivery/modules/restaurant/handler"
	"github.com/definev/200lab_golang/food_delivery/modules/restaurant/storage"
	"github.com/gin-gonic/gin"
)

func DeleteRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, common.ErrInvalidRequest(common.ErrInvalidRequest(err)))
			return
		}

		store := storage.NewSQLStore(appCtx.GetDBConnection())
		handler := handler.NewDeleteRestaurantHandler(store)
		err = handler.DeleteRestaurant(ctx.Request.Context(), id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.String(200, "Success!")
	}
}
