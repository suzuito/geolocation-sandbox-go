package core

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/suzuito/geolocation-sandbox-go/entity/model"
	"github.com/suzuito/geolocation-sandbox-go/store"
)

var loc, _ = time.LoadLocation("Asia/Tokyo")

const (
	// NameUsers ...
	NameUsers = "users"
	// NameLocations ...
	NameLocations = "locations"
	// NamePartitions ...
	NamePartitions = "partitions"
	// NameDatas ...
	NameDatas = "datas"
)

// Client ...
type Client interface {
	PutLocations(ctx context.Context, userID string, locations []*model.Location) store.Error
	Close() store.Error
}

// ClientImpl ...
type ClientImpl struct {
	cli *firestore.Client
}

// NewClientImpl ...
func NewClientImpl(cli *firestore.Client) *ClientImpl {
	return &ClientImpl{
		cli: cli,
	}
}

// PutLocations ...
func (c *ClientImpl) PutLocations(ctx context.Context, userID string, locations []*model.Location) store.Error {
	userDoc := c.cli.Collection(NameUsers).Doc(userID)
	for _, location := range locations {
		t := time.Unix(location.Seconds, 0).In(loc)
		partition := t.Format("2006-01-02")
		partitionDoc := userDoc.Collection(NameLocations).Doc(partition)
	    doc := partitionDoc.Collection(NameDatas).Doc(location.ID)
		if _, err := doc.Set(ctx, location); err != nil {
			return store.NewErrorImpl(err)
		}
	}
	return nil
}

// Close ...
func (c *ClientImpl) Close() store.Error {
	if err := c.cli.Close(); err != nil {
		return store.NewErrorImpl(err)
	}
	return nil
}