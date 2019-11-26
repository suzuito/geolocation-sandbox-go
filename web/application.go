package web

import (
	"context"

	"github.com/suzuito/geolocation-sandbox-go/store/core"
)

// Application ...
type Application interface {
	NewStoreCore(ctx context.Context) (core.Client, *HTTPError)
}
