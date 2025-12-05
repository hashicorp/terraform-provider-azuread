// Copyright IBM Corp. 2019, 2025
// SPDX-License-Identifier: MPL-2.0

package credentials

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

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
)

func GetKeyCredential(keyCredentials *[]stable.KeyCredential, id string) (credential *stable.KeyCredential) {
	if keyCredentials != nil {
		for _, cred := range *keyCredentials {
			if strings.EqualFold(cred.KeyId.GetOrZero(), id) {
				credential = &cred
				break
			}
		}
	}
	return
}

func GetVerifyKeyCredentialFromCustomKeyId(keyCredentials *[]stable.KeyCredential, id string) (credential *stable.KeyCredential) {
	if keyCredentials != nil {
		for _, cred := range *keyCredentials {
			if !cred.KeyId.IsNull() && strings.EqualFold(cred.CustomKeyIdentifier.GetOrZero(), id) && strings.EqualFold(cred.Usage.GetOrZero(), KeyCredentialUsageVerify) {
				credential = &cred
				break
			}
		}
	}
	return
}

func GetPasswordCredential(passwordCredentials *[]stable.PasswordCredential, id string) (credential *stable.PasswordCredential) {
	if passwordCredentials != nil {
		for _, cred := range *passwordCredentials {
			if strings.EqualFold(cred.KeyId.GetOrZero(), id) {
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

func KeyCredentialForResource(d *pluginsdk.ResourceData) (*stable.KeyCredential, error) {
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

	credential := stable.KeyCredential{
		KeyId: nullable.Value(keyId),
		Type:  nullable.Value(keyType),
		Usage: nullable.Value(KeyCredentialUsageVerify),
		Key:   nullable.Value(encodedValue),
	}

	if v, ok := d.GetOk("start_date"); ok {
		startDate, err := time.Parse(time.RFC3339, v.(string))
		if err != nil {
			return nil, CredentialError{str: fmt.Sprintf("Unable to parse the provided start date %q: %+v", v, err), attr: "start_date"}
		}
		credential.StartDateTime = nullable.Value(startDate.Format(time.RFC3339))
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
			startDateTime, err := time.Parse(time.RFC3339, v.(string))
			if err != nil {
				return nil, CredentialError{str: fmt.Sprintf("Unable to parse the provided start date %q: %+v", v, err), attr: "start_date"}
			}
			expiry := startDateTime.Add(d)
			endDate = &expiry
		}
	}

	if endDate != nil {
		credential.EndDateTime = nullable.Value(endDate.Format(time.RFC3339))
	}

	return &credential, nil
}

func PasswordCredential(in map[string]interface{}) (*stable.PasswordCredential, error) {
	credential := stable.PasswordCredential{}

	if v, ok := in["display_name"]; ok {
		credential.DisplayName = nullable.Value(v.(string))
	}

	if v, ok := in["start_date"]; ok && v.(string) != "" {
		startDate, err := time.Parse(time.RFC3339, v.(string))
		if err != nil {
			return nil, CredentialError{str: fmt.Sprintf("Unable to parse the provided start date %q: %+v", v, err), attr: "start_date"}
		}
		credential.StartDateTime = nullable.Value(startDate.Format(time.RFC3339))
	}

	if v, ok := in["end_date"]; ok && v.(string) != "" {
		var err error
		expiry, err := time.Parse(time.RFC3339, v.(string))
		if err != nil {
			return nil, CredentialError{str: fmt.Sprintf("Unable to parse the provided end date %q: %+v", v, err), attr: "end_date"}
		}

		credential.EndDateTime = nullable.Value(expiry.Format(time.RFC3339))
	}

	if v, ok := in["key_id"]; ok && v.(string) != "" {
		credential.KeyId = nullable.Value(v.(string))
	}

	if v, ok := in["value"]; ok && v.(string) != "" {
		credential.SecretText = nullable.Value(v.(string))
	}

	return &credential, nil
}

func PasswordCredentialForResource(d *pluginsdk.ResourceData) (*stable.PasswordCredential, error) {
	data := make(map[string]interface{})

	// display_name, start_date and end_date support intentionally remains for if/when the API supports user-specified values for these
	if v, ok := d.GetOk("display_name"); ok {
		data["display_name"] = v
	}

	if v, ok := d.GetOk("start_date"); ok {
		data["start_date"] = v
	}

	if v, ok := d.GetOk("end_date"); ok && v.(string) != "" {
		data["end_date"] = v
	} else if v, ok := d.GetOk("end_date_relative"); ok && v.(string) != "" {
		data["end_date_relative"] = v
	}

	return PasswordCredential(data)
}
