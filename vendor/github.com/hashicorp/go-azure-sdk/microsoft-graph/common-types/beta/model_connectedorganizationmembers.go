package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ UserSet = ConnectedOrganizationMembers{}

type ConnectedOrganizationMembers struct {
	// The name of the connected organization. Read only.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The ID of the connected organization in entitlement management.
	Id nullable.Type[string] `json:"id,omitempty"`

	// Fields inherited from UserSet

	// For a user in an approval stage, this property indicates whether the user is a backup fallback approver.
	IsBackup nullable.Type[bool] `json:"isBackup,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s ConnectedOrganizationMembers) UserSet() BaseUserSetImpl {
	return BaseUserSetImpl{
		IsBackup:  s.IsBackup,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ConnectedOrganizationMembers{}

func (s ConnectedOrganizationMembers) MarshalJSON() ([]byte, error) {
	type wrapper ConnectedOrganizationMembers
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ConnectedOrganizationMembers: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ConnectedOrganizationMembers: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.connectedOrganizationMembers"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ConnectedOrganizationMembers: %+v", err)
	}

	return encoded, nil
}
