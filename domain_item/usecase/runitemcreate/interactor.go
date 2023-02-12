package runitemcreate

import (
	"context"
	"theitem/domain_item/model/entity"
	"theitem/domain_item/model/service"
)

type runItemCreateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runItemCreateInteractor{
		outport: outputPort,
	}
}

func (r *runItemCreateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	err := service.MakeSureItemNameIsAllowed(ctx, req.Name, r.outport)
	if err != nil {
		return nil, err
	}

	err = service.MakeSureItemNameIsUnique(ctx, req.Name, "", r.outport)
	if err != nil {
		return nil, err
	}

	itemObj, err := entity.NewItem(req.ItemCreateRequest)
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveItem(ctx, itemObj)
	if err != nil {
		return nil, err
	}

	res.ItemID = itemObj.ID

	return res, nil
}
