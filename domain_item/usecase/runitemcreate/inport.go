package runitemcreate

import (
	"theitem/domain_item/model/entity"
	"theitem/domain_item/model/vo"
	"theitem/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	entity.ItemCreateRequest
}

type InportResponse struct {
	ItemID vo.ItemID `json:"item_id"`
}
