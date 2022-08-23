package handler

import (
	"context"

	"github.com/definev/200lab_golang/food_delivery/modules/restaurant/model"
)

type CreateRestaurantStoreInterface interface {
	Create(ctx context.Context, data *model.RestaurantCreate) error
}

type createRestaurantHandler struct {
	store CreateRestaurantStoreInterface
}

func NewCreateRestaurantHandler(store CreateRestaurantStoreInterface) createRestaurantHandler {
	return createRestaurantHandler{store}
}

func (h createRestaurantHandler) CreateRestaurant(ctx context.Context, data *model.RestaurantCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	err := h.store.Create(ctx, data)

	return err
}
