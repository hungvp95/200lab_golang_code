package storage

import (
	"context"

	"github.com/definev/200lab_golang/food_delivery/modules/restaurant/model"
)

func (s sqlStore) Create(ctx context.Context, data *model.RestaurantCreate) error {
	if err := s.db.Create(&data); err != nil {
		return err.Error
	}

	return nil
}
