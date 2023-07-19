package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Restaurant struct {
	Id      int    `json:"id" gorm:"column:id;"`
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:addr;"`
	Status  int    `json:"status" gorm:"column:status;"`
}

func (Restaurant) TableName() string {
	return "restaurants_200lab"
}

type RestaurantUpdate struct {
	Name    *string `json:"name" gorm:"column:name;"`
	Address *string `json:"address" gorm:"column:addr;"`
	Status  *int    `json:"status" gorm:"column:status;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

func main() {
	/*Connect to MySQL db*/
	connStr := os.Getenv("MYSQL_CONN_STR")
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(db, err)

	router := gin.Default()
	group := router.Group("/v1")

	testConnectServer(router)
	createNewRestaurant(db, group)
	getRestaurants(db, group)
	getRestaurantById(db, group)
	updateRestaurantById(db, group)
	deleteRestaurantById(db, group)

	routerErr := router.Run()
	if routerErr != nil {
		return
	}
}

func deleteRestaurantById(db *gorm.DB, group *gin.RouterGroup) {
	group.DELETE("restaurant/:id", func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		db.Table(Restaurant{}.TableName()).
			Where("id = ?", id).
			Delete(nil)
		context.JSON(http.StatusOK, gin.H{
			"isDeleted": true,
		})
	})
}

func updateRestaurantById(db *gorm.DB, group *gin.RouterGroup) {
	group.PATCH("restaurant/:id", func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		var input RestaurantUpdate
		inputErr := context.ShouldBind(&input)
		if inputErr != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": inputErr.Error(),
			})
			return
		}

		db.Where("id = ?", id).Updates(&input)
		context.JSON(http.StatusOK, gin.H{
			"data": input,
		})
	})
}

func getRestaurantById(db *gorm.DB, group *gin.RouterGroup) {
	group.GET("restaurant/:id", func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		var data Restaurant
		db.Where("id = ?", id).First(&data)
		context.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})
}

func getRestaurants(db *gorm.DB, group *gin.RouterGroup) {
	group.GET("restaurants", func(context *gin.Context) {
		type Paging struct {
			Page  int `json:"page" form:"page"`
			Limit int `json:"limit" form:"limit"`
		}

		var paging Paging
		err := context.ShouldBind(&paging)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		if paging.Page <= 0 {
			paging.Page = 1
		}
		if paging.Limit <= 0 {
			paging.Limit = 5
		}
		offset := (paging.Page - 1) * paging.Limit

		var data []Restaurant
		db.Offset(offset).
			Order("id desc").
			Limit(paging.Limit).
			Find(&data)
		context.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})
}

func createNewRestaurant(db *gorm.DB, group *gin.RouterGroup) {
	group.POST("restaurant", func(c *gin.Context) {
		var input Restaurant
		inputErr := c.ShouldBind(&input)
		if inputErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": inputErr.Error(),
			})
			return
		}

		db.Create(&input)

		c.JSON(http.StatusOK, gin.H{
			"message": input,
		})
	})
}

// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
func testConnectServer(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
