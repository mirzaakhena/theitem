package restapi2

import (
	"context"
	"net/http"
	"theitem/domain_item/model/repository"
	"theitem/domain_item/usecase/getallitem"
	"theitem/shared/gogen"
	"theitem/shared/infrastructure/logger"
	"theitem/shared/model/payload"
	"theitem/shared/util"

	"github.com/labstack/echo/v4"
)

func (r *controller) getAllItemHandler() echo.HandlerFunc {

	type InportRequest = getallitem.InportRequest
	type InportResponse = getallitem.InportResponse

	inport := gogen.GetInport[InportRequest, InportResponse](r.GetUsecase(InportRequest{}))

	type request struct {
		Page   int                        `form:"page,omitempty,default=0"`
		Size   int                        `form:"size,omitempty,default=0"`
		Filter repository.ItemQueryFilter `form:"filter,omitempty,default=0"`
	}

	type response struct {
		Count int64 `json:"count"`
		Items []any `json:"items"`
	}

	return func(c echo.Context) error {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		var jsonReq request
		err := c.Bind(&jsonReq)
		if err != nil {
			r.log.Error(ctx, err.Error())
			return c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
		}

		var req InportRequest
		req.Page = jsonReq.Page
		req.Size = jsonReq.Size
		req.Filter = jsonReq.Filter

		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			r.log.Error(ctx, err.Error())
			return c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
		}

		var jsonRes response
		jsonRes.Count = res.Count
		jsonRes.Items = res.Items

		r.log.Info(ctx, util.MustJSON(jsonRes))
		return c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
