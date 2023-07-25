package restaurantstorage

import (
	"context"
	"food-delivery-200lab/common"
	restaurantmodel "food-delivery-200lab/module/restaurant/model"
)

func (store *sqlStore) GetListWithCondition(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	db := store.db.
		Table(restaurantmodel.RestaurantCreate{}.TableName()).
		Where("status in (1)")

	if filter != nil {
		if filter.OwnerId > 0 {
			db = db.Where("owner_id = ?", filter.OwnerId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	// paging use cursor
	if cursorId := paging.FakeCursor; cursorId != "" {
		uid, err := common.FromBase58(cursorId)
		if err != nil {
			return nil, common.ErrDB(err)
		}
		db = db.Where("id < ?", uid.GetLocalID())
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db.Offset(offset)
	}

	var result []restaurantmodel.Restaurant

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	// return next cursor for client
	if len(result) > 0 {
		lastItem := result[len(result)-1]
		lastItem.MaskId(false)
		paging.NextCursor = lastItem.FakeId.String()
	}

	return result, nil
}
