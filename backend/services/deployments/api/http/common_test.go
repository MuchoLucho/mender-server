// Copyright 2022 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package http

import (
	"github.com/gin-gonic/gin"

	"github.com/mendersoftware/mender-server/pkg/accesslog"
	"github.com/mendersoftware/mender-server/pkg/config"
	"github.com/mendersoftware/mender-server/pkg/requestid"

	dconfig "github.com/mendersoftware/mender-server/services/deployments/config"
)

func setUpTestRouter() *gin.Engine {
	router := gin.New()
	router.Use(accesslog.Middleware())
	router.Use(requestid.Middleware())

	return router
}

func init() {
	config.Config.SetDefault(
		dconfig.SettingStorageMaxImageSize,
		dconfig.SettingStorageMaxImageSizeDefault,
	)
}
