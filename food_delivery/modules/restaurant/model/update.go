package model

type RestaurantUpdate struct {
	Name *string `json:"name"`
	Addr *string `json:"addr"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}
