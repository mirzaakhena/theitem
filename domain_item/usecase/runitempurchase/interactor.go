package runitempurchase

import (
	"context"
	"theitem/domain_item/model/service"
)

type runItemPurchaseInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runItemPurchaseInteractor{
		outport: outputPort,
	}
}

func (r *runItemPurchaseInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	itemObj, err := service.FindOneItem(ctx, req.ItemID, r.outport)
	if err != nil {
		return nil, err
	}

	err = itemObj.Purchase(req.Quantity)
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveItem(ctx, itemObj)
	if err != nil {
		return nil, err
	}

	return res, nil
}
