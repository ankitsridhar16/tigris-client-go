// Copyright 2022-2023 Tigris Data, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package timeseries

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/tigrisdata/tigris-client-go/driver"
	"github.com/tigrisdata/tigris-client-go/filter"
	"github.com/tigrisdata/tigris-client-go/mock"
	"github.com/tigrisdata/tigris-client-go/tigris"
)

var toDocument = driver.ToDocument

func TestTimeseries(t *testing.T) {
	type Coll1 struct {
		Model

		Field1 string `json:"field1"`
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ctrl := gomock.NewController(t)
	mdrv := mock.NewMockDriver(ctrl)

	mdb := mock.NewMockDatabase(ctrl)

	mdrv.EXPECT().UseDatabase("db1").Return(mdb)

	mdb.EXPECT().CreateOrUpdateCollections(gomock.Any(), driver.JAM(t, []string{`{"title":"coll_1","properties":{"timestamp":{"type":"string", "format":"date-time", "autoGenerate":true},"id":{"type":"string", "format":"uuid", "autoGenerate":true},"field1":{"type":"string"}},"primary_key":["timestamp", "id"],"collection_type":"documents"}`}))

	db, err := tigris.TestOpenDatabase(ctx, mdrv, "db1", &Coll1{})
	require.NoError(t, err)

	mdrv.EXPECT().UseDatabase("db1").Return(mdb)

	coll := GetCollection[Coll1](db)

	BufSize = 2

	ts := time.Now()

	d1 := &Coll1{Field1: "one"}

	// test append
	err = coll.Append(ctx, d1)
	require.NoError(t, err)

	require.Equal(t, 1, len(coll.buffer))

	d2 := &Coll1{Field1: "two"}

	mdb.EXPECT().Replace(ctx, "coll_1", []driver.Document{toDocument(t, d1), toDocument(t, d2)})

	err = coll.Append(ctx, d2)
	require.NoError(t, err)

	require.Equal(t, 0, len(coll.buffer))

	// test find after
	mit := mock.NewMockIterator(ctrl)

	b, err := json.Marshal(ts)
	require.NoError(t, err)

	mdb.EXPECT().Read(ctx, "coll_1",
		driver.Filter(`{"timestamp":{"$gte":`+string(b)+`}}`),
		driver.Projection(nil),
	).Return(mit, nil)

	it, err := coll.FindAfter(ctx, ts)
	require.NoError(t, err)

	var dd driver.Document

	mit.EXPECT().Next(&dd).SetArg(0, toDocument(t, d1)).Return(false)

	var d Coll1
	for it.Next(&d) {
		spew.Dump(d)
	}

	// test find after with filter
	mdb.EXPECT().Read(ctx, "coll_1",
		driver.Filter(`{"$and":[{"timestamp":{"$gte":`+string(b)+`}},{"field1":{"$eq":"one"}}]}`),
		driver.Projection(nil),
	).Return(mit, nil)

	it, err = coll.FindAfter(ctx, ts, filter.Eq("field1", "one"))
	require.NoError(t, err)

	mit.EXPECT().Next(&dd).SetArg(0, toDocument(t, d1)).Return(false)

	for it.Next(&d) {
		require.Equal(t, d1, d)
	}

	tsAfter := time.Now().Add(1 * time.Second)
	b1, err := json.Marshal(tsAfter)
	require.NoError(t, err)

	// test find before
	mdb.EXPECT().Read(ctx, "coll_1",
		driver.Filter(`{"timestamp":{"$lt":`+string(b1)+`}}`),
		driver.Projection(nil),
	).Return(mit, nil)

	it, err = coll.FindBefore(ctx, tsAfter)
	require.NoError(t, err)

	mit.EXPECT().Next(&dd).SetArg(0, toDocument(t, d1)).Return(false)

	for it.Next(&d) {
		require.Equal(t, d1, d)
	}

	// test find before with filter
	mdb.EXPECT().Read(ctx, "coll_1",
		driver.Filter(`{"$and":[{"timestamp":{"$lt":`+string(b1)+`}},{"field1":{"$eq":"one"}}]}`),
		driver.Projection(nil),
	).Return(mit, nil)

	it, err = coll.FindBefore(ctx, tsAfter, filter.Eq("field1", "one"))
	require.NoError(t, err)

	mit.EXPECT().Next(&dd).SetArg(0, toDocument(t, d1)).Return(false)

	for it.Next(&d) {
		require.Equal(t, d1, d)
	}

	// test find in range
	mdb.EXPECT().Read(ctx, "coll_1",
		driver.Filter(`{"$and":[{"timestamp":{"$gte":`+string(b)+`}},{"timestamp":{"$lt":`+string(b1)+`}}]}`),
		driver.Projection(nil),
	).Return(mit, nil)

	it, err = coll.FindInRange(ctx, ts, tsAfter)
	require.NoError(t, err)

	mit.EXPECT().Next(&dd).SetArg(0, toDocument(t, d1)).Return(false)

	for it.Next(&d) {
		require.Equal(t, d1, d)
	}

	// test find in range with filter
	mdb.EXPECT().Read(ctx, "coll_1",
		driver.Filter(`{"$and":[{"timestamp":{"$gte":`+string(b)+`}},{"timestamp":{"$lt":`+string(b1)+`}},{"field1":{"$eq":"one"}}]}`),
		driver.Projection(nil),
	).Return(mit, nil)

	it, err = coll.FindInRange(ctx, ts, tsAfter, filter.Eq("field1", "one"))
	require.NoError(t, err)

	mit.EXPECT().Next(&dd).SetArg(0, toDocument(t, d1)).Return(false)

	for it.Next(&d) {
		require.Equal(t, d1, d)
	}

	// test flush
	err = coll.Append(ctx, d1)
	require.NoError(t, err)

	require.Equal(t, 1, len(coll.buffer))

	mdb.EXPECT().Replace(ctx, "coll_1", []driver.Document{toDocument(t, d1)})

	err = coll.Flush(ctx)
	require.NoError(t, err)

	require.Equal(t, 0, len(coll.buffer))
}
