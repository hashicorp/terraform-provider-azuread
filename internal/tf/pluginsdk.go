// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tf

// PluginSdkUnknownValue is a dummy value used/sent by the plugin SDK when a real value is not known at plan time,
// e.g. during a CustomizeDiff function
// See https://github.com/hashicorp/terraform-plugin-sdk/blob/main/internal/configs/hcl2shim/values.go#L16
const PluginSdkUnknownValue = "74D93920-ED26-11E3-AC10-0800200C9A66"

// ValueIsNotEmptyOrUnknown returns false if provided a blank string or a string that looks "unknown". Intended for use
// in CustomizeDiff functions to avoid validating a field that can't be validated.
func ValueIsNotEmptyOrUnknown(in interface{}) bool {
	switch val := in.(type) {
	case string:
		// consider strings potentially unknown if empty or set to the SDK dummy value
		return val != "" && val != PluginSdkUnknownValue
	default:
		// for all other types, treat them as known
		return true
	}
}
