//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by conversion-gen. DO NOT EDIT.

package v1

import (
	unsafe "unsafe"

	api "github.com/redhat-marketplace/rhmctl/pkg/rhmctl/api"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	if err := s.AddGeneratedConversionFunc((*FileInfo)(nil), (*api.FileInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_FileInfo_To_api_FileInfo(a.(*FileInfo), b.(*api.FileInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.FileInfo)(nil), (*FileInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_FileInfo_To_v1_FileInfo(a.(*api.FileInfo), b.(*FileInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*GetFileResponse)(nil), (*api.GetFileResponse)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_GetFileResponse_To_api_GetFileResponse(a.(*GetFileResponse), b.(*api.GetFileResponse), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.GetFileResponse)(nil), (*GetFileResponse)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_GetFileResponse_To_v1_GetFileResponse(a.(*api.GetFileResponse), b.(*GetFileResponse), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ListFilesResponse)(nil), (*api.ListFilesResponse)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ListFilesResponse_To_api_ListFilesResponse(a.(*ListFilesResponse), b.(*api.ListFilesResponse), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.ListFilesResponse)(nil), (*ListFilesResponse)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_ListFilesResponse_To_v1_ListFilesResponse(a.(*api.ListFilesResponse), b.(*ListFilesResponse), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Marketplace)(nil), (*api.Marketplace)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Marketplace_To_api_Marketplace(a.(*Marketplace), b.(*api.Marketplace), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.Marketplace)(nil), (*Marketplace)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_Marketplace_To_v1_Marketplace(a.(*api.Marketplace), b.(*Marketplace), scope)
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
	if err := s.AddGeneratedConversionFunc((*MeteringFileSummary)(nil), (*api.MeteringFileSummary)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_MeteringFileSummary_To_api_MeteringFileSummary(a.(*MeteringFileSummary), b.(*api.MeteringFileSummary), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.MeteringFileSummary)(nil), (*MeteringFileSummary)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_MeteringFileSummary_To_v1_MeteringFileSummary(a.(*api.MeteringFileSummary), b.(*MeteringFileSummary), scope)
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
	if err := Convert_v1_Marketplace_To_api_Marketplace(&in.MarketplaceEndpoint, &out.MarketplaceEndpoint, s); err != nil {
		return err
	}
	// WARNING: in.MeteringExports requires manual conversion: inconvertible types ([]*./v1.MeteringExport vs map[string]*github.com/redhat-marketplace/rhmctl/pkg/rhmctl/api.MeteringExport)
	// WARNING: in.DataServiceEndpoints requires manual conversion: inconvertible types ([]*./v1.DataServiceEndpoint vs map[string]*github.com/redhat-marketplace/rhmctl/pkg/rhmctl/api.DataServiceEndpoint)
	return nil
}

func autoConvert_api_Config_To_v1_Config(in *api.Config, out *Config, s conversion.Scope) error {
	if err := Convert_api_Marketplace_To_v1_Marketplace(&in.MarketplaceEndpoint, &out.MarketplaceEndpoint, s); err != nil {
		return err
	}
	// WARNING: in.MeteringExports requires manual conversion: inconvertible types (map[string]*github.com/redhat-marketplace/rhmctl/pkg/rhmctl/api.MeteringExport vs []*./v1.MeteringExport)
	// WARNING: in.DataServiceEndpoints requires manual conversion: inconvertible types (map[string]*github.com/redhat-marketplace/rhmctl/pkg/rhmctl/api.DataServiceEndpoint vs []*./v1.DataServiceEndpoint)
	return nil
}

func autoConvert_v1_DataServiceEndpoint_To_api_DataServiceEndpoint(in *DataServiceEndpoint, out *api.DataServiceEndpoint, s conversion.Scope) error {
	out.ClusterContextName = in.ClusterContextName
	out.URL = in.URL
	out.ServiceAccount = in.ServiceAccount
	out.Token = in.Token
	out.TokenData = in.TokenData
	out.InsecureSkipTLSVerify = in.InsecureSkipTLSVerify
	out.CertificateAuthority = in.CertificateAuthority
	out.CertificateAuthorityData = *(*[]byte)(unsafe.Pointer(&in.CertificateAuthorityData))
	out.ProxyURL = in.ProxyURL
	return nil
}

// Convert_v1_DataServiceEndpoint_To_api_DataServiceEndpoint is an autogenerated conversion function.
func Convert_v1_DataServiceEndpoint_To_api_DataServiceEndpoint(in *DataServiceEndpoint, out *api.DataServiceEndpoint, s conversion.Scope) error {
	return autoConvert_v1_DataServiceEndpoint_To_api_DataServiceEndpoint(in, out, s)
}

func autoConvert_api_DataServiceEndpoint_To_v1_DataServiceEndpoint(in *api.DataServiceEndpoint, out *DataServiceEndpoint, s conversion.Scope) error {
	// INFO: in.LocationOfOrigin opted out of conversion generation
	out.ClusterContextName = in.ClusterContextName
	out.URL = in.URL
	out.Token = in.Token
	out.TokenData = in.TokenData
	out.ServiceAccount = in.ServiceAccount
	out.InsecureSkipTLSVerify = in.InsecureSkipTLSVerify
	out.CertificateAuthority = in.CertificateAuthority
	out.CertificateAuthorityData = *(*[]byte)(unsafe.Pointer(&in.CertificateAuthorityData))
	out.ProxyURL = in.ProxyURL
	return nil
}

// Convert_api_DataServiceEndpoint_To_v1_DataServiceEndpoint is an autogenerated conversion function.
func Convert_api_DataServiceEndpoint_To_v1_DataServiceEndpoint(in *api.DataServiceEndpoint, out *DataServiceEndpoint, s conversion.Scope) error {
	return autoConvert_api_DataServiceEndpoint_To_v1_DataServiceEndpoint(in, out, s)
}

func autoConvert_v1_FileInfo_To_api_FileInfo(in *FileInfo, out *api.FileInfo, s conversion.Scope) error {
	out.Id = in.Id
	out.Name = in.Name
	out.Size = in.Size
	out.Source = in.Source
	out.SourceType = in.SourceType
	out.Checksum = in.Checksum
	out.MimeType = in.MimeType
	out.CreatedAt = (*metav1.Time)(unsafe.Pointer(in.CreatedAt))
	out.UpdatedAt = (*metav1.Time)(unsafe.Pointer(in.UpdatedAt))
	out.DeletedAt = (*metav1.Time)(unsafe.Pointer(in.DeletedAt))
	out.Metadata = *(*map[string]string)(unsafe.Pointer(&in.Metadata))
	out.UploadID = in.UploadID
	out.UploadError = in.UploadError
	return nil
}

// Convert_v1_FileInfo_To_api_FileInfo is an autogenerated conversion function.
func Convert_v1_FileInfo_To_api_FileInfo(in *FileInfo, out *api.FileInfo, s conversion.Scope) error {
	return autoConvert_v1_FileInfo_To_api_FileInfo(in, out, s)
}

func autoConvert_api_FileInfo_To_v1_FileInfo(in *api.FileInfo, out *FileInfo, s conversion.Scope) error {
	out.Id = in.Id
	out.Name = in.Name
	out.Size = in.Size
	out.Source = in.Source
	out.SourceType = in.SourceType
	out.Checksum = in.Checksum
	out.MimeType = in.MimeType
	out.CreatedAt = (*metav1.Time)(unsafe.Pointer(in.CreatedAt))
	out.UpdatedAt = (*metav1.Time)(unsafe.Pointer(in.UpdatedAt))
	out.DeletedAt = (*metav1.Time)(unsafe.Pointer(in.DeletedAt))
	out.Metadata = *(*map[string]string)(unsafe.Pointer(&in.Metadata))
	out.UploadID = in.UploadID
	out.UploadError = in.UploadError
	return nil
}

// Convert_api_FileInfo_To_v1_FileInfo is an autogenerated conversion function.
func Convert_api_FileInfo_To_v1_FileInfo(in *api.FileInfo, out *FileInfo, s conversion.Scope) error {
	return autoConvert_api_FileInfo_To_v1_FileInfo(in, out, s)
}

func autoConvert_v1_GetFileResponse_To_api_GetFileResponse(in *GetFileResponse, out *api.GetFileResponse, s conversion.Scope) error {
	out.Info = (*api.FileInfo)(unsafe.Pointer(in.Info))
	return nil
}

// Convert_v1_GetFileResponse_To_api_GetFileResponse is an autogenerated conversion function.
func Convert_v1_GetFileResponse_To_api_GetFileResponse(in *GetFileResponse, out *api.GetFileResponse, s conversion.Scope) error {
	return autoConvert_v1_GetFileResponse_To_api_GetFileResponse(in, out, s)
}

func autoConvert_api_GetFileResponse_To_v1_GetFileResponse(in *api.GetFileResponse, out *GetFileResponse, s conversion.Scope) error {
	out.Info = (*FileInfo)(unsafe.Pointer(in.Info))
	return nil
}

// Convert_api_GetFileResponse_To_v1_GetFileResponse is an autogenerated conversion function.
func Convert_api_GetFileResponse_To_v1_GetFileResponse(in *api.GetFileResponse, out *GetFileResponse, s conversion.Scope) error {
	return autoConvert_api_GetFileResponse_To_v1_GetFileResponse(in, out, s)
}

func autoConvert_v1_ListFilesResponse_To_api_ListFilesResponse(in *ListFilesResponse, out *api.ListFilesResponse, s conversion.Scope) error {
	out.Files = *(*[]*api.FileInfo)(unsafe.Pointer(&in.Files))
	out.NextPageToken = in.NextPageToken
	out.PageSize = in.PageSize
	return nil
}

// Convert_v1_ListFilesResponse_To_api_ListFilesResponse is an autogenerated conversion function.
func Convert_v1_ListFilesResponse_To_api_ListFilesResponse(in *ListFilesResponse, out *api.ListFilesResponse, s conversion.Scope) error {
	return autoConvert_v1_ListFilesResponse_To_api_ListFilesResponse(in, out, s)
}

func autoConvert_api_ListFilesResponse_To_v1_ListFilesResponse(in *api.ListFilesResponse, out *ListFilesResponse, s conversion.Scope) error {
	out.Files = *(*[]*FileInfo)(unsafe.Pointer(&in.Files))
	out.NextPageToken = in.NextPageToken
	out.PageSize = in.PageSize
	return nil
}

// Convert_api_ListFilesResponse_To_v1_ListFilesResponse is an autogenerated conversion function.
func Convert_api_ListFilesResponse_To_v1_ListFilesResponse(in *api.ListFilesResponse, out *ListFilesResponse, s conversion.Scope) error {
	return autoConvert_api_ListFilesResponse_To_v1_ListFilesResponse(in, out, s)
}

func autoConvert_v1_Marketplace_To_api_Marketplace(in *Marketplace, out *api.Marketplace, s conversion.Scope) error {
	out.Host = in.Host
	out.PullSecret = in.PullSecret
	out.PullSecretData = in.PullSecretData
	out.InsecureSkipTLSVerify = in.InsecureSkipTLSVerify
	out.CertificateAuthority = in.CertificateAuthority
	out.CertificateAuthorityData = *(*[]byte)(unsafe.Pointer(&in.CertificateAuthorityData))
	out.ProxyURL = in.ProxyURL
	return nil
}

// Convert_v1_Marketplace_To_api_Marketplace is an autogenerated conversion function.
func Convert_v1_Marketplace_To_api_Marketplace(in *Marketplace, out *api.Marketplace, s conversion.Scope) error {
	return autoConvert_v1_Marketplace_To_api_Marketplace(in, out, s)
}

func autoConvert_api_Marketplace_To_v1_Marketplace(in *api.Marketplace, out *Marketplace, s conversion.Scope) error {
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

// Convert_api_Marketplace_To_v1_Marketplace is an autogenerated conversion function.
func Convert_api_Marketplace_To_v1_Marketplace(in *api.Marketplace, out *Marketplace, s conversion.Scope) error {
	return autoConvert_api_Marketplace_To_v1_Marketplace(in, out, s)
}

func autoConvert_v1_MeteringExport_To_api_MeteringExport(in *MeteringExport, out *api.MeteringExport, s conversion.Scope) error {
	out.FileName = in.FileName
	out.Active = in.Active
	out.Start = in.Start
	out.End = in.End
	out.FileInfo = *(*[]*api.MeteringFileSummary)(unsafe.Pointer(&in.FileInfo))
	return nil
}

// Convert_v1_MeteringExport_To_api_MeteringExport is an autogenerated conversion function.
func Convert_v1_MeteringExport_To_api_MeteringExport(in *MeteringExport, out *api.MeteringExport, s conversion.Scope) error {
	return autoConvert_v1_MeteringExport_To_api_MeteringExport(in, out, s)
}

func autoConvert_api_MeteringExport_To_v1_MeteringExport(in *api.MeteringExport, out *MeteringExport, s conversion.Scope) error {
	// INFO: in.LocationOfOrigin opted out of conversion generation
	out.FileName = in.FileName
	out.Active = in.Active
	out.Start = in.Start
	out.End = in.End
	out.FileInfo = *(*[]*MeteringFileSummary)(unsafe.Pointer(&in.FileInfo))
	return nil
}

// Convert_api_MeteringExport_To_v1_MeteringExport is an autogenerated conversion function.
func Convert_api_MeteringExport_To_v1_MeteringExport(in *api.MeteringExport, out *MeteringExport, s conversion.Scope) error {
	return autoConvert_api_MeteringExport_To_v1_MeteringExport(in, out, s)
}

func autoConvert_v1_MeteringFileSummary_To_api_MeteringFileSummary(in *MeteringFileSummary, out *api.MeteringFileSummary, s conversion.Scope) error {
	out.DataServiceContext = in.DataServiceContext
	out.Files = *(*[]*api.FileInfo)(unsafe.Pointer(&in.Files))
	out.Committed = in.Committed
	out.Pushed = in.Pushed
	return nil
}

// Convert_v1_MeteringFileSummary_To_api_MeteringFileSummary is an autogenerated conversion function.
func Convert_v1_MeteringFileSummary_To_api_MeteringFileSummary(in *MeteringFileSummary, out *api.MeteringFileSummary, s conversion.Scope) error {
	return autoConvert_v1_MeteringFileSummary_To_api_MeteringFileSummary(in, out, s)
}

func autoConvert_api_MeteringFileSummary_To_v1_MeteringFileSummary(in *api.MeteringFileSummary, out *MeteringFileSummary, s conversion.Scope) error {
	out.DataServiceContext = in.DataServiceContext
	out.Files = *(*[]*FileInfo)(unsafe.Pointer(&in.Files))
	out.Committed = in.Committed
	out.Pushed = in.Pushed
	return nil
}

// Convert_api_MeteringFileSummary_To_v1_MeteringFileSummary is an autogenerated conversion function.
func Convert_api_MeteringFileSummary_To_v1_MeteringFileSummary(in *api.MeteringFileSummary, out *MeteringFileSummary, s conversion.Scope) error {
	return autoConvert_api_MeteringFileSummary_To_v1_MeteringFileSummary(in, out, s)
}
