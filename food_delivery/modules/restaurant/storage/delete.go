package storage

import (
	"context"

	"github.com/definev/200lab_golang/food_delivery/modules/restaurant/model"
)

func (s *sqlStore) SoftDeleteRestaurant(
	ctx context.Context,
	id int,
) error {
	query := s.
		db.
		Where("id = ?").
		Table(model.Restaurant{}.TableName()).
		Update("status", 0)

	if err := query.Error; err != nil {
		return err
	}

	return nil
}
