package restaurantmodel

type Filter struct {
	OwnerId int `json:"ownerId,omitempty" form:"ownerId"`
}
