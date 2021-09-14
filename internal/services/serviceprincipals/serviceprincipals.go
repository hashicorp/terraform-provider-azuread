package serviceprincipals

import (
	"strings"

	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/manicminer/hamilton/msgraph"
)

func expandFeatures(in []interface{}) (out []string) {
	if len(in) == 0 || in[0] == nil {
		return
	}

	features := in[0].(map[string]interface{})

	if v, ok := features["custom_single_sign_on_app"]; ok && v.(bool) {
		out = append(out, "WindowsAzureActiveDirectoryCustomSingleSignOnApplication")
	}

	if v, ok := features["enterprise_application"]; ok && v.(bool) {
		out = append(out, "WindowsAzureActiveDirectoryIntegratedApp")
	}

	if v, ok := features["gallery_application"]; ok && v.(bool) {
		out = append(out, "WindowsAzureActiveDirectoryGalleryApplicationNonPrimaryV1")
	}

	if v, ok := features["visible_to_users"]; ok && !v.(bool) {
		out = append(out, "HideApp")
	}

	return
}

func expandSamlSingleSignOn(in []interface{}) *msgraph.SamlSingleSignOnSettings {
	result := msgraph.SamlSingleSignOnSettings{}
	if len(in) == 0 || in[0] == nil {
		return &result
	}

	samlSingleSignOnSettings := in[0].(map[string]interface{})

	result.RelayState = utils.String(samlSingleSignOnSettings["relay_state"].(string))

	return &result
}

func flattenFeatures(tags *[]string) []interface{} {
	result := map[string]bool{
		"custom_single_sign_on_app": false,
		"enterprise_application":    false,
		"gallery_application":       false,
		"visible_to_users":          true,
	}

	if tags == nil || len(*tags) == 0 {
		return []interface{}{result}
	}

	for _, tag := range *tags {
		if strings.EqualFold(tag, "WindowsAzureActiveDirectoryCustomSingleSignOnApplication") {
			result["custom_single_sign_on_app"] = true
		}
		if strings.EqualFold(tag, "WindowsAzureActiveDirectoryIntegratedApp") {
			result["enterprise_application"] = true
		}
		if strings.EqualFold(tag, "WindowsAzureActiveDirectoryGalleryApplicationNonPrimaryV1") {
			result["gallery_application"] = true
		}
		if strings.EqualFold(tag, "HideApp") {
			result["visible_to_users"] = false
		}
	}

	return []interface{}{result}
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
