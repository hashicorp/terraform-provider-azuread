package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentifierUriRestriction struct {
	// Collection of custom security attribute exemptions. If an actor user or service principal has the custom security
	// attribute defined in this section, they're exempted from the restriction. This means that calls the user or service
	// principal makes to create or update apps are exempt from this policy enforcement.
	ExcludeActors *AppManagementPolicyActorExemptions `json:"excludeActors,omitempty"`

	// If true, the restriction isn't enforced for applications that are configured to receive V2 tokens in Microsoft Entra
	// ID; else, the restriction isn't enforced for those applications.
	ExcludeAppsReceivingV2Tokens nullable.Type[bool] `json:"excludeAppsReceivingV2Tokens,omitempty"`

	// If true, the restriction isn't enforced for SAML applications in Microsoft Entra ID; else, the restriction is
	// enforced for those applications.
	ExcludeSaml nullable.Type[bool] `json:"excludeSaml,omitempty"`

	// If true, Microsoft sets the identifierUriRestriction state. If false, the tenant modifies the
	// identifierUriRestriction state. Read-only.
	IsStateSetByMicrosoft *bool `json:"isStateSetByMicrosoft,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Specifies the date from which the policy restriction applies to newly created applications. For existing
	// applications, the enforcement date can be retroactively applied.
	RestrictForAppsCreatedAfterDateTime nullable.Type[string] `json:"restrictForAppsCreatedAfterDateTime,omitempty"`

	State *AppManagementRestrictionState `json:"state,omitempty"`
}

var _ json.Marshaler = IdentifierUriRestriction{}

func (s IdentifierUriRestriction) MarshalJSON() ([]byte, error) {
	type wrapper IdentifierUriRestriction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IdentifierUriRestriction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentifierUriRestriction: %+v", err)
	}

	delete(decoded, "isStateSetByMicrosoft")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IdentifierUriRestriction: %+v", err)
	}

	return encoded, nil
}
