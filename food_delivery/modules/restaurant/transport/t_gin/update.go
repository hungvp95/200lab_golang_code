package tgin

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/definev/200lab_golang/food_delivery/common"
	"github.com/definev/200lab_golang/food_delivery/component"
	"github.com/definev/200lab_golang/food_delivery/modules/restaurant/handler"
	"github.com/definev/200lab_golang/food_delivery/modules/restaurant/model"
	"github.com/definev/200lab_golang/food_delivery/modules/restaurant/storage"
	"github.com/gin-gonic/gin"
)

func UpdateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.New("id must be int").Error()})
			return
		}
		updateData := model.RestaurantUpdate{}
		if err := ctx.ShouldBind(&updateData); err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				gin.H{"error": errors.New("restaurant update data wrong format").Error()},
			)
			return
		}

		store := storage.NewSQLStore(appCtx.GetDBConnection())
		hand := handler.NewUpdateRestaurantHandler(store)

		if err := hand.UpdateRestaurant(
			ctx.Request.Context(),
			id,
			&updateData,
		); err != nil {
			ctx.JSON(
				http.StatusBadGateway,
				gin.H{"error": err.Error()},
			)
			return
		}

		ctx.JSON(http.StatusOK, common.NewDataSuccessRes(updateData))
	}
}
