package restaurantstorage

import (
	"context"
	restaurantmodel "food-delivery-200lab/module/restaurant/model"
)

func (store *sqlStore) Delete(ctx context.Context, id int) error {
	statusDeleted := map[string]interface{}{"status": 0}

	if err := store.db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Updates(statusDeleted).Error; err != nil {
		return err
	}

	return nil
}
