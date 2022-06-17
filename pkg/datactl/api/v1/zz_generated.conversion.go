//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2022 IBM Corporation.
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

// Code generated by conversion-gen. DO NOT EDIT.

package v1

import (
	unsafe "unsafe"

	api "github.com/redhat-marketplace/datactl/pkg/datactl/api"
	dataservicev1 "github.com/redhat-marketplace/datactl/pkg/datactl/api/dataservice/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*DataServiceEndpoint)(nil), (*api.DataServiceEndpoint)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_DataServiceEndpoint_To_api_DataServiceEndpoint(a.(*DataServiceEndpoint), b.(*api.DataServiceEndpoint), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.DataServiceEndpoint)(nil), (*DataServiceEndpoint)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_DataServiceEndpoint_To_v1_DataServiceEndpoint(a.(*api.DataServiceEndpoint), b.(*DataServiceEndpoint), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ILMTEndpoint)(nil), (*api.ILMTEndpoint)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ILMTEndpoint_To_api_ILMTEndpoint(a.(*ILMTEndpoint), b.(*api.ILMTEndpoint), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.ILMTEndpoint)(nil), (*ILMTEndpoint)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_ILMTEndpoint_To_v1_ILMTEndpoint(a.(*api.ILMTEndpoint), b.(*ILMTEndpoint), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MeteringExport)(nil), (*api.MeteringExport)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_MeteringExport_To_api_MeteringExport(a.(*MeteringExport), b.(*api.MeteringExport), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.MeteringExport)(nil), (*MeteringExport)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_MeteringExport_To_v1_MeteringExport(a.(*api.MeteringExport), b.(*MeteringExport), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Source)(nil), (*api.Source)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Source_To_api_Source(a.(*Source), b.(*api.Source), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.Source)(nil), (*Source)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_Source_To_v1_Source(a.(*api.Source), b.(*Source), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*UploadAPI)(nil), (*api.UploadAPI)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_UploadAPI_To_api_UploadAPI(a.(*UploadAPI), b.(*api.UploadAPI), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.UploadAPI)(nil), (*UploadAPI)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_UploadAPI_To_v1_UploadAPI(a.(*api.UploadAPI), b.(*UploadAPI), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*api.Config)(nil), (*Config)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_Config_To_v1_Config(a.(*api.Config), b.(*Config), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*Config)(nil), (*api.Config)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Config_To_api_Config(a.(*Config), b.(*api.Config), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_Config_To_api_Config(in *Config, out *api.Config, s conversion.Scope) error {
	if err := Convert_v1_UploadAPI_To_api_UploadAPI(&in.MarketplaceEndpoint, &out.MarketplaceEndpoint, s); err != nil {
		return err
	}
	if in.CurrentMeteringExport != nil {
		in, out := &in.CurrentMeteringExport, &out.CurrentMeteringExport
		*out = new(api.MeteringExport)
		if err := Convert_v1_MeteringExport_To_api_MeteringExport(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.CurrentMeteringExport = nil
	}
	// WARNING: in.MeteringExports requires manual conversion: inconvertible types ([]*datactl/pkg/datactl/api/v1.MeteringExport vs map[string]*github.com/redhat-marketplace/datactl/pkg/datactl/api.MeteringExport)
	// WARNING: in.DataServiceEndpoints requires manual conversion: inconvertible types ([]*datactl/pkg/datactl/api/v1.DataServiceEndpoint vs map[string]*github.com/redhat-marketplace/datactl/pkg/datactl/api.DataServiceEndpoint)
	// WARNING: in.ILMTEndpoints requires manual conversion: inconvertible types ([]*datactl/pkg/datactl/api/v1.ILMTEndpoint vs map[string]*github.com/redhat-marketplace/datactl/pkg/datactl/api.ILMTEndpoint)
	// WARNING: in.Sources requires manual conversion: inconvertible types ([]*datactl/pkg/datactl/api/v1.Source vs map[string]*github.com/redhat-marketplace/datactl/pkg/datactl/api.Source)
	return nil
}

func autoConvert_api_Config_To_v1_Config(in *api.Config, out *Config, s conversion.Scope) error {
	if err := Convert_api_UploadAPI_To_v1_UploadAPI(&in.MarketplaceEndpoint, &out.MarketplaceEndpoint, s); err != nil {
		return err
	}
	if in.CurrentMeteringExport != nil {
		in, out := &in.CurrentMeteringExport, &out.CurrentMeteringExport
		*out = new(MeteringExport)
		if err := Convert_api_MeteringExport_To_v1_MeteringExport(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.CurrentMeteringExport = nil
	}
	// WARNING: in.MeteringExports requires manual conversion: inconvertible types (map[string]*github.com/redhat-marketplace/datactl/pkg/datactl/api.MeteringExport vs []*datactl/pkg/datactl/api/v1.MeteringExport)
	// WARNING: in.DataServiceEndpoints requires manual conversion: inconvertible types (map[string]*github.com/redhat-marketplace/datactl/pkg/datactl/api.DataServiceEndpoint vs []*datactl/pkg/datactl/api/v1.DataServiceEndpoint)
	// WARNING: in.ILMTEndpoints requires manual conversion: inconvertible types (map[string]*github.com/redhat-marketplace/datactl/pkg/datactl/api.ILMTEndpoint vs []*datactl/pkg/datactl/api/v1.ILMTEndpoint)
	// WARNING: in.Sources requires manual conversion: inconvertible types (map[string]*github.com/redhat-marketplace/datactl/pkg/datactl/api.Source vs []*datactl/pkg/datactl/api/v1.Source)
	return nil
}

func autoConvert_v1_DataServiceEndpoint_To_api_DataServiceEndpoint(in *DataServiceEndpoint, out *api.DataServiceEndpoint, s conversion.Scope) error {
	out.ClusterName = in.ClusterName
	out.Host = in.Host
	out.TokenData = in.TokenData
	out.TokenExpiration = in.TokenExpiration
	out.ServiceAccount = in.ServiceAccount
	out.Namespace = in.Namespace
	out.InsecureSkipTLSVerify = in.InsecureSkipTLSVerify
	out.CertificateAuthority = in.CertificateAuthority
	out.CertificateAuthorityData = *(*[]byte)(unsafe.Pointer(&in.CertificateAuthorityData))
	return nil
}

// Convert_v1_DataServiceEndpoint_To_api_DataServiceEndpoint is an autogenerated conversion function.
func Convert_v1_DataServiceEndpoint_To_api_DataServiceEndpoint(in *DataServiceEndpoint, out *api.DataServiceEndpoint, s conversion.Scope) error {
	return autoConvert_v1_DataServiceEndpoint_To_api_DataServiceEndpoint(in, out, s)
}

func autoConvert_api_DataServiceEndpoint_To_v1_DataServiceEndpoint(in *api.DataServiceEndpoint, out *DataServiceEndpoint, s conversion.Scope) error {
	// INFO: in.LocationOfOrigin opted out of conversion generation
	out.ClusterName = in.ClusterName
	out.Host = in.Host
	out.TokenData = in.TokenData
	out.TokenExpiration = in.TokenExpiration
	out.ServiceAccount = in.ServiceAccount
	out.Namespace = in.Namespace
	out.InsecureSkipTLSVerify = in.InsecureSkipTLSVerify
	out.CertificateAuthority = in.CertificateAuthority
	out.CertificateAuthorityData = *(*[]byte)(unsafe.Pointer(&in.CertificateAuthorityData))
	return nil
}

// Convert_api_DataServiceEndpoint_To_v1_DataServiceEndpoint is an autogenerated conversion function.
func Convert_api_DataServiceEndpoint_To_v1_DataServiceEndpoint(in *api.DataServiceEndpoint, out *DataServiceEndpoint, s conversion.Scope) error {
	return autoConvert_api_DataServiceEndpoint_To_v1_DataServiceEndpoint(in, out, s)
}

func autoConvert_v1_ILMTEndpoint_To_api_ILMTEndpoint(in *ILMTEndpoint, out *api.ILMTEndpoint, s conversion.Scope) error {
	// INFO: in.LocationOfOrigin opted out of conversion generation
	out.Host = in.Host
	out.Port = in.Port
	out.Token = in.Token
	out.LastPulldate = in.LastPulldate
	return nil
}

// Convert_v1_ILMTEndpoint_To_api_ILMTEndpoint is an autogenerated conversion function.
func Convert_v1_ILMTEndpoint_To_api_ILMTEndpoint(in *ILMTEndpoint, out *api.ILMTEndpoint, s conversion.Scope) error {
	return autoConvert_v1_ILMTEndpoint_To_api_ILMTEndpoint(in, out, s)
}

func autoConvert_api_ILMTEndpoint_To_v1_ILMTEndpoint(in *api.ILMTEndpoint, out *ILMTEndpoint, s conversion.Scope) error {
	// INFO: in.LocationOfOrigin opted out of conversion generation
	out.Host = in.Host
	out.Port = in.Port
	out.Token = in.Token
	out.LastPulldate = in.LastPulldate
	return nil
}

// Convert_api_ILMTEndpoint_To_v1_ILMTEndpoint is an autogenerated conversion function.
func Convert_api_ILMTEndpoint_To_v1_ILMTEndpoint(in *api.ILMTEndpoint, out *ILMTEndpoint, s conversion.Scope) error {
	return autoConvert_api_ILMTEndpoint_To_v1_ILMTEndpoint(in, out, s)
}

func autoConvert_v1_MeteringExport_To_api_MeteringExport(in *MeteringExport, out *api.MeteringExport, s conversion.Scope) error {
	out.FileName = in.FileName
	out.DataServiceCluster = in.DataServiceCluster
	out.Files = *(*[]*dataservicev1.FileInfoCTLAction)(unsafe.Pointer(&in.Files))
	return nil
}

// Convert_v1_MeteringExport_To_api_MeteringExport is an autogenerated conversion function.
func Convert_v1_MeteringExport_To_api_MeteringExport(in *MeteringExport, out *api.MeteringExport, s conversion.Scope) error {
	return autoConvert_v1_MeteringExport_To_api_MeteringExport(in, out, s)
}

func autoConvert_api_MeteringExport_To_v1_MeteringExport(in *api.MeteringExport, out *MeteringExport, s conversion.Scope) error {
	// INFO: in.LocationOfOrigin opted out of conversion generation
	out.FileName = in.FileName
	out.DataServiceCluster = in.DataServiceCluster
	out.Files = *(*[]*dataservicev1.FileInfoCTLAction)(unsafe.Pointer(&in.Files))
	// INFO: in.Committed opted out of conversion generation
	// INFO: in.Pushed opted out of conversion generation
	return nil
}

// Convert_api_MeteringExport_To_v1_MeteringExport is an autogenerated conversion function.
func Convert_api_MeteringExport_To_v1_MeteringExport(in *api.MeteringExport, out *MeteringExport, s conversion.Scope) error {
	return autoConvert_api_MeteringExport_To_v1_MeteringExport(in, out, s)
}

func autoConvert_v1_Source_To_api_Source(in *Source, out *api.Source, s conversion.Scope) error {
	out.Name = in.Name
	out.Type = api.SourceType(in.Type)
	out.LastAccessTime = in.LastAccessTime
	return nil
}

// Convert_v1_Source_To_api_Source is an autogenerated conversion function.
func Convert_v1_Source_To_api_Source(in *Source, out *api.Source, s conversion.Scope) error {
	return autoConvert_v1_Source_To_api_Source(in, out, s)
}

func autoConvert_api_Source_To_v1_Source(in *api.Source, out *Source, s conversion.Scope) error {
	// INFO: in.LocationOfOrigin opted out of conversion generation
	out.Name = in.Name
	out.Type = api.SourceType(in.Type)
	out.LastAccessTime = in.LastAccessTime
	return nil
}

// Convert_api_Source_To_v1_Source is an autogenerated conversion function.
func Convert_api_Source_To_v1_Source(in *api.Source, out *Source, s conversion.Scope) error {
	return autoConvert_api_Source_To_v1_Source(in, out, s)
}

func autoConvert_v1_UploadAPI_To_api_UploadAPI(in *UploadAPI, out *api.UploadAPI, s conversion.Scope) error {
	out.Host = in.Host
	out.PullSecret = in.PullSecret
	out.PullSecretData = in.PullSecretData
	out.InsecureSkipTLSVerify = in.InsecureSkipTLSVerify
	out.CertificateAuthority = in.CertificateAuthority
	out.CertificateAuthorityData = *(*[]byte)(unsafe.Pointer(&in.CertificateAuthorityData))
	out.ProxyURL = in.ProxyURL
	return nil
}

// Convert_v1_UploadAPI_To_api_UploadAPI is an autogenerated conversion function.
func Convert_v1_UploadAPI_To_api_UploadAPI(in *UploadAPI, out *api.UploadAPI, s conversion.Scope) error {
	return autoConvert_v1_UploadAPI_To_api_UploadAPI(in, out, s)
}

func autoConvert_api_UploadAPI_To_v1_UploadAPI(in *api.UploadAPI, out *UploadAPI, s conversion.Scope) error {
	// INFO: in.LocationOfOrigin opted out of conversion generation
	out.Host = in.Host
	out.PullSecret = in.PullSecret
	out.PullSecretData = in.PullSecretData
	out.InsecureSkipTLSVerify = in.InsecureSkipTLSVerify
	out.CertificateAuthority = in.CertificateAuthority
	out.CertificateAuthorityData = *(*[]byte)(unsafe.Pointer(&in.CertificateAuthorityData))
	out.ProxyURL = in.ProxyURL
	return nil
}

// Convert_api_UploadAPI_To_v1_UploadAPI is an autogenerated conversion function.
func Convert_api_UploadAPI_To_v1_UploadAPI(in *api.UploadAPI, out *UploadAPI, s conversion.Scope) error {
	return autoConvert_api_UploadAPI_To_v1_UploadAPI(in, out, s)
}


