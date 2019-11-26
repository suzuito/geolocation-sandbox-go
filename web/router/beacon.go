package router

import (
	"github.com/gin-gonic/gin"
	"github.com/suzuito/geolocation-sandbox-go/entity/model"
	"github.com/suzuito/geolocation-sandbox-go/web"
)

// PostBeaconLocation ...
func PostBeaconLocation(app web.Application) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		locations := []model.Location{}
		ctx.BindJSON(&locations)
		ctx.JSON(200, &locations)
	}
}
