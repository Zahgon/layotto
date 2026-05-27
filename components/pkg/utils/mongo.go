// Copyright 2021 Layotto Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package utils

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

const (
	mongoHost        = "mongoHost"
	mongoPassword    = "mongoPassword"
	username         = "username"
	server           = "server"
	databaseName     = "databaseName"
	collecttionName  = "collectionName"
	writeConcern     = "writeConcern"
	readConcern      = "readConcern"
	operationTimeout = "operationTimeout"
	params           = "params"

	defaultDatabase       = "layottoStore"
	defaultCollectionName = "layottoCollection"
	defaultTimeout        = 5 * time.Second

	// mongodb://<username>:<password@<host>/<database><params>
	connectionURIFormatWithAuthentication = "mongodb://%s:%s@%s/%s%s"

	// mongodb://<host>/<database><params>
	connectionURIFormat = "mongodb://%s/%s%s"

	// mongodb+srv://<server>/<params>
	connectionURIFormatWithSrv = "mongodb+srv://%s/%s"
)

type MongoMetadata struct {
	Host             string
	Username         string
	Password         string
	DatabaseName     string
	CollectionName   string
	Server           string
	Params           string
	WriteConcern     string
	ReadConcern      string
	OperationTimeout time.Duration
}

// Item is Mongodb document wrapper.
type Item struct {
	Key   string      `bson:"_id"`
	Value interface{} `bson:"value"`
	Etag  string      `bson:"_etag"`
}

type MongoFactory interface {
	NewMongoClient(m MongoMetadata) (MongoClient, error)
	NewMongoCollection(m *mongo.Database, collectionName string, opts *options.CollectionOptions) MongoCollection
	NewSingleResult(sr *mongo.SingleResult) MongoSingleResult
}

type MongoClient interface {
	StartSession(opts ...*options.SessionOptions) (mongo.Session, error)
	Ping(ctx context.Context, rp *readpref.ReadPref) error
	Database(name string, opts ...*options.DatabaseOptions) *mongo.Database
	Disconnect(ctx context.Context) error
}

type MongoSession interface {
	AbortTransaction(context.Context) error
	CommitTransaction(context.Context) error
	WithTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) (interface{}, error),
		opts ...*options.TransactionOptions) (interface{}, error)
	EndSession(context.Context)
}

type MongoCollection interface {
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
	Indexes() mongo.IndexView
	UpdateOne(ctx context.Context, filter interface{}, update interface{},
		opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	FindOneAndUpdate(ctx context.Context, filter interface{},
		update interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult
}

type MongoSingleResult interface {
	Decode(v interface{}) error
	Err() error
	DecodeBytes() (bson.Raw, error)
}

type MongoFactoryImpl struct{}

func (c *MongoFactoryImpl) NewSingleResult(sr *mongo.SingleResult) MongoSingleResult {
	_ = "STUB: not implemented"
	return *new(MongoSingleResult)
}

func (c *MongoFactoryImpl) NewMongoCollection(m *mongo.Database, collectionName string, opts *options.CollectionOptions) MongoCollection {
	_ = "STUB: not implemented"
	return *new(MongoCollection)
}

func (c *MongoFactoryImpl) NewMongoClient(m MongoMetadata) (MongoClient, error) {
	_ = "STUB: not implemented"
	return *

	// Set client options
	new(MongoClient), nil
}

// Connect to MongoDB

func ParseMongoMetadata(properties map[string]string) (MongoMetadata, error) {
	_ = "STUB: not implemented"
	return *new(MongoMetadata), nil
}

func getString(properties map[string]string, key string) string {
	_ = "STUB: not implemented"
	return ""
}

func getMongoURI(m MongoMetadata) string { _ = "STUB: not implemented"; return "" }

func GetWriteConcernObject(cn string) (*writeconcern.WriteConcern, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetReadConcrenObject(cn string) (*readconcern.ReadConcern, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SetConcern(m MongoMetadata) (*options.CollectionOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// set mongo options of collection

func SetCollection(c MongoClient, f MongoFactory, m MongoMetadata) (MongoCollection, error) {
	_ = "STUB: not implemented"
	return *new(MongoCollection), nil
}
