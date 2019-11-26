package router

import (
	"github.com/gin-gonic/gin"
	"github.com/suzuito/geolocation-sandbox-go/entity/model"
	"github.com/suzuito/geolocation-sandbox-go/web"
)

// GetUserLocations ...
func GetUserLocations(app web.Application) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		partition := ctx.Param("partition")
		cli, err := app.NewStoreCore(ctx)
		if err != nil {
			responseError(ctx, err)
			return
		}
		locations := []model.Location{}
		if err := cli.GetLocations(ctx, "dummyUser", partition, &locations); err != nil {
			responseError(ctx, web.FromStoreError(err))
			return
		}
		ctx.JSON(200, &locations)
	}
}
