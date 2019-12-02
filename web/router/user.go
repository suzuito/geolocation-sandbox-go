package router

import (
	"github.com/gin-gonic/gin"
	"github.com/suzuito/geolocation-sandbox-go/entity/model"
	"github.com/suzuito/geolocation-sandbox-go/web"
)

// GetUserLocationsPartitions ...
func GetUserLocationsPartitions(app web.Application) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		partition := ctx.Param("partition")
		cli, err := app.NewStoreCore(ctx)
		if err != nil {
			responseError(ctx, err)
			return
		}
		locations := []*model.Location{}
		if err := cli.GetLocations(ctx, "dummyUser", partition, &locations); err != nil {
			responseError(ctx, web.FromStoreError(err))
			return
		}
		ctx.JSON(200, newLocations(partition, locations))
	}
}

// GetUserLocations ...
func GetUserLocations(app web.Application) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		currentPartition := ctx.DefaultQuery("current", "")
		cli, err := app.NewStoreCore(ctx)
		if err != nil {
			responseError(ctx, err)
			return
		}
		partitions := []string{}
		if err := cli.GetUserLocationsPartitions(ctx, "dummyUser", &partitions); err != nil {
			responseError(ctx, web.FromStoreError(err))
			return
		}
		if len(partitions) <= 0 {
			ctx.JSON(200, []string{})
			return
		}
		partition := ""
		if currentPartition == "" {
			partition = partitions[len(partitions)-1]
		} else {
			j := -1
			for i := range partitions {
				if partitions[i] == currentPartition {
					if i > 0 {
						j = i - 1
						break
					}
				}
			}
			if j < 0 {
				ctx.JSON(200, []string{})
				return
			}
			partition = partitions[j]
		}
		locations := []*model.Location{}
		if err := cli.GetLocations(ctx, "dummyUser", partition, &locations); err != nil {
			responseError(ctx, web.FromStoreError(err))
			return
		}
		ctx.JSON(200, newLocations(partition, locations))
	}
}
