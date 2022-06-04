package videostream

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"testing"

	miniredis "github.com/alicebob/miniredis"
	redismock "github.com/elliotchance/redismock/v8"
	redis "github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	yaml "gopkg.in/yaml.v3"
)

// newTestRedis returns a redis.Cmdable.
func newTestRedis() (*redis.Client, *redismock.ClientMock) {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	return client, redismock.NewNiceMock(client)
}

func TestGetErr(t *testing.T) {
	t.Run("error due to empty search result object response", func(t *testing.T) {
		redisClient, _ := newTestRedis()
		streamStorage := NewStorage(redisClient)
		_, err := streamStorage.Get(context.TODO(), "some id")
		assert.Error(t, err)
	})
}

// Writing VideoStreamPool and then trying to gather it
func TestGetOK(t *testing.T) {
	t.Run("error due to test search result dto doesnt received", func(t *testing.T) {

		redisClient, redisMock := newTestRedis()
		streamStorage := NewStorage(redisClient)

		//loading test streampool
		stream := VideoStreamPool{
			VideoId: "",
			Streams: []VideoStream{},
		}

		buf, err := ioutil.ReadFile("./test_data/data.yaml")
		require.NoError(t, err)

		err = yaml.Unmarshal([]byte(buf), &stream)
		require.NoError(t, err)
		ctx := context.Background()

		redisMock.HSet(ctx, "video_stream_pool:"+stream.VideoId, "videoId", stream.VideoId)
		data, _ := json.Marshal(stream.Streams)
		redisMock.HSet(ctx, "video_stream_pool:"+stream.VideoId, "streams", data)

		retrivenResult, err := streamStorage.Get(ctx, stream.VideoId)

		require.NoError(t, err)
		assert.Equal(t, stream, *retrivenResult, err)
	})
}

func TestCreatetOK(t *testing.T) {
	t.Run("error due to the video  dosn't properly created", func(t *testing.T) {
		redisClient, _ := newTestRedis()
		streamStorage := NewStorage(redisClient)

		//loading test streampool
		stream := VideoStreamPool{
			VideoId: "",
			Streams: []VideoStream{},
		}

		buf, err := ioutil.ReadFile("./test_data/data.yaml")
		require.NoError(t, err)

		err = yaml.Unmarshal([]byte(buf), &stream)
		require.NoError(t, err)

		ctx := context.Background()
		_, err = streamStorage.Get(ctx, stream.VideoId)
		assert.Error(t, err)

		err = streamStorage.Create(ctx, &stream)
		assert.NoError(t, err)

		// If it succesfully written then it has to be same as origin
		retrivenSearchResult, err := streamStorage.Get(ctx, stream.VideoId)
		require.NoError(t, err)
		assert.Equal(t, stream, *retrivenSearchResult, err)
	})
}
