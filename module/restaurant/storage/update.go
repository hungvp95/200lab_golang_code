package restaurantstorage

import (
	"context"
	"food-delivery-200lab/common"
	restaurantmodel "food-delivery-200lab/module/restaurant/model"
)

func (store *sqlStore) Update(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	if err := store.db.
		Where("id = ?", id).
		Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
