package storage

import (
	"context"

	"github.com/definev/200lab_golang/food_delivery/modules/restaurant/model"
)

func (s *sqlStore) GetRestaurantById(
	ctx context.Context,
	id string,
	moreKeys ...string,
) (model.Restaurant, error) {
	result := model.Restaurant{}
	db := s.db.Table(model.Restaurant{}.TableName())

	for _, key := range moreKeys {
		db = db.Preload(key)
	}

	if err := db.
		Where("id = ?", id).
		First(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}
