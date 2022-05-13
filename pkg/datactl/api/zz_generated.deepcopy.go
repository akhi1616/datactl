//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package api

import (
	"github.com/redhat-marketplace/datactl/pkg/datactl/api/dataservice/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Config) DeepCopyInto(out *Config) {
	*out = *in
	in.MarketplaceEndpoint.DeepCopyInto(&out.MarketplaceEndpoint)
	if in.CurrentMeteringExport != nil {
		in, out := &in.CurrentMeteringExport, &out.CurrentMeteringExport
		*out = new(MeteringExport)
		(*in).DeepCopyInto(*out)
	}
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
	if in.Sources != nil {
		in, out := &in.Sources, &out.Sources
		*out = make(map[string]*Source, len(*in))
		for key, val := range *in {
			var outVal *Source
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(Source)
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
	in.TokenExpiration.DeepCopyInto(&out.TokenExpiration)
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
func (in *MeteringExport) DeepCopyInto(out *MeteringExport) {
	*out = *in
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make([]*v1.FileInfoCTLAction, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(v1.FileInfoCTLAction)
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
		*out = make([]*v1.FileInfoCTLAction, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(v1.FileInfoCTLAction)
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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Source) DeepCopyInto(out *Source) {
	*out = *in
	in.LastAccessTime.DeepCopyInto(&out.LastAccessTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Source.
func (in *Source) DeepCopy() *Source {
	if in == nil {
		return nil
	}
	out := new(Source)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UploadAPI) DeepCopyInto(out *UploadAPI) {
	*out = *in
	if in.CertificateAuthorityData != nil {
		in, out := &in.CertificateAuthorityData, &out.CertificateAuthorityData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UploadAPI.
func (in *UploadAPI) DeepCopy() *UploadAPI {
	if in == nil {
		return nil
	}
	out := new(UploadAPI)
	in.DeepCopyInto(out)
	return out
}
