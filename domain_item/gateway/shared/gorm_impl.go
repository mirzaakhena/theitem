package shared

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"theitem/domain_item/model/entity"
	"theitem/domain_item/model/repository"
	"theitem/domain_item/model/vo"
	"theitem/shared/config"
	"theitem/shared/gogen"
	"theitem/shared/infrastructure/logger"
)

type GormImpl struct {
	*ForbiddenNameList
	AppData gogen.ApplicationData
	Cfg     *config.Config
	Log     logger.Logger
	Db      *gorm.DB
}

func NewGormImpl(appData gogen.ApplicationData, cfg *config.Config, log logger.Logger, db *gorm.DB) *GormImpl {

	return &GormImpl{
		ForbiddenNameList: NewForbiddenNameList(log),
		AppData:           appData,
		Cfg:               cfg,
		Log:               log,
		Db:                db,
	}
}

func (r *GormImpl) FindAllItem(ctx context.Context, page, size int, query repository.ItemQueryFilter) ([]*entity.Item, int64, error) {
	r.Log.Info(ctx, "called")

	sql := r.Db.
		Model(&entity.Item{})

	if query.Category != "" {
		sql = sql.Where("category = ?", query.Category)
	}

	if query.Rating != -1 {
		sql = sql.Where("rating = ?", query.Rating)
	}

	if query.ReputationBadge != "" {
		sql = sql.Where("reputation_badge = ?", query.ReputationBadge)
	}

	if query.AvailabilityMore != -1 {
		sql = sql.Where("availability >= ?", query.AvailabilityMore)
	}

	if query.AvailabilityLess != -1 {
		sql = sql.Where("availability <= ?", query.AvailabilityLess)
	}

	var objs []*entity.Item
	var count int64

	err := sql.
		Count(&count).
		Limit(size).
		Offset((page - 1) * size).
		Find(&objs).Error

	if err != nil {
		return nil, 0, err
	}

	return objs, count, nil
}

func (r *GormImpl) FindOneItem(ctx context.Context, itemID vo.ItemID) (*entity.Item, error) {
	r.Log.Info(ctx, "called")

	var obj entity.Item

	err := r.Db.First(&obj, itemID).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, nil
	}

	return &obj, nil
}

func (r *GormImpl) SaveItem(ctx context.Context, obj *entity.Item) error {
	r.Log.Info(ctx, "called")

	err := r.Db.Save(obj).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *GormImpl) HardDeleteOneItem(ctx context.Context, item *entity.Item) error {
	r.Log.Info(ctx, "called")

	err := r.Db.
		Unscoped(). // HARD Delete!
		Delete(item).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *GormImpl) FindOneItemByName(ctx context.Context, name string) (*entity.Item, error) {
	r.Log.Info(ctx, "called")

	var obj entity.Item

	err := r.Db.First(&obj, "name = ?", name).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, nil
	}

	return &obj, nil
}
