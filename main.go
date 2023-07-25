package main

import (
	"food-delivery-200lab/component"
	"food-delivery-200lab/middleware"
	ginrestaurant "food-delivery-200lab/module/restaurant/transport/gin"
	"food-delivery-200lab/module/upload/transport/ginupload"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	/*Connect to MySQL db*/
	connStr := os.Getenv("MYSQL_CONN_STR")
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(db, err)

	db = db.Debug()

	appCtx := component.NewAppContext(db)
	router := gin.Default()
	router.Use(middleware.Recover(appCtx))
	// only use for case: save file to static folder
	router.Static("/static", "./static")
	group := router.Group("/v1")

	testConnectServer(router)
	createNewRestaurant(appCtx, group)
	getRestaurants(appCtx, group)
	getRestaurantById(appCtx, group)
	updateRestaurantById(appCtx, group)
	deleteRestaurantById(appCtx, group)
	uploadImageFile(appCtx, group)

	routerErr := router.Run()
	if routerErr != nil {
		return
	}
}

func deleteRestaurantById(appCtx component.AppContext, group *gin.RouterGroup) {
	group.DELETE("restaurant/:id", ginrestaurant.DeleteRestaurant(appCtx))
}

func updateRestaurantById(appCtx component.AppContext, group *gin.RouterGroup) {
	group.PATCH("restaurant/:id", ginrestaurant.UpdateRestaurant(appCtx))
}

func getRestaurantById(appCtx component.AppContext, group *gin.RouterGroup) {
	group.GET("restaurant/:id", ginrestaurant.GetRestaurantById(appCtx))
}

func getRestaurants(appCtx component.AppContext, group *gin.RouterGroup) {
	group.GET("restaurants", ginrestaurant.GetListRestaurant(appCtx))
}

func createNewRestaurant(appCtx component.AppContext, group *gin.RouterGroup) {
	group.POST("restaurant", ginrestaurant.CreateRestaurant(appCtx))
}

func uploadImageFile(appCtx component.AppContext, group *gin.RouterGroup) {
	group.POST("uploadImage", ginupload.UploadImageSimple(appCtx))
}

// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
func testConnectServer(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
