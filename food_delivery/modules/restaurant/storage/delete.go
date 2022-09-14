package storage

import (
	"context"

	"github.com/definev/200lab_golang/food_delivery/common"
	"github.com/definev/200lab_golang/food_delivery/modules/restaurant/model"
)

func (s *sqlStore) SoftDeleteRestaurant(
	ctx context.Context,
	id int,
) error {
	query := s.db.
		Table(model.Restaurant{}.TableName()).
		Where("id = ?", id).
		Update("status", 0)

	if err := query.Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
