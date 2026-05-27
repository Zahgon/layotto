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
package mock

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"mosn.io/layotto/components/pkg/utils"
)

type MockMongoFactory struct{}

// MockMongoClient is a mock of MongoClient interface
type MockMongoClient struct{}

// MockMongoSession is a mock of MongoSession interface
type MockMongoSession struct {
	mongo.SessionContext
}

// MockMongoCollection is a mock of MongoCollection interface
type MockMongoCollection struct {
	// '_id' document
	Result           map[string]bson.M
	InsertManyResult *mongo.InsertManyResult
	InsertOneResult  *mongo.InsertOneResult
	SingleResult     *mongo.SingleResult
	DeleteResult     *mongo.DeleteResult
}

func NewMockMongoFactory() *MockMongoFactory { _ = "STUB: not implemented"; return nil }

func NewMockMongoClient() *MockMongoClient { _ = "STUB: not implemented"; return nil }

func NewMockMongoCollection() *MockMongoCollection { _ = "STUB: not implemented"; return nil }

func NewMockMongoSession() *MockMongoSession { _ = "STUB: not implemented"; return nil }

func (f *MockMongoFactory) NewMongoClient(m utils.MongoMetadata) (utils.MongoClient, error) {
	_ = "STUB: not implemented"
	return *new(utils.MongoClient), nil
}

func (f *MockMongoFactory) NewMongoCollection(m *mongo.Database, collectionName string, opts *options.CollectionOptions) utils.MongoCollection {
	_ = "STUB: not implemented"
	return *new(utils.MongoCollection)
}

func (f *MockMongoFactory) NewSingleResult(sr *mongo.SingleResult) utils.MongoSingleResult {
	_ = "STUB: not implemented"
	return *new(utils.MongoSingleResult)
}

func (mc *MockMongoCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	_ = "STUB: not implemented"
	return nil
}

func (mc *MockMongoCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// insert cache

func (mc *MockMongoCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mc *MockMongoCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mc *MockMongoCollection) Indexes() mongo.IndexView {
	_ = "STUB: not implemented"
	return *new(mongo.IndexView)
}

func (mc *MockMongoCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mc *MockMongoCollection) FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult {
	_ = "STUB: not implemented"
	return nil
}

func (c *MockMongoClient) StartSession(opts ...*options.SessionOptions) (mongo.Session, error) {
	_ = "STUB: not implemented"
	return *new(mongo.Session), nil
}

func (c *MockMongoClient) Ping(ctx context.Context, rp *readpref.ReadPref) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *MockMongoClient) Database(name string, opts ...*options.DatabaseOptions) *mongo.Database {
	_ = "STUB: not implemented"
	return nil
}

func (c *MockMongoClient) Disconnect(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *MockMongoSession) AbortTransaction(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *MockMongoSession) CommitTransaction(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *MockMongoSession) WithTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) (interface{}, error),
	opts ...*options.TransactionOptions) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *MockMongoSession) EndSession(context.Context) { _ = "STUB: not implemented"; return }
