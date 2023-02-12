package withmongodb

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"strings"
	"theitem/domain_item/gateway/shared"
	"theitem/domain_item/model/entity"
	"theitem/domain_item/model/repository"
	"theitem/domain_item/model/vo"
	"theitem/shared/config"
	"theitem/shared/gogen"
	"theitem/shared/infrastructure/logger"
)

type gateway struct {
	*shared.ForbiddenNameList
	appData gogen.ApplicationData
	cfg     *config.Config
	log     logger.Logger
	client  *mongo.Client

	collName string
}

// NewGateway ...
func NewGateway(log logger.Logger, appData gogen.ApplicationData, cfg *config.Config) *gateway {

	cdb := cfg.Database
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin", cdb.Username, cdb.Password, cdb.Host, cdb.Port, cdb.DBName)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		panic(err)
	}

	err = client.Database(cfg.Database.DBName).RunCommand(context.Background(), bson.D{{"create", "item"}}).Err()
	if err != nil {
		if !strings.Contains(err.Error(), "already exist") {
			panic(err)
		}
	}

	return &gateway{
		ForbiddenNameList: shared.NewForbiddenNameList(log),
		log:               log,
		appData:           appData,
		cfg:               cfg,
		client:            client,
		collName:          "item",
	}
}

func (r *gateway) FindAllItem(ctx context.Context, page, size int, query repository.ItemQueryFilter) ([]*entity.Item, int64, error) {
	r.log.Info(ctx, "called")

	coll := r.client.Database(r.cfg.Database.DBName).Collection(r.collName)
	filter := bson.M{}

	if query.Category != "" {
		filter["category"] = query.Category
	}

	if query.Rating != -1 {
		filter["rating"] = query.Rating
	}

	if query.ReputationBadge != "" {
		filter["reputation_badge"] = query.ReputationBadge
	}

	if query.AvailabilityMore != -1 {
		filter["availability"] = bson.M{"$gte": query.AvailabilityMore}
	}

	if query.AvailabilityLess != -1 {
		filter["availability"] = bson.M{"$lte": query.AvailabilityLess}
	}

	limit := int64(size)
	skip := int64((page - 1) * size)
	countOpts := options.CountOptions{Limit: &limit, Skip: &skip}
	findOpts := options.FindOptions{Limit: &limit, Skip: &skip}

	count, err := coll.CountDocuments(ctx, filter, &countOpts)
	if err != nil {
		return nil, 0, err
	}

	findCursor, err := coll.Find(ctx, filter, &findOpts)
	if err != nil {
		return nil, 0, err
	}

	objs := make([]*entity.Item, 0)
	err = findCursor.All(ctx, &objs)
	if err != nil {
		return nil, 0, err
	}

	return objs, count, nil
}

func (r *gateway) FindOneItem(ctx context.Context, itemID vo.ItemID) (*entity.Item, error) {
	r.log.Info(ctx, "called")

	coll := r.client.Database(r.cfg.Database.DBName).Collection(r.collName)

	var obj entity.Item
	err := coll.FindOne(ctx, bson.M{"_id": itemID}).Decode(&obj)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		}
		return nil, nil
	}

	return &obj, nil
}

func (r *gateway) SaveItem(ctx context.Context, obj *entity.Item) error {
	r.log.Info(ctx, "called")

	coll := r.client.Database(r.cfg.Database.DBName).Collection(r.collName)

	filter := bson.D{{"_id", obj.ID}}
	update := bson.D{{"$set", obj}}
	opts := options.Update().SetUpsert(true)

	_, err := coll.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}

func (r *gateway) FindOneItemByName(ctx context.Context, name string) (*entity.Item, error) {
	r.log.Info(ctx, "called")

	coll := r.client.Database(r.cfg.Database.DBName).Collection(r.collName)

	var obj entity.Item
	err := coll.FindOne(ctx, bson.M{"name": name}).Decode(&obj)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		}
		return nil, nil
	}

	return nil, nil
}

func (r *gateway) HardDeleteOneItem(ctx context.Context, item *entity.Item) error {
	r.log.Info(ctx, "called")

	coll := r.client.Database(r.cfg.Database.DBName).Collection(r.collName)

	filter := bson.D{{"_id", item.ID}}

	_, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
