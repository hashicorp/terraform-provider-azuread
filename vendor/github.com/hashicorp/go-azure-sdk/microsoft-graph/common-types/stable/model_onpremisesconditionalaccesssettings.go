package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = OnPremisesConditionalAccessSettings{}

type OnPremisesConditionalAccessSettings struct {
	// Indicates if on premises conditional access is enabled for this organization
	Enabled *bool `json:"enabled,omitempty"`

	// User groups that will be exempt by on premises conditional access. All users in these groups will be exempt from the
	// conditional access policy.
	ExcludedGroups *[]string `json:"excludedGroups,omitempty"`

	// User groups that will be targeted by on premises conditional access. All users in these groups will be required to
	// have mobile device managed and compliant for mail access.
	IncludedGroups *[]string `json:"includedGroups,omitempty"`

	// Override the default access rule when allowing a device to ensure access is granted.
	OverrideDefaultRule *bool `json:"overrideDefaultRule,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s OnPremisesConditionalAccessSettings) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OnPremisesConditionalAccessSettings{}

func (s OnPremisesConditionalAccessSettings) MarshalJSON() ([]byte, error) {
	type wrapper OnPremisesConditionalAccessSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OnPremisesConditionalAccessSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OnPremisesConditionalAccessSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.onPremisesConditionalAccessSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OnPremisesConditionalAccessSettings: %+v", err)
	}

	return encoded, nil
}
