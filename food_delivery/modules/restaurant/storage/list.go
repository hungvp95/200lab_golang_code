package storage

import (
	"context"

	"github.com/definev/200lab_golang/food_delivery/modules/restaurant/model"
)

func (s *sqlStore) ListRestaurantByFilter(
	ctx context.Context,
	conditions map[string]any,
	filter *model.RestaurantFilter,
	moreKeys ...string,
) []model.Restaurant {
	var result []model.Restaurant

	db := s.db

	db = db.Where(conditions)

	for _, key := range moreKeys {
		db = db.Where(key)
	}

	if filter != nil {
		if filter.CityId != nil {
			db = db.Where("city_id = ?", filter.CityId)
		}
	}

	db.Find(&result)

	return result
}
