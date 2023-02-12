package getallitem

import (
	"theitem/domain_item/model/repository"
	"theitem/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	Page   int                        `form:"page,omitempty,default=1"`
	Size   int                        `form:"size,omitempty,default=30"`
	Filter repository.ItemQueryFilter `form:""`
}

type InportResponse struct {
	Count int64 `json:"count"`
	Items []any `json:"items"`
}
