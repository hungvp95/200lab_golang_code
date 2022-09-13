package handler

import (
	"context"
	"errors"

	"github.com/definev/200lab_golang/food_delivery/modules/restaurant/model"
)

type DeleteRestaurantInterface interface {
	FindRestaurantById(
		ctx context.Context,
		id int,
		moreKeys ...string,
	) (*model.Restaurant, error)
	SoftDeleteRestaurant(ctx context.Context, id int) error
}

type deleteRestaurantHandler struct {
	store DeleteRestaurantInterface
}

func NewDeleteRestaurantHandler(store DeleteRestaurantInterface) *deleteRestaurantHandler {
	return &deleteRestaurantHandler{store}
}

func (h *deleteRestaurantHandler) DeleteRestaurant(ctx context.Context, id int) error {
	oldData, err := h.store.FindRestaurantById(ctx, id)
	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("Data already deleted!")
	}
	if err := h.store.SoftDeleteRestaurant(ctx, id); err != nil {
		return err
	}

	return nil
}
