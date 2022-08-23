package tgin

import (
	"log"
	"net/http"

	"github.com/definev/200lab_golang/food_delivery/modules/restaurant/handler"
	"github.com/definev/200lab_golang/food_delivery/modules/restaurant/model"
	"github.com/definev/200lab_golang/food_delivery/modules/restaurant/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data model.RestaurantCreate

		if err := ctx.Bind(&data); err != nil {
			log.Println(err)
			return
		}

		store := storage.NewSQLStore(db)
		handler := handler.NewCreateRestaurantHandler(store)

		if err := handler.CreateRestaurant(ctx, &data); err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadGateway, nil)
		}

		ctx.JSON(http.StatusOK, gin.H{"ok": 1})
	}
}
