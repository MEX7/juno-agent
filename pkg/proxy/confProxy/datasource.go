// Copyright 2020 Douyu
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package confProxy

import (
	"github.com/douyu/juno-agent/pkg/structs"
	"github.com/labstack/echo/v4"
)

// DataSource confu proxy dataSource interface ...
type DataSource interface {
	ListenAppConfig(ctx echo.Context, key string) chan *structs.ConfNode
	GetValues(ctx echo.Context, keys ...string) (map[string]string, error)
	GetRawValues(ctx echo.Context, rawKey string) (map[string]string, error)
	AppConfigScanner() []*structs.ConfNode
	Reload() error
	Stop()
}
