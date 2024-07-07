package helpers

import (
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// To create test certificates:
// openssl req -subj '/CN=hashicorptest/O=HashiCorp, Inc./ST=CA/C=US' -new -newkey rsa:2048 -sha256 -days 3650 -nodes -x509 -keyout server.key -out server.crt
// grep -v \\----- server.crt >server.b64
// cat server.b64 | base64 -d | xxd -p

const applicationCertificatePem string = `-----BEGIN CERTIFICATE-----
MIIDFDCCAfwCCQCvHp+vopfOOTANBgkqhkiG9w0BAQsFADBMMRYwFAYDVQQDDA1o
YXNoaWNvcnB0ZXN0MRgwFgYDVQQKDA9IYXNoaUNvcnAsIEluYy4xCzAJBgNVBAgM
AkNBMQswCQYDVQQGEwJVUzAeFw0yMTAzMDkxMTAyMTNaFw0zMTAzMDcxMTAyMTNa
MEwxFjAUBgNVBAMMDWhhc2hpY29ycHRlc3QxGDAWBgNVBAoMD0hhc2hpQ29ycCwg
SW5jLjELMAkGA1UECAwCQ0ExCzAJBgNVBAYTAlVTMIIBIjANBgkqhkiG9w0BAQEF
AAOCAQ8AMIIBCgKCAQEAlVmb5pmoASvZ5pxD6CEBiPYqADb7teCHV54RRwv1aJjS
eiPUW/1WNQooIQF0M0yzFdHmwx3HSoxCkQwwxVMAPsuqFJVabs/eAr41NpxQCncb
i+vKlbmaAWbaIdidxeUe1jXB2N0YXRCg7Ps8IGA0UochvRGypfciy4k6/xEfrrQP
FlrPeDeaurNUjJ4IotTBLzWNAX9nT1HKzvljYNg4A0PwuzPNOmgxUSpAeiPbDoQo
D/YcQUKWzBlW8qt9ZnuRMGNi6V2fnQeTLblfsheaavXyP11syJ9owz6mDffZELHd
SYC7j2EOqG+Pndd55MLOac8cF4D9Y91PkLKKjNIrWwIDAQABMA0GCSqGSIb3DQEB
CwUAA4IBAQBlVJLn17BFmigbqS8JIx0/RTbGokRoLKdg7SZAQJWn20jDtunSo+sp
ZzuZ4uS8WbgZ+SFD1rrQy3s0F9HssZFBwDGyn31z/sGjkwWpoAP65v1DCaNzmAsz
xMNijhYlShv61g2IEO9Q98bgBW9LNwmJRGnGxz0ufzeZuUr9IV9EjeoJCKPIbwJC
lab0Ty/kRC13JgNhHtNFwYVwK6NDt46IRsjxqWQ6bVakrEROlfuoY8sxUjunj+hB
2vZTkZKaPc0sFvUQjNHxHX4jMeTwCopQCo+qF3lPde+G7C1MNf30kDZlks++GLNs
0/0Ayfjh6JllWqW482dIIqMErl6s5DuK
-----END CERTIFICATE-----`

func getKeyCredentialSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"value": {
			Type:     schema.TypeString,
			Required: true,
		},
		"encoding": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"start_date": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"end_date": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"end_date_relative": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func TestKeyCredentialForResource_withEndDate(t *testing.T) {

	// Arrange
	d := schema.TestResourceDataRaw(t, getKeyCredentialSchema(), nil)
	d.Set("type", "AsymmetricX509Cert")
	d.Set("value", applicationCertificatePem)
	d.Set("encoding", "pem")

	now := time.Now()
	d.Set("end_date", now.Format(time.RFC3339))

	// Act
	keyCredential, err := KeyCredentialForResource(d)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	actual := keyCredential.EndDateTime
	expected := now

	// In order to compare the time.Time objects, we need to truncate the time.
	// 3 minutes is a reasonable amount of time to allow for the time it takes to run the test.
	trunc := 3 * time.Minute
	if actual.Truncate(trunc).Equal(expected.Truncate(trunc)) == false {
		t.Fatalf("expected end date to be %v, got %v", expected, actual)
	}
}

func TestKeyCredentialForResource_withEndDateRelative(t *testing.T) {

	// Arrange
	d := schema.TestResourceDataRaw(t, getKeyCredentialSchema(), nil)
	d.Set("type", "AsymmetricX509Cert")
	d.Set("value", applicationCertificatePem)
	d.Set("encoding", "pem")

	d.Set("end_date_relative", "240h") // 10 days

	// Act
	keyCredential, err := KeyCredentialForResource(d)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	actual := keyCredential.EndDateTime
	expected := time.Now().Add(time.Hour * 240) // 10 days

	// In order to compare the time.Time objects, we need to truncate the time.
	// 3 minutes is a reasonable amount of time to allow for the time it takes to run the test.
	trunc := 3 * time.Minute
	if actual.Truncate(trunc).Equal(expected.Truncate(trunc)) == false {
		t.Fatalf("expected end date to be 10 days from now, got %v", actual)
	}
}

func TestKeyCredentialForResource_withStartDateAndEndDateRelative(t *testing.T) {

	// Arrange
	d := schema.TestResourceDataRaw(t, getKeyCredentialSchema(), nil)
	d.Set("type", "AsymmetricX509Cert")
	d.Set("value", applicationCertificatePem)
	d.Set("encoding", "pem")

	start_date := time.Now().Add(time.Hour * 984)
	d.Set("start_date", start_date.Format(time.RFC3339)) // 41 days
	d.Set("end_date_relative", "240h")                   // 10 days

	// Act
	keyCredential, err := KeyCredentialForResource(d)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	actual := keyCredential.EndDateTime
	expected := start_date.Add(time.Hour * 240) // 10 days

	// In order to compare the time.Time objects, we need to truncate the time.
	// 3 minutes is a reasonable amount of time to allow for the time it takes to run the test.
	trunc := 3 * time.Minute
	if actual.Truncate(trunc).Equal(expected.Truncate(trunc)) == false {
		t.Fatalf("expected end date to be 10 days from %v, got %v", start_date, actual)
	}

}
