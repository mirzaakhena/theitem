package runitemupdate

import (
	"theitem/domain_item/model/entity"
	"theitem/domain_item/model/vo"
	"theitem/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	ItemID                   vo.ItemID `uri:"item_id"`
	entity.ItemUpdateRequest `json:""`
}

type InportResponse struct {
}
