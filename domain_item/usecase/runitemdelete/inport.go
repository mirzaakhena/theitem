package runitemdelete

import (
	"theitem/domain_item/model/vo"
	"theitem/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	ItemID vo.ItemID `form:"item_id"`
}

type InportResponse struct {
}
