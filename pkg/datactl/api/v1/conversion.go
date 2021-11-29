// Copyright 2021 IBM Corporation.
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

package v1

import (
	api "github.com/redhat-marketplace/datactl/pkg/datactl/api"
	"k8s.io/apimachinery/pkg/conversion"
)

func Convert_v1_Config_To_api_Config(a *Config, b *api.Config, scope conversion.Scope) error {
	err := autoConvert_v1_Config_To_api_Config(a, b, scope)

	if err != nil {
		return err
	}

	b.DataServiceEndpoints = make(map[string]*api.DataServiceEndpoint)

	for _, aD := range a.DataServiceEndpoints {
		bD := &api.DataServiceEndpoint{}

		err := autoConvert_v1_DataServiceEndpoint_To_api_DataServiceEndpoint(aD, bD, scope)
		if err != nil {
			return err
		}
		b.DataServiceEndpoints[aD.ClusterName] = bD
	}

	b.MeteringExports = make(map[string]*api.MeteringExport)

	for _, aD := range a.MeteringExports {
		bD := &api.MeteringExport{}

		err := autoConvert_v1_MeteringExport_To_api_MeteringExport(aD, bD, scope)
		if err != nil {
			return err
		}

		b.MeteringExports[bD.DataServiceCluster] = bD
	}

	return nil
}

func Convert_api_Config_To_v1_Config(a *api.Config, b *Config, scope conversion.Scope) error {
	err := autoConvert_api_Config_To_v1_Config(a, b, scope)

	if err != nil {
		return err
	}

	b.DataServiceEndpoints = make([]*DataServiceEndpoint, 0, len(a.DataServiceEndpoints))

	for _, aD := range a.DataServiceEndpoints {
		bD := &DataServiceEndpoint{}

		err := autoConvert_api_DataServiceEndpoint_To_v1_DataServiceEndpoint(aD, bD, scope)
		if err != nil {
			return err
		}
		b.DataServiceEndpoints = append(b.DataServiceEndpoints, bD)
	}

	b.MeteringExports = make([]*MeteringExport, 0, len(a.MeteringExports))

	for _, aD := range a.MeteringExports {
		bD := &MeteringExport{}

		err := autoConvert_api_MeteringExport_To_v1_MeteringExport(aD, bD, scope)
		if err != nil {
			return err
		}
		b.MeteringExports = append(b.MeteringExports, bD)
	}

	return nil
}
