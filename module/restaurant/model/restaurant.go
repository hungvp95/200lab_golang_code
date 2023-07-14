package restaurantmodel

type Restaurant struct {
	Id      int    `json:"id" gorm:"column:id;"`
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:addr;"`
	Status  int    `json:"status" gorm:"column:status;"`
}
