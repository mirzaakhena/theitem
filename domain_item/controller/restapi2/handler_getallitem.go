package restapi2

import (
	"context"
	"net/http"
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
		InportRequest
	}

	type response struct {
		*InportResponse
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
		req = jsonReq.InportRequest

		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			r.log.Error(ctx, err.Error())
			return c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
		}

		var jsonRes response
		jsonRes.InportResponse = res

		r.log.Info(ctx, util.MustJSON(jsonRes))
		return c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
