package apiv1

import (
	"crypto"
	"crypto/x509"
	"time"

	"github.com/smallstep/certificates/kms/apiv1"
)

// CertificateAuthorityType indicates the type of Certificate Authority to
// create.
type CertificateAuthorityType int

const (
	// RootCA is the type used to create a self-signed certificate suitable for
	// use as a root CA.
	RootCA CertificateAuthorityType = iota + 1

	// IntermediateCA is the type used to create a subordinated certificate that
	// can be used to sign additional leaf certificates.
	IntermediateCA
)

// SignatureAlgorithm used for cryptographic signing.
type SignatureAlgorithm int

const (
	// Not specified.
	UnspecifiedSignAlgorithm SignatureAlgorithm = iota
	// RSASSA-PKCS1-v1_5 key and a SHA256 digest.
	SHA256WithRSA
	// RSASSA-PKCS1-v1_5 key and a SHA384 digest.
	SHA384WithRSA
	// RSASSA-PKCS1-v1_5 key and a SHA512 digest.
	SHA512WithRSA
	// RSASSA-PSS key with a SHA256 digest.
	SHA256WithRSAPSS
	// RSASSA-PSS key with a SHA384 digest.
	SHA384WithRSAPSS
	// RSASSA-PSS key with a SHA512 digest.
	SHA512WithRSAPSS
	// ECDSA on the NIST P-256 curve with a SHA256 digest.
	ECDSAWithSHA256
	// ECDSA on the NIST P-384 curve with a SHA384 digest.
	ECDSAWithSHA384
	// ECDSA on the NIST P-521 curve with a SHA512 digest.
	ECDSAWithSHA512
	// EdDSA on Curve25519 with a SHA512 digest.
	PureEd25519
)

// CreateCertificateRequest is the request used to sign a new certificate.
type CreateCertificateRequest struct {
	Template  *x509.Certificate
	Lifetime  time.Duration
	Backdate  time.Duration
	RequestID string
}

// CreateCertificateResponse is the response to a create certificate request.
type CreateCertificateResponse struct {
	Certificate      *x509.Certificate
	CertificateChain []*x509.Certificate
}

// RenewCertificateRequest is the request used to re-sign a certificate.
type RenewCertificateRequest struct {
	Template  *x509.Certificate
	Lifetime  time.Duration
	Backdate  time.Duration
	RequestID string
}

// RenewCertificateResponse is the response to a renew certificate request.
type RenewCertificateResponse struct {
	Certificate      *x509.Certificate
	CertificateChain []*x509.Certificate
}

// RevokeCertificateRequest is the request used to revoke a certificate.
type RevokeCertificateRequest struct {
	Certificate *x509.Certificate
	Reason      string
	ReasonCode  int
	RequestID   string
}

// RevokeCertificateResponse is the response to a revoke certificate request.
type RevokeCertificateResponse struct {
	Certificate      *x509.Certificate
	CertificateChain []*x509.Certificate
}

// GetCertificateAuthorityRequest is the request used to get the root
// certificate from a CAS.
type GetCertificateAuthorityRequest struct {
	Name string
}

// GetCertificateAuthorityResponse is the response that contains
// the root certificate.
type GetCertificateAuthorityResponse struct {
	RootCertificate *x509.Certificate
}

// CreateCertificateAuthorityRequest is the request used to generate a root or
// intermediate certificate.
type CreateCertificateAuthorityRequest struct {
	Name      string
	Type      CertificateAuthorityType
	Template  *x509.Certificate
	Lifetime  time.Duration
	Backdate  time.Duration
	RequestID string
	Project   string
	Location  string

	// Parent is the signer of the new CertificateAuthority.
	Parent *CreateCertificateAuthorityResponse

	// CreateKey defines the KMS CreateKeyRequest to use when creating a new
	// CertificateAuthority. If CreateKey is nil, a default algorithm will be
	// used.
	CreateKey *apiv1.CreateKeyRequest
}

// CreateCertificateAuthorityResponse is the response for
// CreateCertificateAuthority method and contains the root or intermediate
// certificate generated as well as the CA chain.
type CreateCertificateAuthorityResponse struct {
	Name             string
	Certificate      *x509.Certificate
	CertificateChain []*x509.Certificate
	PublicKey        crypto.PublicKey
	PrivateKey       crypto.PrivateKey
	Signer           crypto.Signer
}
