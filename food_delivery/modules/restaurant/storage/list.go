package storage

import (
	"context"

	"github.com/definev/200lab_golang/food_delivery/common"
	"github.com/definev/200lab_golang/food_delivery/modules/restaurant/model"
	"gorm.io/gorm"
)

func (s *sqlStore) ListRestaurantByFilter(
	ctx context.Context,
	conditions map[string]any,
	filter *model.RestaurantFilter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.Restaurant, error) {
	
	var result []model.Restaurant

	db := s.db.Table(model.Restaurant{}.TableName())

	for _, key := range moreKeys {
		db = db.Preload(key)
	}

	db = db.Where(conditions)
	filterDB(db, filter)

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

func filterDB(db *gorm.DB, filter *model.RestaurantFilter) *gorm.DB {
	if filter != nil {
		if filter.CityId != nil {
			db = db.Where("city_id = ?", filter.CityId)
		}
	}

	return db
}
