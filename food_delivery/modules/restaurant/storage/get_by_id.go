package storage

import (
	"context"

	"github.com/definev/200lab_golang/food_delivery/modules/restaurant/model"
)

func (s *sqlStore) FindRestaurantById(
	ctx context.Context,
	id int,
	moreKeys ...string,
) (*model.Restaurant, error) {
	result := model.Restaurant{}
	db := s.db.Table(model.Restaurant{}.TableName())

	for _, key := range moreKeys {
		db = db.Preload(key)
	}

	if err := db.
		Where("id = ?", id).
		First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}
