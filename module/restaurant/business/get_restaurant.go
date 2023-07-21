package reataurantbusiness

import (
	"context"
	"food-delivery-200lab/common"
	restaurantmodel "food-delivery-200lab/module/restaurant/model"
)

type GetRestaurantStore interface {
	GetListWithCondition(
		context context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)

	FindDataWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (*restaurantmodel.Restaurant, error)
}

type getRestaurantBusiness struct {
	store GetRestaurantStore
}

func InitGetRestaurantBusiness(store GetRestaurantStore) *getRestaurantBusiness {
	return &getRestaurantBusiness{store: store}
}

func (business *getRestaurantBusiness) GetListRestaurant(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := business.store.GetListWithCondition(context, filter, paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (business *getRestaurantBusiness) GetRestaurantById(
	context context.Context,
	id int,
) (*restaurantmodel.Restaurant, error) {
	condition := map[string]interface{}{"id": id}
	result, err := business.store.FindDataWithCondition(context, condition)

	if err != nil {
		return nil, err
	}

	return result, nil
}
