package runitemdelete

import (
	"context"
	"theitem/domain_item/model/service"
)

type runItemDeleteInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runItemDeleteInteractor{
		outport: outputPort,
	}
}

func (r *runItemDeleteInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	err := service.DeleteItem(ctx, req.ItemID, r.outport)
	if err != nil {
		return nil, err
	}

	return res, nil
}
