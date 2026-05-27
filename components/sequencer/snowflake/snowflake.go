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
package snowflake

import (
	"database/sql"
	"net"
	"time"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlHost         = "mysqlHost"
	mysqlDatabaseName = "databaseName"
	mysqlTableName    = "tableName"
	mysqlKeyTableName = "keyTableName"
	mysqlUserName     = "userName"
	mysqlPassword     = "password"
	mysqlCharset      = "utf8"
	timeBits          = "timeBits"
	workerBits        = "workerBits"
	seqBits           = "seqBits"
	startTime         = "startTime"
	reqTimeout        = "reqTimeout"
	keyTimeout        = "keyTimeout"

	defaultMysqlTableName = "layotto_sequencer_snowflake"
	defaultKeyTableName   = "layotto_sequencer_snowflake_key"
	defaultTimeBits       = 28
	defaultWorkerBits     = 22
	defaultSeqBits        = 13
	defaultStartTime      = "2022-01-01"
	defaultReqTimeout     = 500
	defaultKeyTimeout     = 24
)

type SnowflakeMetadata struct {
	MysqlMetadata SnowflakeMysqlMetadata

	WorkerBits     int64
	TimeBits       int64
	SeqBits        int64
	WorkidShift    int64
	TimestampShift int64
	StartTime      int64
	ReqTimeout     time.Duration
	KeyTimeout     time.Duration
	LogInfo        bool
}

type SnowflakeMysqlMetadata struct {
	//ip:port
	MysqlHost    string
	UserName     string
	Password     string
	DatabaseName string
	TableName    string
	KeyTableName string
	Db           *sql.DB
}

func ParseSnowflakeMysqlMetadata(properties map[string]string) (SnowflakeMysqlMetadata, error) {
	_ = "STUB: not implemented"
	return *new(SnowflakeMysqlMetadata), nil
}

func Parsebits(val string, defaultVal int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func Parsetime(val string, defaultVal int) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func ParseSnowflakeMetadata(properties map[string]string) (SnowflakeMetadata, error) {
	_ = "STUB: not implemented"
	return *new(SnowflakeMetadata), nil
}

func NewMysqlClient(meta *SnowflakeMysqlMetadata) (int64, error) {
	_ = "STUB: not implemented"

	//for unit test
	return 0, nil
}

// get id from mysql
// host_name = "ip"
// port = "timestamp-random number"
func NewWorkId(meta SnowflakeMysqlMetadata) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//insert a new record if the records are duplicated, to avoid clock rollback problems after shutdown

func MysqlRecord(db *sql.DB, keyTableName, key string, workerId, timestamp int64) error {
	_ = "STUB: not implemented"
	return nil
}

func getMysqlPort() string { _ = "STUB: not implemented"; return "" }

func getIP() (net.IP, error) { _ = "STUB: not implemented"; return *new(net.IP), nil }

func getIpFromAddr(addr net.Addr) net.IP { _ = "STUB: not implemented"; return *new(net.IP) }
