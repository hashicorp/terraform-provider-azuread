package serviceprincipals

import (
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

func expandSamlSingleSignOn(in []interface{}) *msgraph.SamlSingleSignOnSettings {
	result := msgraph.SamlSingleSignOnSettings{}
	if len(in) == 0 || in[0] == nil {
		return &result
	}

	samlSingleSignOnSettings := in[0].(map[string]interface{})

	result.RelayState = utils.String(samlSingleSignOnSettings["relay_state"].(string))

	return &result
}

func flattenSamlSingleSignOn(in *msgraph.SamlSingleSignOnSettings) []map[string]interface{} {
	if in == nil {
		return []map[string]interface{}{}
	}

	relayState := ""
	if in.RelayState != nil {
		relayState = *in.RelayState
	}

	return []map[string]interface{}{{
		"relay_state": relayState,
	}}
}
