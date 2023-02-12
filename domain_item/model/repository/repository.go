package repository

import (
	"context"
	"theitem/domain_item/model/entity"
	"theitem/domain_item/model/vo"
)

type SaveItemRepo interface {
	SaveItem(ctx context.Context, obj *entity.Item) error
}

type ExistInForbiddenNameListRepo interface {
	ExistInForbiddenNameList(ctx context.Context, name string) bool
}

type ItemQueryFilter struct {
	Rating           vo.Rating   `form:"rating,omitempty,default=-1"`
	ReputationBadge  string      `form:"reputation_badge,omitempty"`
	AvailabilityMore int         `form:"availability_more,omitempty,default=-1"`
	AvailabilityLess int         `form:"availability_less,omitempty,default=-1"`
	Category         vo.Category `form:"category,omitempty"`
}

type FindAllItemRepo interface {
	FindAllItem(ctx context.Context, page, size int, query ItemQueryFilter) ([]*entity.Item, int64, error)
}

type FindOneItemRepo interface {
	FindOneItem(ctx context.Context, itemID vo.ItemID) (*entity.Item, error)
}

type FindOneItemByNameRepo interface {
	FindOneItemByName(ctx context.Context, name string) (*entity.Item, error)
}

type DeleteOneItemRepo interface {
	FindOneItemRepo
	HardDeleteOneItem(ctx context.Context, item *entity.Item) error
}
