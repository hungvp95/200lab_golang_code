package fooddelivery

import (
	"log"
	"net/http"
	"os"

	"github.com/definev/200lab_golang/food_delivery/component"
	tgin "github.com/definev/200lab_golang/food_delivery/modules/restaurant/transport/t_gin"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Main() {
	connStr := os.Getenv("DBConnectionStr")
	db, err := gorm.Open(mysql.Open(connStr))
	if err != nil {
		log.Panicln(err)
	}

	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	appCtx := component.CreateAppComponent(db)

	gRestaurant := r.Group("/restaurant")
	{
		gRestaurant.GET("/ping", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "Pong!")
		})
		gRestaurant.POST("/", tgin.CreateRestaurant(appCtx))
	}

	r.Run("localhost:8080")
}
