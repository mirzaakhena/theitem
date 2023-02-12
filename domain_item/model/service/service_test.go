package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"theitem/domain_item/model/entity"
	"theitem/domain_item/model/vo"
	"theitem/shared/model/apperror"
)

type mock1 struct {
	t *testing.T
}

func (r *mock1) ExistInForbiddenNameList(ctx context.Context, name string) bool {
	if name == "a" {
		return true
	}
	return false
}

func (r *mock1) FindOneItem(ctx context.Context, itemID vo.ItemID) (*entity.Item, error) {

	if itemID.String() == "a" {
		return &entity.Item{ID: itemID}, nil
	}

	return nil, nil
}

func TestMakeSureItemNameIsAllowed001(t *testing.T) {
	assert.Nil(t, MakeSureItemNameIsAllowed(context.TODO(), "b", &mock1{}))
	assert.Equal(t, "ER0005", MakeSureItemNameIsAllowed(context.TODO(), "a", &mock1{}).(apperror.ErrorType).Code())
}

func TestFindOneItem(t *testing.T) {
	{
		itemObj, err := FindOneItem(context.TODO(), "a", &mock1{})
		assert.Nil(t, err)
		assert.Equal(t, "a", itemObj.ID.String())
	}

	{
		_, err := FindOneItem(context.TODO(), "b", &mock1{})
		assert.Equal(t, "ER0007", err.(apperror.ErrorType).Code())
	}

}
