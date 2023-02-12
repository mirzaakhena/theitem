package restapi

import (
	"context"
	"net/http"
	"theitem/domain_item/model/vo"
	"theitem/domain_item/usecase/runitemupdate"
	"theitem/shared/gogen"
	"theitem/shared/infrastructure/logger"
	"theitem/shared/model/payload"
	"theitem/shared/util"
	"time"

	"github.com/gin-gonic/gin"
)

func (r *controller) runItemUpdateHandler() gin.HandlerFunc {

	type InportRequest = runitemupdate.InportRequest
	type InportResponse = runitemupdate.InportResponse

	inport := gogen.GetInport[InportRequest, InportResponse](r.GetUsecase(InportRequest{}))

	type request struct {
		InportRequest
	}

	type response struct {
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		var jsonReq request
		err := c.BindJSON(&jsonReq)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req InportRequest
		req = jsonReq.InportRequest
		req.Now = time.Now()
		req.ItemID = vo.ItemID(c.Param("item_id"))

		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		_ = res

		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
