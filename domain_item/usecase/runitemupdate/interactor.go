package runitemupdate

import (
	"context"
	"theitem/domain_item/model/service"
)

type runItemUpdateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runItemUpdateInteractor{
		outport: outputPort,
	}
}

func (r *runItemUpdateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	err := service.MakeSureItemNameIsAllowed(ctx, req.Name, r.outport)
	if err != nil {
		return nil, err
	}

	itemObj, err := service.FindOneItem(ctx, req.ItemID, r.outport)
	if err != nil {
		return nil, err
	}

	err = service.MakeSureItemNameIsUnique(ctx, req.Name, itemObj.ID, r.outport)
	if err != nil {
		return nil, err
	}

	err = itemObj.Update(req.ItemUpdateRequest)
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveItem(ctx, itemObj)
	if err != nil {
		return nil, err
	}

	return res, nil
}
