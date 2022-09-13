package handler

import (
	"context"

	"github.com/definev/200lab_golang/food_delivery/modules/restaurant/model"
)

type GetRestaurantByIdInterface interface {
	FindRestaurantById(
		ctx context.Context,
		id int,
		moreKeys ...string,
	) (*model.Restaurant, error)
}

type getRestaurantByIdHandler struct {
	store GetRestaurantByIdInterface
}

func NewGetRestaurantByIdHandler(store GetRestaurantByIdInterface) getRestaurantByIdHandler {
	return getRestaurantByIdHandler{store}
}

func (h getRestaurantByIdHandler) GetRestaurantById(ctx context.Context, id int) (*model.Restaurant, error) {
	return h.store.FindRestaurantById(ctx, id)
}
