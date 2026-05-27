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

var Result = make(map[string]bson.M)
var id string

type MockMongoSequencerFactory struct{}

// MockMongoClient is a mock of MongoClient interface
type MockMongoSequencerClient struct{}

// MockMongoSession is a mock of MongoSession interface
type MockMongoSequencerSession struct {
	mongo.SessionContext
}

// MockMongoCollection is a mock of MongoCollection interface
type MockMongoSequencerCollection struct {
	// '_id' document
	InsertOneResult *mongo.InsertOneResult
	Result          map[string]bson.M
}

type MockMongoSequencerSingleResult struct{}

func NewMockMongoSequencerFactory() *MockMongoSequencerFactory {
	_ = "STUB: not implemented"
	return nil
}

func NewMockMongoSequencerSession() *MockMongoSequencerSession {
	_ = "STUB: not implemented"
	return nil
}

func (f *MockMongoSequencerFactory) NewSingleResult(sr *mongo.SingleResult) utils.MongoSingleResult {
	_ = "STUB: not implemented"
	return *new(utils.MongoSingleResult)
}

func (f *MockMongoSequencerFactory) NewMongoClient(m utils.MongoMetadata) (utils.MongoClient, error) {
	_ = "STUB: not implemented"
	return *new(utils.MongoClient), nil
}

func (f *MockMongoSequencerFactory) NewMongoCollection(m *mongo.Database, collectionName string, opts *options.CollectionOptions) utils.MongoCollection {
	_ = "STUB: not implemented"
	return *new(utils.MongoCollection)
}

func (mc *MockMongoSequencerCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	_ = "STUB: not implemented"
	return nil
}

func (mc *MockMongoSequencerCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// insert cache

func (mc *MockMongoSequencerCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mc *MockMongoSequencerCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mc *MockMongoSequencerCollection) Indexes() mongo.IndexView {
	_ = "STUB: not implemented"
	return *new(mongo.IndexView)
}

func (mc *MockMongoSequencerCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mc *MockMongoSequencerCollection) FindOneAndUpdate(ctx context.Context, filter interface{},
	update interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult {
	_ = "STUB: not implemented"
	return nil
}

func (c *MockMongoSequencerClient) StartSession(opts ...*options.SessionOptions) (mongo.Session, error) {
	_ = "STUB: not implemented"
	return *new(mongo.Session), nil
}

func (c *MockMongoSequencerClient) Ping(ctx context.Context, rp *readpref.ReadPref) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *MockMongoSequencerClient) Database(name string, opts ...*options.DatabaseOptions) *mongo.Database {
	_ = "STUB: not implemented"
	return nil
}

func (c *MockMongoSequencerClient) Disconnect(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *MockMongoSequencerSession) AbortTransaction(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *MockMongoSequencerSession) CommitTransaction(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *MockMongoSequencerSession) WithTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) (interface{}, error),
	opts ...*options.TransactionOptions) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *MockMongoSequencerSession) EndSession(context.Context) { _ = "STUB: not implemented"; return }

func (d *MockMongoSequencerSingleResult) Decode(v interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *MockMongoSequencerSingleResult) Err() error { _ = "STUB: not implemented"; return nil }

func (d *MockMongoSequencerSingleResult) DecodeBytes() (bson.Raw, error) {
	_ = "STUB: not implemented"
	return *new(bson.Raw), nil
}
