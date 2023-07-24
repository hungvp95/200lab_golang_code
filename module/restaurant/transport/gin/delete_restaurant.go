package ginrestaurant

import (
	"food-delivery-200lab/common"
	"food-delivery-200lab/component"
	reataurantbusiness "food-delivery-200lab/module/restaurant/business"
	restaurantstorage "food-delivery-200lab/module/restaurant/storage"

	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		store := restaurantstorage.NewSqlStore(db)

		business := reataurantbusiness.InitDeleteRestaurantBusiness(store)
		if err := business.DeleteRestaurant(ctx.Request.Context(), id); err != nil {
			panic(err)
		}

		deleteSuccess := map[string]interface{}{
			"isDeleted": true,
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(deleteSuccess))
	}
}