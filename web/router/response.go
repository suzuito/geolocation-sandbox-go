package router

import (
	"github.com/gin-gonic/gin"
	"github.com/suzuito/geolocation-sandbox-go/web"
)

func responseError(ctx *gin.Context, err *web.HTTPError) {
	ctx.IndentedJSON(err.Status, err)
}
