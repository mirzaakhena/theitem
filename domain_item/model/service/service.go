package service

import (
	"context"
	"theitem/domain_item/model/entity"
	"theitem/domain_item/model/errorenum"
	"theitem/domain_item/model/repository"
	"theitem/domain_item/model/vo"
)

func MakeSureItemNameIsAllowed(ctx context.Context, name string, repo repository.ExistInForbiddenNameListRepo) error {

	if repo.ExistInForbiddenNameList(ctx, name) {
		return errorenum.ForbiddenWord.Var(name)
	}

	return nil
}

func FindOneItem(ctx context.Context, itemID vo.ItemID, repo repository.FindOneItemRepo) (*entity.Item, error) {

	item, err := repo.FindOneItem(ctx, itemID)
	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, errorenum.UnavailableItem.Var(itemID)
	}

	return item, nil
}

func DeleteItem(ctx context.Context, itemID vo.ItemID, repo repository.DeleteOneItemRepo) error {

	item, err := repo.FindOneItem(ctx, itemID)
	if err != nil {
		return err
	}

	if item == nil {
		return errorenum.UnavailableItem.Var(itemID)
	}

	err = repo.HardDeleteOneItem(ctx, item)
	if err != nil {
		return err
	}

	return nil
}

func MakeSureItemNameIsUnique(ctx context.Context, name string, itemID vo.ItemID, repo repository.FindOneItemByNameRepo) error {

	itemObj, err := repo.FindOneItemByName(ctx, name)
	if err != nil {
		return err
	}

	// if it is update case (itemID is exist)
	// it is not the current object that we want to update
	if itemID != "" {

		if itemObj != nil && itemObj.ID != itemID {
			return errorenum.ItemNameAlreadyExist.Var(itemObj.ID)
		}

	} else if itemObj != nil { // it is create new case (itemID is blank)

		return errorenum.ItemNameAlreadyExist.Var(name)
	}

	return nil
}
