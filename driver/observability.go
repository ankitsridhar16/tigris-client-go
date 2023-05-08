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

package driver

import (
	"context"
	"fmt"

	"github.com/tigrisdata/tigris-client-go/config"
)

// Observability declares Tigris Observability APIs.
type Observability interface {
	QuotaLimits(ctx context.Context) (*QuotaLimits, error)
	QuotaUsage(ctx context.Context) (*QuotaUsage, error)

	Close() error
}

// NewObservability instantiates observability API client.
func NewObservability(ctx context.Context, cfg *config.Driver) (Observability, error) {
	var err error

	cfg, err = initConfig(cfg)
	if err != nil {
		return nil, err
	}

	proto := DefaultProtocol
	if cfg.Protocol != "" {
		proto = cfg.Protocol
	}

	initDrv, ok := drivers[proto]
	if !ok {
		return nil, fmt.Errorf("unsupported protocol")
	}

	_, _, o11y, err := initDrv(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return o11y, nil
}
