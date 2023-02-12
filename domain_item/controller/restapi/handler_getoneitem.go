package restapi

import (
	"context"
	"net/http"
	"theitem/domain_item/usecase/getoneitem"
	"theitem/shared/gogen"
	"theitem/shared/infrastructure/logger"
	"theitem/shared/model/payload"
	"theitem/shared/util"

	"github.com/gin-gonic/gin"
)

func (r *controller) getOneItemHandler() gin.HandlerFunc {

	type InportRequest = getoneitem.InportRequest
	type InportResponse = getoneitem.InportResponse

	inport := gogen.GetInport[InportRequest, InportResponse](r.GetUsecase(InportRequest{}))

	type request struct {
		InportRequest
	}

	type response struct {
		*InportResponse
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		var jsonReq request
		err := c.BindUri(&jsonReq)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req InportRequest
		req = jsonReq.InportRequest

		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		jsonRes.InportResponse = res

		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
