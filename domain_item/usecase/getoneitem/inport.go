package getoneitem

import (
	"theitem/domain_item/model/entity"
	"theitem/domain_item/model/vo"
	"theitem/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	ItemID vo.ItemID `uri:"item_id"`
}

type InportResponse struct {
	Item *entity.Item `json:"item"`
}
