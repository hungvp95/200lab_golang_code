package reataurantbusiness

import (
	"context"
	restaurantmodel "food-delivery-200lab/module/restaurant/model"
)

type UpdateRestaurantStore interface {
	Update(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error
}

type updateRestaurantBusiness struct {
	store UpdateRestaurantStore
}

func InitUpdateRestaurantBusiness(store UpdateRestaurantStore) *updateRestaurantBusiness {
	return &updateRestaurantBusiness{store: store}
}

func (business *updateRestaurantBusiness) UpdateRestaurant(
	ctx context.Context,
	id int,
	data *restaurantmodel.RestaurantUpdate,
) error {
	if err := business.store.Update(ctx, id, data); err != nil {
		return err
	}
	return nil
}
