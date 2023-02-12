package getallitem

import (
	"context"
	"theitem/shared/util"
)

type getAllItemInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &getAllItemInteractor{
		outport: outputPort,
	}
}

func (r *getAllItemInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	itemObjs, count, err := r.outport.FindAllItem(ctx, req.Page, req.Size, req.Filter)
	if err != nil {
		return nil, err
	}

	res.Count = count
	res.Items = util.ToSliceAny(itemObjs)

	return res, nil
}
