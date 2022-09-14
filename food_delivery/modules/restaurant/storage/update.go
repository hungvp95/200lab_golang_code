package storage

import (
	"context"

	"github.com/definev/200lab_golang/food_delivery/common"
	"github.com/definev/200lab_golang/food_delivery/modules/restaurant/model"
)

func (s *sqlStore) UpdateRestaurant(
	ctx context.Context,
	id int,
	data *model.RestaurantUpdate,
	moreKeys ...string,
) error {
	db := s.db.Where("id = ?", id)

	if err := db.Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
