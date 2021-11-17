package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
import dataservicev1 "github.com/redhat-marketplace/rhmctl/pkg/rhmctl/api/dataservice/v1"

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Config struct {
	MarketplaceEndpoint Marketplace `json:"marketplace"`

	MeteringExports []*MeteringExport `json:"metering-export-history,omitempty"`

	DataServiceEndpoints []*DataServiceEndpoint `json:"data-service-endpoints"`
}

type MeteringExport struct {
	FileName string `json:"name"`

	// +optional
	DataServiceContext string `json:"data-service-context,omitempty"`

	// +optional
	Files []*dataservicev1.FileInfoCTLAction `json:"files,omitempty"`

	// DEPRECATED
	// +optional
	Start *metav1.Time `json:"start,omitempty"`

	// DEPRECATED
	// +optional
	End *metav1.Time `json:"end,omitempty"`

	// DEPRECATED
	// +optional
	Active bool `json:"active,omitempty"`

	// DEPRECATED
	// +optional
	DFileInfo []*MeteringFileSummary `json:"info,omitempty"`
}

// DEPRECATED
type MeteringFileSummary struct {
	DataServiceContext string `json:"data-service-context"`
	// +optional
	Files []*dataservicev1.FileInfoCTLAction `json:"files,omitempty"`
	// +optional
	Committed bool `json:"committed,omitempty"`
	// +optional
	Pushed bool `json:"pushed,omitempty"`
}

type Marketplace struct {
	// Host is the url of the marketplace i.e. marketplace.redhat.com
	Host string `json:"host"`

	// +optional
	PullSecret string `json:"pull-secret,omitempty"`

	// +optional
	PullSecretData string `json:"pull-secret-data,omitempty"`

	// InsecureSkipTLSVerify skips the validity check for the server's certificate. This will make your HTTPS connections insecure.
	// +optional
	InsecureSkipTLSVerify bool `json:"insecure-skip-tls-verify,omitempty"`

	// CertificateAuthority is the path to a cert file for the certificate authority.
	// +optional
	CertificateAuthority string `json:"certificate-authority,omitempty"`

	// CertificateAuthorityData contains PEM-encoded certificate authority certificates. Overrides CertificateAuthority
	// +optional
	CertificateAuthorityData []byte `json:"certificate-authority-data,omitempty"`

	// ProxyURL is the URL to the proxy to be used for all requests made by this
	// client. URLs with "http", "https", and "socks5" schemes are supported.  If
	// this configuration is not provided or the empty string, the client
	// attempts to construct a proxy configuration from http_proxy and
	// https_proxy environment variables. If these environment variables are not
	// set, the client does not attempt to proxy requests.
	//
	// socks5 proxying does not currently support spdy streaming endpoints (exec,
	// attach, port forward).
	// +optional
	ProxyURL string `json:"proxy-url,omitempty"`
}

type DataServiceEndpoint struct {
	ClusterContextName string `json:"cluster-context-name"`

	URL string `json:"url"`

	ServiceAccount string `json:"service-account,omitempty"`

	// Token is a filepath
	Token string `json:"token,omitempty"`

	TokenData string `json:"token-data,omitempty"`

	// InsecureSkipTLSVerify skips the validity check for the server's certificate. This will make your HTTPS connections insecure.
	// +optional
	InsecureSkipTLSVerify bool `json:"insecure-skip-tls-verify,omitempty"`

	// CertificateAuthority is the path to a cert file for the certificate authority.
	// +optional
	CertificateAuthority string `json:"certificate-authority,omitempty"`

	// CertificateAuthorityData contains PEM-encoded certificate authority certificates. Overrides CertificateAuthority
	// +optional
	CertificateAuthorityData []byte `json:"certificate-authority-data,omitempty"`

	// ProxyURL is the URL to the proxy to be used for all requests made by this
	// client. URLs with "http", "https", and "socks5" schemes are supported.  If
	// this configuration is not provided or the empty string, the client
	// attempts to construct a proxy configuration from http_proxy and
	// https_proxy environment variables. If these environment variables are not
	// set, the client does not attempt to proxy requests.
	//
	// socks5 proxying does not currently support spdy streaming endpoints (exec,
	// attach, port forward).
	// +optional
	ProxyURL string `json:"proxy-url,omitempty"`
}
