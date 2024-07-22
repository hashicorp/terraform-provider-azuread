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

const (
	applicationCertificatePem string = `-----BEGIN CERTIFICATE-----
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

	// In order to compare the time.Time objects, we need to truncate the time.
	truncateTime time.Duration = time.Hour
)

func getNowTime() time.Time {
	return time.Now().Truncate(truncateTime)
}

func getSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"display_name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"key_id": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"type": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"value": {
			Type:     schema.TypeString,
			Optional: true,
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

func TestKeyCredentialForResource(t *testing.T) {

	// Arrange
	expectedType := "AsymmetricX509Cert"
	expectedKeyId := "test_key_id"
	expectStartDate := getNowTime()
	expectedEndDate := expectStartDate.Add(time.Hour * 984) // 41 days from now

	d := schema.TestResourceDataRaw(t, getSchema(), nil)
	d.Set("type", expectedType)
	d.Set("key_id", expectedKeyId)
	d.Set("value", applicationCertificatePem)
	d.Set("encoding", "pem")
	d.Set("start_date", expectStartDate.Format(time.RFC3339))
	d.Set("end_date", expectedEndDate.Format(time.RFC3339))

	// Act
	keyCredential, err := KeyCredentialForResource(d)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if keyCredential.Type != expectedType {
		t.Fatalf("expected type to be '%v', got %v", expectedType, keyCredential.Type)
	}

	if keyCredential.KeyId == nil || *keyCredential.KeyId != expectedKeyId {
		t.Fatalf("expected key id to be '%v', got %v", expectedKeyId, keyCredential.KeyId)
	}

	if keyCredential.Key == nil {
		t.Fatal("expected key to be set, got nil")
	}

	if keyCredential.StartDateTime == nil || keyCredential.StartDateTime.Equal(expectStartDate) == false {
		t.Fatalf("expected start date to be %v, got %v", expectStartDate, keyCredential.StartDateTime)
	}

	if keyCredential.EndDateTime == nil || keyCredential.EndDateTime.Equal(expectedEndDate) == false {
		t.Fatalf("expected end date to be %v, got %v", expectedEndDate, keyCredential.EndDateTime)
	}
}

func TestKeyCredentialForResource_withEndDateRelative(t *testing.T) {

	// Arrange
	d := schema.TestResourceDataRaw(t, getSchema(), nil)
	d.Set("type", "AsymmetricX509Cert")
	d.Set("value", applicationCertificatePem)
	d.Set("encoding", "pem")

	expectedEndDate := getNowTime().Add(time.Hour * 240) // 10 days
	d.Set("end_date_relative", "240h")                   // 10 days

	// Act
	keyCredential, err := KeyCredentialForResource(d)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if keyCredential.EndDateTime.Truncate(truncateTime).Equal(expectedEndDate) == false {
		t.Fatalf("expected end date to be %v got %v", expectedEndDate, keyCredential.EndDateTime)
	}
}

func TestKeyCredentialForResource_withStartDateAndEndDateRelative(t *testing.T) {

	// Arrange
	d := schema.TestResourceDataRaw(t, getSchema(), nil)
	d.Set("type", "AsymmetricX509Cert")
	d.Set("value", applicationCertificatePem)
	d.Set("encoding", "pem")

	start_date := getNowTime().Add(time.Hour * 984)
	expectedEndDate := start_date.Add(time.Hour * 240) // 10 days
	d.Set("start_date", start_date.Format(time.RFC3339))
	d.Set("end_date_relative", "240h") // 10 days

	// Act
	keyCredential, err := KeyCredentialForResource(d)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if keyCredential.EndDateTime.Equal(expectedEndDate) == false {
		t.Fatalf("expected end date to be %v got %v", expectedEndDate, keyCredential.EndDateTime)
	}
}

func TestPasswordCredentialForResource(t *testing.T) {

	// Arrange
	expectedDisplayName := "test_display_name"
	expectStartDate := getNowTime()
	expectedEndDate := expectStartDate.Add(time.Hour * 984) // 41 days from now

	d := schema.TestResourceDataRaw(t, getSchema(), nil)
	d.Set("display_name", expectedDisplayName)
	d.Set("start_date", expectStartDate.Format(time.RFC3339))
	d.Set("end_date", expectedEndDate.Format(time.RFC3339))

	// Act
	passwordCredential, err := PasswordCredentialForResource(d)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if passwordCredential.DisplayName == nil || *passwordCredential.DisplayName != expectedDisplayName {
		t.Fatalf("expected display name to be '%v', got %v", expectedDisplayName, passwordCredential.DisplayName)
	}

	if passwordCredential.StartDateTime == nil || passwordCredential.StartDateTime.Equal(expectStartDate) == false {
		t.Fatalf("expected start date to be %v, got %v", expectStartDate, passwordCredential.StartDateTime)
	}

	if passwordCredential.EndDateTime == nil || passwordCredential.EndDateTime.Equal(expectedEndDate) == false {
		t.Fatalf("expected end date to be %v, got %v", expectedEndDate, passwordCredential.EndDateTime)
	}
}

func TestPasswordCredentialForResource_withEndDateRelative(t *testing.T) {

	// Arrange
	d := schema.TestResourceDataRaw(t, getSchema(), nil)
	d.Set("display_name", "test_display_name")

	expectedEndDate := getNowTime().Add(time.Hour * 240) // 10 days
	d.Set("end_date_relative", "240h")                   // 10 days

	// Act
	passwordCredential, err := PasswordCredentialForResource(d)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// In order to compare the time.Time objects, we need to truncate the time.
	if passwordCredential.EndDateTime.Truncate(truncateTime).Equal(expectedEndDate) == false {
		t.Fatalf("expected end date to be %v got %v", expectedEndDate, passwordCredential.EndDateTime)
	}
}

func TestPasswordCredentialForResource_withStartDateAndEndDateRelative(t *testing.T) {

	// Arrange
	d := schema.TestResourceDataRaw(t, getSchema(), nil)
	d.Set("display_name", "test_display_name")

	start_date := getNowTime().Add(time.Hour * 984)
	expectedEndDate := start_date.Add(time.Hour * 240) // 10 days
	d.Set("start_date", start_date.Format(time.RFC3339))
	d.Set("end_date_relative", "240h") // 10 days

	// Act
	passwordCredential, err := PasswordCredentialForResource(d)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if passwordCredential.EndDateTime.Equal(expectedEndDate) == false {
		t.Fatalf("expected end date to be %v, got %v", expectedEndDate, passwordCredential.EndDateTime)
	}
}
