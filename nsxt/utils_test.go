/* Copyright © 2018 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: MPL-2.0 */

package nsxt

import (
	"fmt"
	"github.com/vmware/go-vmware-nsxt/trust"
	"net/http"
	"os"
	"testing"
)

// Default names or prefixed of NSX backend existing objects used in the acceptance tests.
// Those defaults can be overridden using environment parameters
const tier0RouterDefaultName string = "PLR-1 LogicalRouterTier0"
const edgeClusterDefaultName string = "edgecluster1"
const vlanTransportZoneName string = "transportzone2"
const overlayTransportZoneNamePrefix string = "1-transportzone"
const macPoolDefaultName string = "DefaultMacPool"

const realizationResourceName string = "data.nsxt_policy_realization_info.realization_info"

const singleTag string = `
  tag {
    scope = "scope1"
    tag   = "tag1"
  }`

const doubleTags string = `
  tag {
    scope = "scope1"
    tag   = "tag1"
  }
  tag {
    scope = "scope2"
    tag   = "tag2"
  }`

func getTier0RouterName() string {
	name := os.Getenv("NSXT_TEST_TIER0_ROUTER")
	if name == "" {
		name = tier0RouterDefaultName
	}
	return name
}

func getEdgeClusterName() string {
	name := os.Getenv("NSXT_TEST_EDGE_CLUSTER")
	if name == "" {
		name = edgeClusterDefaultName
	}
	return name
}

func getVlanTransportZoneName() string {
	name := os.Getenv("NSXT_TEST_VLAN_TRANSPORT_ZONE")
	if name == "" {
		name = vlanTransportZoneName
	}
	return name
}

func getOverlayTransportZoneName() string {
	name := os.Getenv("NSXT_TEST_OVERLAY_TRANSPORT_ZONE")
	if name == "" {
		name = overlayTransportZoneNamePrefix
	}
	return name
}

func getMacPoolName() string {
	name := os.Getenv("NSXT_TEST_MAC_POOL")
	if name == "" {
		name = macPoolDefaultName
	}
	return name
}

func getIPPoolName() string {
	return os.Getenv("NSXT_TEST_IP_POOL")
}

func getTestVMID() string {
	return os.Getenv("NSXT_TEST_VM_ID")
}

func getTestVMName() string {
	return os.Getenv("NSXT_TEST_VM_NAME")
}

func getTestCertificateName(isClient bool) string {
	if isClient {
		return os.Getenv("NSXT_TEST_CLIENT_CERTIFICATE_NAME")
	}
	return os.Getenv("NSXT_TEST_CERTIFICATE_NAME")
}

func testAccEnvDefined(t *testing.T, envVar string) {
	if len(os.Getenv(envVar)) == 0 {
		t.Skipf("This test requires %s environment variable to be set", envVar)
	}
}

// Create and delete CA and client cert for various tests
func testAccNSXCreateCert(t *testing.T, name string, certPem string, certPK string, certType string) string {
	nsxClient, err := testAccGetClient()
	if err != nil {
		t.Fatal(err)
	}

	object := trust.TrustObjectData{
		DisplayName:  name,
		ResourceType: certType,
		PemEncoded:   certPem,
		PrivateKey:   certPK,
	}

	certList, response, err := nsxClient.NsxComponentAdministrationApi.AddCertificateImport(nsxClient.Context, object)

	if err != nil {
		t.Fatal(fmt.Sprintf("Error while creating %s certificate. Error: %v", certType, err))
	}
	if response.StatusCode != http.StatusCreated {
		t.Fatal(fmt.Errorf("Error while creating %s certificate. HTTP return code %d", certType, response.StatusCode))
	}
	certID := ""
	for _, cert := range certList.Results {
		certID = cert.Id
	}

	return certID
}

func testAccNSXDeleteCert(t *testing.T, id string) {
	nsxClient, err := testAccGetClient()
	if err != nil {
		t.Fatal(err)
	}

	response, err := nsxClient.NsxComponentAdministrationApi.DeleteCertificate(nsxClient.Context, id)

	if err != nil {
		t.Fatal(fmt.Sprintf("Error while deleting certificate %s. Error: %v", id, err))
	}
	if response.StatusCode != http.StatusOK {
		t.Fatal(fmt.Errorf("Error while deleting certificate %s. HTTP return code %d", id, response.StatusCode))
	}
}

