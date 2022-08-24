package handler

import (
	"context"

	"github.com/definev/200lab_golang/food_delivery/modules/restaurant/model"
)

type GetRestaurantByIdInterface interface {
	GetRestaurantById(
		ctx context.Context,
		id string,
		moreKeys ...string,
	) (model.Restaurant, error)
}

type getRestaurantByIdHandler struct {
	store GetRestaurantByIdInterface
}

func NewGetRestaurantByIdHandler(store GetRestaurantByIdInterface) getRestaurantByIdHandler {
	return getRestaurantByIdHandler{store}
}

func (h getRestaurantByIdHandler) GetRestaurantById(ctx context.Context, id string) (model.Restaurant, error) {
	return h.store.GetRestaurantById(ctx, id)
}
