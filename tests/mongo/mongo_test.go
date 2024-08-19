package mongo

import (
	"context"
	"testing"
	"time"

	"github.com/qiniu/qmgo"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestMongo_1(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	require.NoError(t, err)
	require.NotNil(t, client)

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

}

func TestMongo_qmgo_1(t *testing.T) {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: "mongodb://localhost:27017"})
	require.NoError(t, err)

	db := client.Database("class")
	coll := db.Collection("user")
	require.NotNil(t, coll)
}
