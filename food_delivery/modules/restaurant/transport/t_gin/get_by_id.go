package tgin

import (
	"net/http"

	"github.com/definev/200lab_golang/food_delivery/common"
	"github.com/definev/200lab_golang/food_delivery/component"
	"github.com/definev/200lab_golang/food_delivery/modules/restaurant/handler"
	"github.com/definev/200lab_golang/food_delivery/modules/restaurant/storage"
	"github.com/gin-gonic/gin"
)

func GetRestaurantById(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		store := storage.NewSQLStore(appCtx.GetDBConnection())
		hand := handler.NewGetRestaurantByIdHandler(store)
	
		result, err := hand.GetRestaurantById(ctx.Request.Context(), id)
		if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": err})
			return
		}

		ctx.JSON(http.StatusOK, common.NewDataSuccessRes(result))
	}
}