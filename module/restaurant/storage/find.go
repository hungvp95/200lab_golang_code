package restaurantstorage

import (
	"context"
	restaurantmodel "food-delivery-200lab/module/restaurant/model"
)

func (store *sqlStore) FindDataWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKey ...string,
) (*restaurantmodel.Restaurant, error) {
	var inputData restaurantmodel.Restaurant
	if err := store.db.Where(condition).First(&inputData).Error; err != nil {
		return nil, err
	}

	return &inputData, nil
}
