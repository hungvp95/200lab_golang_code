package reataurantbusiness

import (
	"context"
	"food-delivery-200lab/common"
	restaurantmodel "food-delivery-200lab/module/restaurant/model"
)

type createRestaurantBusiness struct {
	store CreateRestaurantStore
}

type CreateRestaurantStore interface {
	Create(ctx context.Context, input *restaurantmodel.RestaurantCreate) error
}

func InitCreateRestaurantBusiness(store CreateRestaurantStore) *createRestaurantBusiness {
	return &createRestaurantBusiness{store: store}
}

func (business *createRestaurantBusiness) CreateRestaurant(
	ctx context.Context,
	input *restaurantmodel.RestaurantCreate,
) error {
	if err := input.ValidateInputData(); err != nil {
		return common.ErrInvalidRequest(err)
	}
	if err := business.store.Create(ctx, input); err != nil {
		return common.ErrCannotCreateEntity(restaurantmodel.EntityName, err)
	}
	return nil
}
