/*
 * Copyright 2021 Layotto Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ref

//Config is ref json config
type Config struct {
	ComponentRef     []*ComponentRef `json:"component_ref"`
	SecretRef        []*Item         `json:"secret_ref"`
	ConfigurationRef []*Item         `json:"configuration_ref"`
}

type ComponentRef struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type Item struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}
