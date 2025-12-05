// Copyright IBM Corp. 2019, 2025
// SPDX-License-Identifier: MPL-2.0

package serviceprincipals

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

func expandSamlSingleSignOn(in []interface{}) *stable.SamlSingleSignOnSettings {
	result := stable.SamlSingleSignOnSettings{}
	if len(in) == 0 || in[0] == nil {
		return &result
	}

	samlSingleSignOnSettings := in[0].(map[string]interface{})

	result.RelayState = nullable.Value(samlSingleSignOnSettings["relay_state"].(string))

	return &result
}

func flattenSamlSingleSignOn(in *stable.SamlSingleSignOnSettings) []map[string]interface{} {
	if in == nil {
		return []map[string]interface{}{}
	}

	return []map[string]interface{}{{
		"relay_state": in.RelayState.GetOrZero(),
	}}
}
