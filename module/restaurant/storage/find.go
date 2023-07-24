package restaurantstorage

import (
	"context"
	"food-delivery-200lab/common"
	restaurantmodel "food-delivery-200lab/module/restaurant/model"
	"gorm.io/gorm"
)

func (store *sqlStore) FindDataWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKey ...string,
) (*restaurantmodel.Restaurant, error) {
	var inputData restaurantmodel.Restaurant
	if err := store.db.Where(condition).First(&inputData).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &inputData, nil
}
