package getoneitem

import (
	"context"
	"theitem/domain_item/model/service"
)

type getOneItemInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &getOneItemInteractor{
		outport: outputPort,
	}
}

func (r *getOneItemInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	itemObj, err := service.FindOneItem(ctx, req.ItemID, r.outport)
	if err != nil {
		return nil, err
	}

	res.Item = itemObj

	return res, nil
}
