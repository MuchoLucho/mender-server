// Copyright 2021 Northern.tech AS
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

package utils

import (
	"io"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

func ReadBodyRaw(r *http.Request) ([]byte, error) {
	if r.Body == nil {
		return nil, errors.New("empty request body")
	}
	content, err := io.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		return nil, err
	}

	return content, nil
}

func JoinURL(base, url string) string {
	url = strings.TrimPrefix(url, "/")
	if !strings.HasSuffix(base, "/") {
		base = base + "/"
	}
	return base + url

}
