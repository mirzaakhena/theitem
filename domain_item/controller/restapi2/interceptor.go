package restapi2

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r *controller) authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

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

		return next(c)

	}
}

func (r *controller) authorization(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		authorized := true

		if !authorized {
			return echo.NewHTTPError(http.StatusForbidden)
		}

		return next(c)
	}
}
