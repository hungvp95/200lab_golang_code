package model

type RestaurantFilter struct {
	CityId *int `json:"city_id,omitempty" form:"city_id,omitempty"`
}