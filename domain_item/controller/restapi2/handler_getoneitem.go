package restapi2

import (
	"context"
	"net/http"
	"theitem/domain_item/model/entity"
	"theitem/domain_item/model/vo"
	"theitem/domain_item/usecase/getoneitem"
	"theitem/shared/gogen"
	"theitem/shared/infrastructure/logger"
	"theitem/shared/model/payload"
	"theitem/shared/util"

	"github.com/labstack/echo/v4"
)

func (r *controller) getOneItemHandler() echo.HandlerFunc {

	type InportRequest = getoneitem.InportRequest
	type InportResponse = getoneitem.InportResponse

	inport := gogen.GetInport[InportRequest, InportResponse](r.GetUsecase(InportRequest{}))

	type request struct {
		ItemID vo.ItemID `form:"item_id,omitempty,default=0"`
	}

	type response struct {
		Item *entity.Item `json:"item"`
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
		req.ItemID = jsonReq.ItemID

		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			r.log.Error(ctx, err.Error())
			return c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
		}

		var jsonRes response
		jsonRes.Item = res.Item

		r.log.Info(ctx, util.MustJSON(jsonRes))
		return c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
