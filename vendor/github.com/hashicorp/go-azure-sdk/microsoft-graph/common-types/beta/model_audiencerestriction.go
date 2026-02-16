package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AudienceRestriction struct {
	// Collection of custom security attribute exemptions. If an actor user or service principal has the custom security
	// attribute defined in this section, they're exempted from the restriction. This means that calls the user or service
	// principal makes to create or update apps are exempt from this policy enforcement.
	ExcludeActors *AppManagementPolicyActorExemptions `json:"excludeActors,omitempty"`

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

var _ json.Marshaler = AudienceRestriction{}

func (s AudienceRestriction) MarshalJSON() ([]byte, error) {
	type wrapper AudienceRestriction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AudienceRestriction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AudienceRestriction: %+v", err)
	}

	delete(decoded, "isStateSetByMicrosoft")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AudienceRestriction: %+v", err)
	}

	return encoded, nil
}
