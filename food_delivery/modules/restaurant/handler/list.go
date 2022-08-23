package handler

import (
	"context"

	"github.com/definev/200lab_golang/food_delivery/common"
	"github.com/definev/200lab_golang/food_delivery/modules/restaurant/model"
)

type ListRestaurantByFilterInterface interface {
	ListRestaurantByFilter(
		ctx context.Context,
		conditions map[string]any,
		filter *model.RestaurantFilter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]model.Restaurant, error)
}

type listRestaurantByFilterHandler struct {
	store ListRestaurantByFilterInterface
}

func NewListRestaurantHandler(store ListRestaurantByFilterInterface) listRestaurantByFilterHandler {
	return listRestaurantByFilterHandler{store}
}

func (h listRestaurantByFilterHandler) ListRestaurant(
	ctx context.Context,
	filter *model.RestaurantFilter,
	paging *common.Paging,
) ([]model.Restaurant, error) {
	res, err := h.store.ListRestaurantByFilter(
		ctx,
		map[string]any{},
		filter,
		paging,
	)

	if err != nil {
		return nil, err
	}

	return res, nil
}
