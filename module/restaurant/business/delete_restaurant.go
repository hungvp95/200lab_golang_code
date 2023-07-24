package reataurantbusiness

import (
	"context"
	"food-delivery-200lab/common"
	restaurantmodel "food-delivery-200lab/module/restaurant/model"
)

type deleteRestaurantBusiness struct {
	store DeleteRestaurantStore
}

type DeleteRestaurantStore interface {
	FindDataWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (*restaurantmodel.Restaurant, error)
	Delete(ctx context.Context, id int) error
}

func InitDeleteRestaurantBusiness(store DeleteRestaurantStore) *deleteRestaurantBusiness {
	return &deleteRestaurantBusiness{store: store}
}

func (business *deleteRestaurantBusiness) DeleteRestaurant(ctx context.Context, id int) error {
	oldData, err := business.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrEntityNotFound(restaurantmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}

	if err := business.store.Delete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(restaurantmodel.EntityName, err)
	}

	return nil
}