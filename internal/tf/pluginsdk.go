package tf

// PluginSdkUnknownValue is a dummy value used/sent by the plugin SDK when a real value is not known at plan time,
// e.g. during a CustomizeDiff function
// See https://github.com/hashicorp/terraform-plugin-sdk/blob/main/internal/configs/hcl2shim/values.go#L16
const PluginSdkUnknownValue = "74D93920-ED26-11E3-AC10-0800200C9A66"
