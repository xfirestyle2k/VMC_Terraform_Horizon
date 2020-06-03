/*
 * NSX API
 *
 * VMware NSX REST API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package loadbalancer

// SSL cipher
type LbSslCipherInfo struct {

	// SSL cipher
	Cipher string `json:"cipher"`

	// Default SSL cipher flag
	IsDefault bool `json:"is_default"`

	// Secure/insecure SSL cipher flag
	IsSecure bool `json:"is_secure"`
}
