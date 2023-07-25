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

func GetListRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging common.Paging
		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.FullFill()

		var filter restaurantmodel.Filter
		if err := ctx.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		store := restaurantstorage.NewSqlStore(db)

		business := reataurantbusiness.InitGetRestaurantBusiness(store)
		result, err := business.GetListRestaurant(ctx.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].MaskId(false)
		}

		ctx.JSON(http.StatusOK, common.FullSuccessResponse(result, paging, filter))
	}
}

func GetRestaurantById(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//id, err := strconv.Atoi(ctx.Param("id"))
		uid, err := common.FromBase58(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		db := appCtx.GetMainDBConnection()
		store := restaurantstorage.NewSqlStore(db)

		business := reataurantbusiness.InitGetRestaurantBusiness(store)
		result, getErr := business.GetRestaurantById(ctx.Request.Context(), int(uid.GetLocalID()))
		if getErr != nil {
			panic(getErr)
		}

		result.MaskId(false)
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
