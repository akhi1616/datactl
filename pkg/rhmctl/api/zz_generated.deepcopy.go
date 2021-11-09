//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package api

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Config) DeepCopyInto(out *Config) {
	*out = *in
	in.MarketplaceEndpoint.DeepCopyInto(&out.MarketplaceEndpoint)
	if in.MeteringExports != nil {
		in, out := &in.MeteringExports, &out.MeteringExports
		*out = make(map[string]*MeteringExport, len(*in))
		for key, val := range *in {
			var outVal *MeteringExport
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(MeteringExport)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.DataServiceEndpoints != nil {
		in, out := &in.DataServiceEndpoints, &out.DataServiceEndpoints
		*out = make(map[string]*DataServiceEndpoint, len(*in))
		for key, val := range *in {
			var outVal *DataServiceEndpoint
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(DataServiceEndpoint)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Config.
func (in *Config) DeepCopy() *Config {
	if in == nil {
		return nil
	}
	out := new(Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Config) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataServiceEndpoint) DeepCopyInto(out *DataServiceEndpoint) {
	*out = *in
	if in.CertificateAuthorityData != nil {
		in, out := &in.CertificateAuthorityData, &out.CertificateAuthorityData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataServiceEndpoint.
func (in *DataServiceEndpoint) DeepCopy() *DataServiceEndpoint {
	if in == nil {
		return nil
	}
	out := new(DataServiceEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileInfo) DeepCopyInto(out *FileInfo) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(v1.Timestamp)
		**out = **in
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = new(v1.Timestamp)
		**out = **in
	}
	if in.DeletedAt != nil {
		in, out := &in.DeletedAt, &out.DeletedAt
		*out = new(v1.Timestamp)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileInfo.
func (in *FileInfo) DeepCopy() *FileInfo {
	if in == nil {
		return nil
	}
	out := new(FileInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FileInfo) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GetFileResponse) DeepCopyInto(out *GetFileResponse) {
	*out = *in
	if in.Info != nil {
		in, out := &in.Info, &out.Info
		*out = new(FileInfo)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GetFileResponse.
func (in *GetFileResponse) DeepCopy() *GetFileResponse {
	if in == nil {
		return nil
	}
	out := new(GetFileResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GetFileResponse) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListFilesResponse) DeepCopyInto(out *ListFilesResponse) {
	*out = *in
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make([]*FileInfo, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(FileInfo)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListFilesResponse.
func (in *ListFilesResponse) DeepCopy() *ListFilesResponse {
	if in == nil {
		return nil
	}
	out := new(ListFilesResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ListFilesResponse) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Marketplace) DeepCopyInto(out *Marketplace) {
	*out = *in
	if in.CertificateAuthorityData != nil {
		in, out := &in.CertificateAuthorityData, &out.CertificateAuthorityData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Marketplace.
func (in *Marketplace) DeepCopy() *Marketplace {
	if in == nil {
		return nil
	}
	out := new(Marketplace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeteringExport) DeepCopyInto(out *MeteringExport) {
	*out = *in
	out.Start = in.Start
	out.End = in.End
	if in.FileInfo != nil {
		in, out := &in.FileInfo, &out.FileInfo
		*out = make([]*MeteringFileSummary, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(MeteringFileSummary)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeteringExport.
func (in *MeteringExport) DeepCopy() *MeteringExport {
	if in == nil {
		return nil
	}
	out := new(MeteringExport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeteringFileSummary) DeepCopyInto(out *MeteringFileSummary) {
	*out = *in
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make([]*FileInfo, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(FileInfo)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeteringFileSummary.
func (in *MeteringFileSummary) DeepCopy() *MeteringFileSummary {
	if in == nil {
		return nil
	}
	out := new(MeteringFileSummary)
	in.DeepCopyInto(out)
	return out
}
