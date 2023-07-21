package restaurantstorage

import (
	"context"
	restaurantmodel "food-delivery-200lab/module/restaurant/model"
)

func (store *sqlStore) Create(ctx context.Context, input *restaurantmodel.RestaurantCreate) error {
	err := store.db.Create(&input).Error
	if err != nil {
		return err
	}
	return nil
}
