package dumongo

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"net/url"
)

type Config struct {
	Hosts      string `yaml:"hosts" toml:"hosts"`             // required (single：ip:port，shard cluster：ip:port,ip:port)
	Username   string `yaml:"username" toml:"username" `      // required
	Password   string `yaml:"password" toml:"password"`       // required
	DbName     string `yaml:"db_name" toml:"db_name"`         // required
	AuthSource string `yaml:"auth_source" toml:"auth_source"` // required (db for authentication)
	// SetMaxPoolSize specifies that maximum number of connections allowed in the driver's connection pool to each server.
	// Requests to a server will block if this maximum is reached. This can also be set through the "maxPoolSize" URI option
	// (e.g. "maxPoolSize=100"). The default is 100. If this is 0, it will be set to math.MaxInt64.
	MaxPoolSize int64 `yaml:"max_pool_size" toml:"max_pool_size"`
	// SetMinPoolSize specifies the minimum number of connections allowed in the driver's connection pool to each server. If
	// this is non-zero, each server's pool will be maintained in the background to ensure that the size does not fall below
	// the minimum. This can also be set through the "minPoolSize" URI option (e.g. "minPoolSize=100"). The default is 0.
	MinPoolSize int64 `yaml:"min_pool_size" toml:"min_pool_size"`
	// SetMaxConnIdleTime specifies the maximum amount of time that a connection will remain idle in a connection pool
	// before it is removed from the pool and closed. This can also be set through the "maxIdleTimeMS" URI option (e.g.
	// "maxIdleTimeMS=10000"). The default is 0, meaning a connection can remain unused indefinitely.
	MaxIdleTimeMS int64 `yaml:"max_idle_time_ms" toml:"max_idle_time_ms"`
	// SetMaxConnecting specifies the maximum number of connections a connection pool may establish simultaneously. This can
	// also be set through the "maxConnecting" URI option (e.g. "maxConnecting=2"). If this is 0, the default is used. The
	// default is 2. Values greater than 100 are not recommended.
	MaxConnecting int64 `yaml:"max_connecting" toml:"max_connecting"`
}

func NewMongo(config Config, opts ...*options.ClientOptions) (*mongo.Client, error) {
	if config.MaxPoolSize == 0 {
		config.MaxPoolSize = 100
	}
	pw := url.QueryEscape(config.Password)
	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s/%s?maxPoolSize=%d&authMechanism=SCRAM-SHA-1&authSource=%s&"+
			"minPoolSize=%d&maxIdleTimeMS=%d&maxConnecting=%d&readPreference=primaryPreferred",
		config.Username, pw, config.Hosts, config.DbName, config.MaxPoolSize, config.AuthSource,
		config.MinPoolSize, config.MaxIdleTimeMS, config.MaxConnecting)
	opts = append(opts, options.Client().ApplyURI(uri))

	client, err := mongo.Connect(context.TODO(), opts...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return client, nil
}

func CloseMongo(client *mongo.Client) error {
	err := client.Disconnect(context.TODO())
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
