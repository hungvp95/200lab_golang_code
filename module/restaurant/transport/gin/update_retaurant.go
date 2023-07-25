package ginrestaurant

import (
	"food-delivery-200lab/common"
	"food-delivery-200lab/component"
	reataurantbusiness "food-delivery-200lab/module/restaurant/business"
	restaurantmodel "food-delivery-200lab/module/restaurant/model"
	restaurantstorage "food-delivery-200lab/module/restaurant/storage"

	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		//id, err := strconv.Atoi(context.Param("id"))
		uid, err := common.FromBase58(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		var input restaurantmodel.RestaurantUpdate
		inputErr := context.ShouldBind(&input)
		if inputErr != nil {
			panic(common.ErrInvalidRequest(inputErr))
		}

		db := appCtx.GetMainDBConnection()
		store := restaurantstorage.NewSqlStore(db)

		business := reataurantbusiness.InitUpdateRestaurantBusiness(store)
		if updateErr := business.UpdateRestaurant(context, int(uid.GetLocalID()), &input); updateErr != nil {
			panic(updateErr)
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(input))
	}
}
