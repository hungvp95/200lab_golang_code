package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"food-delivery-200lab/common"
	"food-delivery-200lab/component"
	reataurantbusiness "food-delivery-200lab/module/restaurant/business"
	restaurantmodel "food-delivery-200lab/module/restaurant/model"
	restaurantstorage "food-delivery-200lab/module/restaurant/storage"
)

func CreateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input restaurantmodel.RestaurantCreate
		if inputErr := ctx.ShouldBind(&input); inputErr != nil {
			panic(common.ErrInvalidRequest(inputErr))
		}

		// panic and recover in goroutine
		/*go func() {
			defer common.AppRecover()
			var arr []int
			log.Println(arr[0])
		}()*/

		db := appCtx.GetMainDBConnection()
		store := restaurantstorage.NewSqlStore(db)

		business := reataurantbusiness.InitCreateRestaurantBusiness(store)
		if err := business.CreateRestaurant(ctx.Request.Context(), &input); err != nil {
			// Note: panic should only use in transport layer
			panic(err)
		}

		input.MaskId(false)
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(input.FakeId))
	}
}
