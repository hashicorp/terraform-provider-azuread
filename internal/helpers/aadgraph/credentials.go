package aadgraph

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

func PasswordCredentialForResource(d *schema.ResourceData) (*graphrbac.PasswordCredential, error) {
	value := d.Get("value").(string)

	// errors should be handled by the validation
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

	var endDate time.Time
	if v := d.Get("end_date").(string); v != "" {
		var err error
		endDate, err = time.Parse(time.RFC3339, v)
		if err != nil {
			return nil, CredentialError{str: fmt.Sprintf("Unable to parse the provided end date %q: %+v", v, err), attr: "end_date"}
		}
	} else if v := d.Get("end_date_relative").(string); v != "" {
		d, err := time.ParseDuration(v)
		if err != nil {
			return nil, CredentialError{str: fmt.Sprintf("Unable to parse `end_date_relative` (%q) as a duration", v), attr: "end_date_relative"}
		}
		endDate = time.Now().Add(d)
	} else {
		return nil, CredentialError{str: "One of `end_date` or `end_date_relative` must be specified", attr: "end_date"}
	}

	credential := graphrbac.PasswordCredential{
		KeyID:   utils.String(keyId),
		Value:   utils.String(value),
		EndDate: &date.Time{Time: endDate},
	}

	if v, ok := d.GetOk("display_name"); ok {
		customIdentifier := []byte(v.(string))
		credential.CustomKeyIdentifier = &customIdentifier
	} else if v, ok := d.GetOk("description"); ok {
		customIdentifier := []byte(v.(string))
		credential.CustomKeyIdentifier = &customIdentifier
	}

	if v, ok := d.GetOk("start_date"); ok {
		startDate, err := time.Parse(time.RFC3339, v.(string))
		if err != nil {
			return nil, CredentialError{str: fmt.Sprintf("Unable to parse the provided start date %q: %+v", v, err), attr: "start_date"}
		}
		credential.StartDate = &date.Time{Time: startDate}
	}

	return &credential, nil
}

func PasswordCredentialResultFindByKeyId(creds graphrbac.PasswordCredentialListResult, keyId string) *graphrbac.PasswordCredential {
	var cred *graphrbac.PasswordCredential

	if creds.Value != nil {
		for _, c := range *creds.Value {
			if c.KeyID == nil {
				continue
			}

			if *c.KeyID == keyId {
				cred = &c
				break
			}
		}
	}

	return cred
}

func PasswordCredentialResultAdd(existing graphrbac.PasswordCredentialListResult, cred *graphrbac.PasswordCredential) (*[]graphrbac.PasswordCredential, error) {
	if cred == nil {
		return nil, fmt.Errorf("credential to be added is nil")
	}

	newCreds := make([]graphrbac.PasswordCredential, 0)

	if existing.Value != nil {
		for _, v := range *existing.Value {
			if v.KeyID == nil {
				continue
			}
			if *v.KeyID == *cred.KeyID {
				return nil, &AlreadyExistsError{"Password Credential", *cred.KeyID}
			}
		}

		newCreds = *existing.Value
	}
	newCreds = append(newCreds, *cred)

	return &newCreds, nil
}

func PasswordCredentialResultRemoveByKeyId(existing graphrbac.PasswordCredentialListResult, keyId string) (*[]graphrbac.PasswordCredential, error) {
	if keyId == "" {
		return nil, fmt.Errorf("ID of key to be removed is empty")
	}

	newCreds := make([]graphrbac.PasswordCredential, 0)

	if existing.Value != nil {
		for _, v := range *existing.Value {
			if v.KeyID == nil {
				continue
			}

			if *v.KeyID == keyId {
				continue
			}

			newCreds = append(newCreds, v)
		}
	}

	return &newCreds, nil
}

func WaitForPasswordCredentialReplication(ctx context.Context, keyId string, timeout time.Duration, f func() (graphrbac.PasswordCredentialListResult, error)) (interface{}, error) {
	return (&resource.StateChangeConf{
		Pending:                   []string{"NotFound"},
		Target:                    []string{"Found"},
		Timeout:                   timeout,
		MinTimeout:                1 * time.Second,
		ContinuousTargetOccurence: 10,
		Refresh: func() (interface{}, string, error) {
			creds, err := f()
			if err != nil {
				if utils.ResponseWasNotFound(creds.Response) {
					return creds, "NotFound", nil
				}
				return creds, "Error", fmt.Errorf("unable to retrieve object, received response with status %d: %v", creds.Response.StatusCode, err)
			}

			credential := PasswordCredentialResultFindByKeyId(creds, keyId)
			if credential == nil {
				return creds, "NotFound", nil
			}

			return creds, "Found", nil
		},
	}).WaitForStateContext(ctx)
}

