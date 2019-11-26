package router

import (
	"github.com/gin-gonic/gin"
	"github.com/suzuito/geolocation-sandbox-go/entity/model"
	"github.com/suzuito/geolocation-sandbox-go/web"
)

// PostBeaconLocation ...
func PostBeaconLocation(app web.Application) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		locations := []*model.Location{}
		if err := ctx.BindJSON(&locations); err != nil {
			responseError(ctx, &web.HTTPError{Status: 404, Message: err.Error()})
			return
		}
		cli, err := app.NewStoreCore(ctx)
		if err != nil {
			responseError(ctx, err)
			return
		}
		if err := cli.PutLocations(ctx, "dummyUser", locations); err != nil {
			responseError(ctx, web.FromStoreError(err))
			return
		}
		ctx.JSON(200, &locations)
	}
}
