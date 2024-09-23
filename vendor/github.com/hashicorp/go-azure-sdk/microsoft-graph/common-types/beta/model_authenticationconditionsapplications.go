package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationConditionsApplications struct {
	// Whether the custom authentication extension should trigger for all applications with appIds specified in the
	// includeApplications relationship. This property must be set to false for listener of type
	// onTokenIssuanceStartListener.
	IncludeAllApplications *bool `json:"includeAllApplications,omitempty"`

	IncludeApplications *[]AuthenticationConditionApplication `json:"includeApplications,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
