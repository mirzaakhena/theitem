package runitempurchase

import (
	"theitem/domain_item/model/vo"
	"theitem/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	ItemID   vo.ItemID `uri:"item_id"`
	Quantity int       `json:"quantity"`
}

type InportResponse struct {
}
