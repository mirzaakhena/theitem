package restapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *controller) authentication() gin.HandlerFunc {
	return func(c *gin.Context) {

		// tokenInBytes, err := r.JwtToken.VerifyToken(c.GetHeader("token"))
		// if err != nil {
		// 	c.AbortWithStatus(http.StatusForbidden)
		// 	return
		// }
		//
		// var dataToken payload.DataToken
		// err = json.Unmarshal(tokenInBytes, &dataToken)
		// if err != nil {
		// 	c.AbortWithStatus(http.StatusForbidden)
		// 	return
		// }
		//
		// c.Set("data", dataToken)
		//
		// c.AbortWithStatus(http.StatusForbidden)
		// return

	}
}

func (r *controller) authorization() gin.HandlerFunc {

	return func(c *gin.Context) {

		authorized := true

		if !authorized {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
	}
}
