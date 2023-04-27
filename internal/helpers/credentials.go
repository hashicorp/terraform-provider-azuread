package helpers

import (
	"bytes"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/manicminer/hamilton/msgraph"
)

type CredentialError struct {
	str  string
	attr string
}

func (e CredentialError) Attr() string {
	return e.attr
}

func (e CredentialError) Error() string {
	return e.str
}

func GetKeyCredential(keyCredentials *[]msgraph.KeyCredential, id string) (credential *msgraph.KeyCredential) {
	if keyCredentials != nil {
		for _, cred := range *keyCredentials {
			if cred.KeyId != nil && strings.EqualFold(*cred.KeyId, id) {
				credential = &cred
				break
			}
		}
	}
	return
}

func GetVerifyKeyCredentialFromCustomKeyId(keyCredentials *[]msgraph.KeyCredential, id string) (credential *msgraph.KeyCredential) {
	if keyCredentials != nil {
		for _, cred := range *keyCredentials {
			if cred.KeyId != nil && strings.EqualFold(*cred.CustomKeyIdentifier, id) && strings.EqualFold(cred.Usage, msgraph.KeyCredentialUsageVerify) {
				credential = &cred
				break
			}
		}
	}
	return
}

func GetPasswordCredential(passwordCredentials *[]msgraph.PasswordCredential, id string) (credential *msgraph.PasswordCredential) {
	if passwordCredentials != nil {
		for _, cred := range *passwordCredentials {
			if cred.KeyId != nil && strings.EqualFold(*cred.KeyId, id) {
				credential = &cred
				break
			}
		}
	}
	return
}

func GetTokenSigningCertificateThumbprint(certByte []byte) (string, error) {
	block, _ := pem.Decode(certByte)
	if block == nil {
		return "", fmt.Errorf("decoding certificate block")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("parsing certificate block data: %+v", err)
	}
	thumbprint := sha1.Sum(cert.Raw)

	var buf bytes.Buffer
	for _, f := range thumbprint {
		fmt.Fprintf(&buf, "%02X", f)
	}
	return buf.String(), nil
}

func KeyCredentialForResource(d *schema.ResourceData) (*msgraph.KeyCredential, error) {
	keyType := d.Get("type").(string)
	value := d.Get("value").(string)

	var encodedValue string
	encoding := d.Get("encoding").(string)
	switch encoding {
	case "base64":
		der, err := base64.StdEncoding.DecodeString(strings.TrimSpace(value))
		if err != nil {
			return nil, fmt.Errorf("failed to decode base64 certificate data")
		}
		block := pem.Block{
			Type:  "CERTIFICATE",
			Bytes: der,
		}
		pemVal := pem.EncodeToMemory(&block)
		if pemVal == nil {
			return nil, fmt.Errorf("failed to PEM-encode certificate")
		}
		encodedValue = base64.StdEncoding.EncodeToString(pemVal)
	case "hex":
		bytesVal := []byte(strings.TrimSpace(value))
		der := make([]byte, hex.DecodedLen(len(bytesVal)))
		_, err := hex.Decode(der, bytesVal)
		if err != nil {
			return nil, fmt.Errorf("failed to decode hexadecimal certificate data: %+v", err)
		}
		block := pem.Block{
			Type:  "CERTIFICATE",
			Bytes: der,
		}
		pemVal := pem.EncodeToMemory(&block)
		if pemVal == nil {
			return nil, fmt.Errorf("failed to PEM-encode certificate")
		}
		encodedValue = base64.StdEncoding.EncodeToString(pemVal)
	case "pem":
		encodedValue = base64.StdEncoding.EncodeToString([]byte(value))
	}

	var keyId string
	if v, ok := d.GetOk("key_id"); ok {
		keyId = v.(string)
	} else {
		kid, err := uuid.GenerateUUID()
		if err != nil {
			return nil, err
		}

		keyId = kid
	}

	credential := msgraph.KeyCredential{
		KeyId: utils.String(keyId),
		Type:  keyType,
		Usage: msgraph.KeyCredentialUsageVerify,
		Key:   utils.String(encodedValue),
	}

	if v, ok := d.GetOk("start_date"); ok {
		startDate, err := time.Parse(time.RFC3339, v.(string))
		if err != nil {
			return nil, CredentialError{str: fmt.Sprintf("Unable to parse the provided start date %q: %+v", v, err), attr: "start_date"}
		}
		credential.StartDateTime = &startDate
	}

	var endDate *time.Time
	if v, ok := d.GetOk("end_date"); ok && v.(string) != "" {
		var err error
		expiry, err := time.Parse(time.RFC3339, v.(string))
		if err != nil {
			return nil, CredentialError{str: fmt.Sprintf("Unable to parse the provided end date %q: %+v", v, err), attr: "end_date"}
		}
		endDate = &expiry
	} else if v, ok := d.GetOk("end_date_relative"); ok && v.(string) != "" {
		d, err := time.ParseDuration(v.(string))
		if err != nil {
			return nil, CredentialError{str: fmt.Sprintf("Unable to parse `end_date_relative` (%q) as a duration", v), attr: "end_date_relative"}
		}

		if credential.StartDateTime == nil {
			expiry := time.Now().Add(d)
			endDate = &expiry
		} else {
			expiry := credential.StartDateTime.Add(d)
			endDate = &expiry
		}
	}

	if endDate != nil {
		credential.EndDateTime = endDate
	}

	return &credential, nil
}

func PasswordCredentialForResource(d *schema.ResourceData) (*msgraph.PasswordCredential, error) {
	credential := msgraph.PasswordCredential{}

	// display_name, start_date and end_date support intentionally remains for if/when the API supports user-specified values for these
	if v, ok := d.GetOk("display_name"); ok {
		credential.DisplayName = utils.String(v.(string))
	}

	if v, ok := d.GetOk("start_date"); ok {
		startDate, err := time.Parse(time.RFC3339, v.(string))
		if err != nil {
			return nil, CredentialError{str: fmt.Sprintf("Unable to parse the provided start date %q: %+v", v, err), attr: "start_date"}
		}
		credential.StartDateTime = &startDate
	}

	var endDate *time.Time
	if v, ok := d.GetOk("end_date"); ok && v.(string) != "" {
		var err error
		expiry, err := time.Parse(time.RFC3339, v.(string))
		if err != nil {
			return nil, CredentialError{str: fmt.Sprintf("Unable to parse the provided end date %q: %+v", v, err), attr: "end_date"}
		}
		endDate = &expiry
	} else if v, ok := d.GetOk("end_date_relative"); ok && v.(string) != "" {
		d, err := time.ParseDuration(v.(string))
		if err != nil {
			return nil, CredentialError{str: fmt.Sprintf("Unable to parse `end_date_relative` (%q) as a duration", v), attr: "end_date_relative"}
		}

		if credential.StartDateTime == nil {
			expiry := time.Now().Add(d)
			endDate = &expiry
		} else {
			expiry := credential.StartDateTime.Add(d)
			endDate = &expiry
		}
	}
	if endDate != nil {
		credential.EndDateTime = endDate
	}

	return &credential, nil
}
