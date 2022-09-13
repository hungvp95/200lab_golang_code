package handler

import (
	"context"
	"errors"

	"github.com/definev/200lab_golang/food_delivery/modules/restaurant/model"
)

type UpdateRestaurantInterface interface {
	FindRestaurantById(
		ctx context.Context,
		id int,
		moreKeys ...string,
	) (*model.Restaurant, error)
	UpdateRestaurant(
		ctx context.Context,
		id int,
		data *model.RestaurantUpdate,
		moreKeys ...string,
	) error
}

type updateRestaurantHandler struct {
	store UpdateRestaurantInterface
}

func NewUpdateRestaurantHandler(store UpdateRestaurantInterface) updateRestaurantHandler {
	return updateRestaurantHandler{store}
}

func (h updateRestaurantHandler) UpdateRestaurant(
	ctx context.Context,
	id int,
	data *model.RestaurantUpdate,
) error {
	oldData, err := h.store.FindRestaurantById(ctx, id)
	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("data deleted")
	}

	if err := h.store.UpdateRestaurant(ctx, id, data); err != nil {
		return err
	}

	return nil
}