func testAccNSXCreateCerts(t *testing.T) (string, string, string) {

	certPem := "-----BEGIN CERTIFICATE-----\n" +
		"MIICVjCCAb8CAg37MA0GCSqGSIb3DQEBBQUAMIGbMQswCQYDVQQGEwJKUDEOMAwG\n" +
		"A1UECBMFVG9reW8xEDAOBgNVBAcTB0NodW8ta3UxETAPBgNVBAoTCEZyYW5rNERE\n" +
		"MRgwFgYDVQQLEw9XZWJDZXJ0IFN1cHBvcnQxGDAWBgNVBAMTD0ZyYW5rNEREIFdl\n" +
		"YiBDQTEjMCEGCSqGSIb3DQEJARYUc3VwcG9ydEBmcmFuazRkZC5jb20wHhcNMTIw\n" +
		"ODIyMDUyNzIzWhcNMTcwODIxMDUyNzIzWjBKMQswCQYDVQQGEwJKUDEOMAwGA1UE\n" +
		"CAwFVG9reW8xETAPBgNVBAoMCEZyYW5rNEREMRgwFgYDVQQDDA93d3cuZXhhbXBs\n" +
		"ZS5jb20wgZ8wDQYJKoZIhvcNAQEBBQADgY0AMIGJAoGBAMYBBrx5PlP0WNI/ZdzD\n" +
		"+6Pktmurn+F2kQYbtc7XQh8/LTBvCo+P6iZoLEmUA9e7EXLRxgU1CVqeAi7QcAn9\n" +
		"MwBlc8ksFJHB0rtf9pmf8Oza9E0Bynlq/4/Kb1x+d+AyhL7oK9tQwB24uHOueHi1\n" +
		"C/iVv8CSWKiYe6hzN1txYe8rAgMBAAEwDQYJKoZIhvcNAQEFBQADgYEAASPdjigJ\n" +
		"kXCqKWpnZ/Oc75EUcMi6HztaW8abUMlYXPIgkV2F7YanHOB7K4f7OOLjiz8DTPFf\n" +
		"jC9UeuErhaA/zzWi8ewMTFZW/WshOrm3fNvcMrMLKtH534JKvcdMg6qIdjTFINIr\n" +
		"evnAhf0cwULaebn+lMs8Pdl7y37+sfluVok=\n" +
		"-----END CERTIFICATE-----"

	certPKPem := "-----BEGIN RSA PRIVATE KEY-----\n" +
		"MIICWwIBAAKBgQDGAQa8eT5T9FjSP2Xcw/uj5LZrq5/hdpEGG7XO10IfPy0wbwqP\n" +
		"j+omaCxJlAPXuxFy0cYFNQlangIu0HAJ/TMAZXPJLBSRwdK7X/aZn/Ds2vRNAcp5\n" +
		"av+Pym9cfnfgMoS+6CvbUMAduLhzrnh4tQv4lb/AkliomHuoczdbcWHvKwIDAQAB\n" +
		"AoGAXzxrIwgmBHeIqUe5FOBnDsOZQlyAQA+pXYjCf8Rll2XptFwUdkzAUMzWUGWT\n" +
		"G5ZspA9l8Wc7IozRe/bhjMxuVK5yZhPDKbjqRdWICA95Jd7fxlIirHOVMQRdzI7x\n" +
		"NKqMNQN05MLJfsEHUYtOLhZE+tfhJTJnnmB7TMwnJgc4O5ECQQD8oOJ45tyr46zc\n" +
		"OAt6ao7PefVLiW5Qu+PxfoHmZmDV2UQqeM5XtZg4O97VBSugOs3+quIdAC6LotYl\n" +
		"/6N+E4y3AkEAyKWD2JNCrAgtjk2bfF1HYt24tq8+q7x2ek3/cUhqwInkrZqOFoke\n" +
		"x3+yBB879TuUOadvBXndgMHHcJQKSAJlLQJAXRuGnHyptAhTe06EnHeNbtZKG67p\n" +
		"I4Q8PJMdmSb+ZZKP1v9zPUxGb+NQ+z3OmF1T8ppUf8/DV9+KAbM4NI1L/QJAdGBs\n" +
		"BKYFObrUkYE5+fwwd4uao3sponqBTZcH3jDemiZg2MCYQUHu9E+AdRuYrziLVJVk\n" +
		"s4xniVLb1tRG0lVxUQJASfjdGT81HDJSzTseigrM+JnBKPPrzpeEp0RbTP52Lm23\n" +
		"YARjLCwmPMMdAwYZsvqeTuHEDQcOHxLHWuyN/zgP2A==\n" +
		"-----END RSA PRIVATE KEY-----"

	certID := testAccNSXCreateCert(t, "test", certPem, certPKPem, "certificate_signed")

	clientCertPem := "-----BEGIN CERTIFICATE-----\n" +
		"MIID6jCCAtKgAwIBAgIJAOtKKdMP6oZcMA0GCSqGSIb3DQEBCwUAMHYxCzAJBgNV\n" +
		"BAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBX\n" +
		"aWRnaXRzIFB0eSBMdGQxETAPBgNVBAMMCHRlc3QuY29tMRwwGgYJKoZIhvcNAQkB\n" +
		"Fg10ZXN0QHRlc3QuY29tMB4XDTE4MDUwNzE3MTkxOVoXDTE5MDUwNzE3MTkxOVow\n" +
		"djELMAkGA1UEBhMCQVUxEzARBgNVBAgMClNvbWUtU3RhdGUxITAfBgNVBAoMGElu\n" +
		"dGVybmV0IFdpZGdpdHMgUHR5IEx0ZDERMA8GA1UEAwwIdGVzdC5jb20xHDAaBgkq\n" +
		"hkiG9w0BCQEWDXRlc3RAdGVzdC5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw\n" +
		"ggEKAoIBAQC8GhgSX8nV0QEStgQ208PDXsl62xefAWVYMVjF9+xUWHY6zp/sTTQA\n" +
		"uFGD6DBiyljYYShj85ha7N3m70Xmur5Bbpa7ca010R0oCLPQRcySvsom190OfeUG\n" +
		"LW396gZUjiY9+8OQqI8BrTcNkI1DOpC5fBbbXkvLN4pOIf+/JqZNtr1igQBv2g+E\n" +
		"FCiPxDuWPOagi2xeEQv3SLLBhfF92UAlV6JRziCwaILlC3Zn/6B8E8rpuqVxUShW\n" +
		"wDmLowoT4BZP6/OHRneyEvVi2L/Ucsdk3nwUTe5Q4ojHY21ftW6by6uXYvK4IePC\n" +
		"ZXeTjdNtRX4MaAtJ4iLY7E/9d0BdTNRvAgMBAAGjezB5MAkGA1UdEwQCMAAwLAYJ\n" +
		"YIZIAYb4QgENBB8WHU9wZW5TU0wgR2VuZXJhdGVkIENlcnRpZmljYXRlMB0GA1Ud\n" +
		"DgQWBBTP56rycJh9hpWKOmvTzYg3YEkqiDAfBgNVHSMEGDAWgBTP56rycJh9hpWK\n" +
		"OmvTzYg3YEkqiDANBgkqhkiG9w0BAQsFAAOCAQEAZ90xoDf4gRbh7/PHxbokP69K\n" +
		"uqt7s78JDKrGxCDwiezUhZrdOBwi2r/sg4cWi43uuwCfjGNCd824EQYRaSCjanAn\n" +
		"5OH14KshCOBi66CaWDzJK6v4X/hbrKtUmXvbvUjrQCEHVuLueEQfvJB5/8O0dpEA\n" +
		"xF4MhnSaF2Id5O7tlXnFoXZh0YI9QnTwHLQ4L9+3PS5LqWd0peV1XqgWy0CXwjZF\n" +
		"nEpHq+TGDwLRoAgnoBGrbaFJRmvm+iVU4J76AtV7B3keckVMyMIeBR9CWB7kDm64\n" +
		"86qiRcEGN7V5mMJtxF49l0F01qdOgrictZRf+gMMrtGmX4KkZ6DKrl278HPs7A==\n" +
		"-----END CERTIFICATE-----"

	clientPKPem := "-----BEGIN PRIVATE KEY-----\n" +
		"MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC8GhgSX8nV0QES\n" +
		"tgQ208PDXsl62xefAWVYMVjF9+xUWHY6zp/sTTQAuFGD6DBiyljYYShj85ha7N3m\n" +
		"70Xmur5Bbpa7ca010R0oCLPQRcySvsom190OfeUGLW396gZUjiY9+8OQqI8BrTcN\n" +
		"kI1DOpC5fBbbXkvLN4pOIf+/JqZNtr1igQBv2g+EFCiPxDuWPOagi2xeEQv3SLLB\n" +
		"hfF92UAlV6JRziCwaILlC3Zn/6B8E8rpuqVxUShWwDmLowoT4BZP6/OHRneyEvVi\n" +
		"2L/Ucsdk3nwUTe5Q4ojHY21ftW6by6uXYvK4IePCZXeTjdNtRX4MaAtJ4iLY7E/9\n" +
		"d0BdTNRvAgMBAAECggEAYS8KKNQcr7/gUg6AduNKTXJ3nmX7+kb6WWqFdyL0k09x\n" +
		"JkkDD0+EAqs4UnJwLVpmNMVd3OZKXQ/sRhhxgRgSnDPK5OWCnD+CVODKJl0pqNey\n" +
		"EgeNSqN45IwsO/fhdWZME9Iz5FVyLWeU/gklMwrbIzodhRFfD4uOhXfDbrtFSPix\n" +
		"KGzhDNCn3z/Y9Sml6tUMgCLPisQbxWcQ+bbTHCgXXEdK2yG45j4q3FqBGy8IExou\n" +
		"cxw0aWu8cIqIGlXBQA5SZ78csBisKxEmlHjkV1i7Jn2DIARUluU/nmYKGC4ElWGp\n" +
		"yuURRJhcjO17MkCEr+niTxnTptkHi9e1p4d/k6HCgQKBgQD0Uj/JPr9vIM6Mt3vm\n" +
		"kK0GSTfpMUDDwaAMR+Pip6qocEYD5nZ6rnSuXIqlKyP/YxtQqsDqFdYHbPJet/Lg\n" +
		"3HFKerQ7xHiral95rsP0TWiRtxCHaF1xahgQWlVaYqGDeBcL7EavGJwQy3crFiVo\n" +
		"6wHwbVfNV/hoXHHK7pX4J8M3sQKBgQDFF+P/gM86QrvDbJnH/7PuF6QzcdvvtacH\n" +
		"i+5aTG9YQaUGldQIx2Fjpn48V73/YGK04DTrXqACjE2XxgEu0/B0wAGkOwGZztAr\n" +
		"q54O4AEWgVAyUukarkjdGtgeUAf/UMeXf7D2tGhjyNtORrOyZCdKS3kZ3GfZuEXj\n" +
		"FjhbJbX2HwKBgCMh8Kovq7d/MDRr7hUpmLfer3uI6Zc8sJcTf2GIWrH98xN8gG0D\n" +
		"ySOJiyZVHcgLqFHhO/xtR2mp8PBN408SY/ghzOkLR47erPwCdYsb1n2dpXLTPxyf\n" +
		"9PXlB4EHzdHp4uaEA2YKU+bWWzyG4rpDkPPRxV5x1/ap1HMp+8bDcP8BAoGAJTdQ\n" +
		"pxNUjgTB3bHpC9ndyOyP5eLvC8F6S7OBi215bOngVnD+O7YiTqXGmnBbARjbKppX\n" +
		"g8Y3YqPJlwodeREuC22iIbe+oqNprYVXcCmeKvi6Avai65XTTmTeQEMOb4h6V8IV\n" +
		"0U/ZklYACzTQg7Pjs2Sy9k4nEfZ4w9uTQqrJRDMCgYEAvgB29WCwEJ81pjBIGiYw\n" +
		"FbqLxQtIFiXKvRW+Qr1EKcV89kUYvsH3+TO/hZpf6rBtRf9yLJvw2fR4tVouowbn\n" +
		"/w+mKvKmcm11NOuobla4w9gztptqeUzcW03JmHmcnyHlnnUwpsU/XcgnsToV6kJB\n" +
		"bS6lN3HXkJnFnAzK2BKoZCA=\n" +
		"-----END PRIVATE KEY-----\n"

	clientCertID := testAccNSXCreateCert(t, "test_client", clientCertPem, clientPKPem, "certificate_self_signed")

	caCertPem := "-----BEGIN CERTIFICATE-----\n" +
		"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\n" +
		"-----END CERTIFICATE-----"

	caCertID := testAccNSXCreateCert(t, "test_ca", caCertPem, "", "certificate_ca")

	return certID, clientCertID, caCertID
}

func testAccNSXDeleteCerts(t *testing.T, certID string, clientCertID string, caCertID string) {

	testAccNSXDeleteCert(t, certID)
	testAccNSXDeleteCert(t, clientCertID)
	testAccNSXDeleteCert(t, caCertID)
}

func testAccNsxtPolicyEmptyTemplate() string {
	return " "
}