func KeyCredentialForResource(d *schema.ResourceData) (*graphrbac.KeyCredential, error) {
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

	// errors should be handled by the validation
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

	var endDate time.Time
	if v := d.Get("end_date").(string); v != "" {
		var err error
		endDate, err = time.Parse(time.RFC3339, v)
		if err != nil {
			return nil, CredentialError{str: fmt.Sprintf("Unable to parse the provided end date %q: %+v", v, err), attr: "end_date"}
		}
	} else if v := d.Get("end_date_relative").(string); v != "" {
		d, err := time.ParseDuration(v)
		if err != nil {
			return nil, CredentialError{str: fmt.Sprintf("Unable to parse `end_date_relative` (%q) as a duration", v), attr: "end_date_relative"}
		}
		endDate = time.Now().Add(d)
	} else {
		return nil, CredentialError{str: "One of `end_date` or `end_date_relative` must be specified", attr: "end_date"}
	}

	credential := graphrbac.KeyCredential{
		KeyID:   utils.String(keyId),
		Type:    utils.String(keyType),
		Usage:   utils.String("verify"),
		Value:   utils.String(encodedValue),
		EndDate: &date.Time{Time: endDate},
	}

	if v, ok := d.GetOk("start_date"); ok {
		startDate, err := time.Parse(time.RFC3339, v.(string))
		if err != nil {
			return nil, CredentialError{str: fmt.Sprintf("Unable to parse the provided start date %q: %+v", v, err), attr: "start_date"}
		}
		credential.StartDate = &date.Time{Time: startDate}
	}

	return &credential, nil
}

func KeyCredentialResultFindByKeyId(creds graphrbac.KeyCredentialListResult, keyId string) *graphrbac.KeyCredential {
	if creds.Value != nil {
		for _, c := range *creds.Value {
			if c.KeyID == nil {
				continue
			}
			if *c.KeyID == keyId {
				return &c
			}
		}
	}

	return nil
}

func KeyCredentialResultAdd(existing graphrbac.KeyCredentialListResult, cred *graphrbac.KeyCredential) (*[]graphrbac.KeyCredential, error) {
	newCreds := make([]graphrbac.KeyCredential, 0)

	if existing.Value != nil {
		for _, v := range *existing.Value {
			if v.KeyID == nil {
				continue
			}

			if *v.KeyID == *cred.KeyID {
				return nil, &AlreadyExistsError{"Key Credential", *cred.KeyID}
			}
		}

		newCreds = *existing.Value
	}
	newCreds = append(newCreds, *cred)

	return &newCreds, nil
}

func KeyCredentialResultRemoveByKeyId(existing graphrbac.KeyCredentialListResult, keyId string) (*[]graphrbac.KeyCredential, error) {
	if keyId == "" {
		return nil, fmt.Errorf("ID of key to be removed is empty")
	}

	newCreds := make([]graphrbac.KeyCredential, 0)

	if existing.Value != nil {
		for _, v := range *existing.Value {
			if v.KeyID == nil {
				continue
			}

			if *v.KeyID == keyId {
				continue
			}

			newCreds = append(newCreds, v)
		}
	}

	return &newCreds, nil
}

func WaitForKeyCredentialReplication(ctx context.Context, keyId string, timeout time.Duration, f func() (graphrbac.KeyCredentialListResult, error)) (interface{}, error) {
	return (&resource.StateChangeConf{
		Pending:                   []string{"NotFound"},
		Target:                    []string{"Found"},
		Timeout:                   timeout,
		MinTimeout:                1 * time.Second,
		ContinuousTargetOccurence: 10,
		Refresh: func() (interface{}, string, error) {
			creds, err := f()
			if err != nil {
				if utils.ResponseWasNotFound(creds.Response) {
					return creds, "NotFound", nil
				}
				return creds, "Error", fmt.Errorf("unable to retrieve object, received response with status %d: %v", creds.Response.StatusCode, err)
			}

			credential := KeyCredentialResultFindByKeyId(creds, keyId)
			if credential == nil {
				return creds, "NotFound", nil
			}

			return creds, "Found", nil
		},
	}).WaitForStateContext(ctx)
}
