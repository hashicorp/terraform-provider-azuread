package stable

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodFeatureConfiguration struct {
	// A single entity that is excluded from this feature.
	ExcludeTarget *FeatureTarget `json:"excludeTarget,omitempty"`

	// A single entity that is included in this feature.
	IncludeTarget *FeatureTarget `json:"includeTarget,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Enable or disable the feature. Possible values are: default, enabled, disabled, unknownFutureValue. The default value
	// is used when the configuration hasn't been explicitly set and uses the default behavior of Microsoft Entra ID for the
	// setting. The default value is disabled.
	State *AdvancedConfigState `json:"state,omitempty"`
}
