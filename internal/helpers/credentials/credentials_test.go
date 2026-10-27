// Copyright IBM Corp. 2019, 2026
// SPDX-License-Identifier: MPL-2.0

package credentials

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func keyCredentialTestSchema() map[string]*schema.Schema {
	schemaWithDescription := func(description string) *schema.Schema {
		return &schema.Schema{Description: description, Type: schema.TypeString, Optional: true}
	}

	return map[string]*schema.Schema{
		"type":              schemaWithDescription("Type"),
		"value":             schemaWithDescription("Value"),
		"encoding":          schemaWithDescription("Encoding"),
		"key_id":            schemaWithDescription("Key id"),
		"start_date":        schemaWithDescription("Start date"),
		"end_date":          schemaWithDescription("End date"),
		"end_date_relative": schemaWithDescription("End date relative"),
	}
}

func TestKeyCredentialForResource_endDate(t *testing.T) {
	cases := []struct {
		Name            string
		Config          map[string]interface{}
		ExpectedEndDate string
	}{
		{
			Name: "end_date",
			Config: map[string]interface{}{
				"end_date": "2030-01-01T00:00:00Z",
			},
			ExpectedEndDate: "2030-01-01T00:00:00Z",
		},
		{
			Name: "start_date_and_end_date_relative",
			Config: map[string]interface{}{
				"start_date":        "2030-01-01T00:00:00Z",
				"end_date_relative": "240h",
			},
			ExpectedEndDate: "2030-01-11T00:00:00Z",
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.Config["type"] = "AsymmetricX509Cert"
			tc.Config["encoding"] = "pem"
			tc.Config["value"] = "test"
			tc.Config["key_id"] = "00000000-0000-0000-0000-000000000000"

			d := schema.TestResourceDataRaw(t, keyCredentialTestSchema(), tc.Config)

			credential, err := KeyCredentialForResource(d)
			if err != nil {
				t.Fatalf("unexpected error: %+v", err)
			}

			if actual := credential.EndDateTime.GetOrZero(); actual != tc.ExpectedEndDate {
				t.Fatalf("expected end date %q, got %q", tc.ExpectedEndDate, actual)
			}
		})
	}
}
