package restapi2

import (
	"context"
	"net/http"
	"theitem/domain_item/model/vo"
	"theitem/domain_item/usecase/runitemdelete"
	"theitem/shared/gogen"
	"theitem/shared/infrastructure/logger"
	"theitem/shared/model/payload"
	"theitem/shared/util"

	"github.com/labstack/echo/v4"
)

func (r *controller) runItemDeleteHandler() echo.HandlerFunc {

	type InportRequest = runitemdelete.InportRequest
	type InportResponse = runitemdelete.InportResponse

	inport := gogen.GetInport[InportRequest, InportResponse](r.GetUsecase(InportRequest{}))

	type request struct {
		ItemID vo.ItemID `json:"item_id"`
	}

	type response struct {
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
		_ = res

		r.log.Info(ctx, util.MustJSON(jsonRes))
		return c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
