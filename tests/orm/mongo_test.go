package orm

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func TestMongo_1(t *testing.T) {
	cc, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	require.NoError(t, err)
	require.NotNil(t, cc)

	defer func() {
		if err = cc.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}
