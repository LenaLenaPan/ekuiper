// Copyright 2022 EMQ Technologies Co., Ltd.
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

package json

import (
	"encoding/json"
)

type Converter struct {
}

var converter = &Converter{}

func GetConverter() (*Converter, error) {
	return converter, nil
}

func (c *Converter) Encode(d interface{}) ([]byte, error) {
	return json.Marshal(d)
}

func (c *Converter) Decode(b []byte) (interface{}, error) {
	result := make(map[string]interface{})
	e := json.Unmarshal(b, &result)
	return result, e
}
