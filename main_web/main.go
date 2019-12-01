package main

import (
	"context"
	"os"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/suzuito/geolocation-sandbox-go/store/core"
	"github.com/suzuito/geolocation-sandbox-go/web"
	"github.com/suzuito/geolocation-sandbox-go/web/router"
	"google.golang.org/api/option"
)

// ApplicationImpl ...
type ApplicationImpl struct {
	appFirebase *firebase.App
}

// NewStoreCore ...
func (a *ApplicationImpl) NewStoreCore(ctx context.Context) (core.Client, *web.HTTPError) {
	cli, err := a.appFirebase.Firestore(ctx)
	if err != nil {
		return nil, web.NewNewStoreCoreError(err)
	}
	return core.NewClientImpl(cli), nil
}

func main() {
	ctx := context.Background()
	sa := option.WithCredentialsFile("gls-minilla-4441b18f2ad3.json")
	// Firestore app
	firebaseConf := &firebase.Config{ProjectID: os.Getenv("GOOGLE_CLOUD_PROJECT")}
	appFirebase, err := firebase.NewApp(ctx, firebaseConf, sa)
	if err != nil {
		panic(err)
	}
	// App
	app := ApplicationImpl{
		appFirebase: appFirebase,
	}
	// Router
	root := gin.Default()
	root.POST("/beacon", router.PostBeaconLocation(&app))
	root.GET("/locations/partitions/:partition", router.GetUserLocationsPartitions(&app))
	root.GET("/locations/latest", router.GetUserLocationsLatest(&app))
	root.Run()
}
