package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EntitlementManagementSettings{}

type EntitlementManagementSettings struct {
	// If externalUserLifecycleAction is BlockSignInAndDelete, the number of days after an external user is blocked from
	// sign in before their account is deleted.
	DaysUntilExternalUserDeletedAfterBlocked nullable.Type[int64] `json:"daysUntilExternalUserDeletedAfterBlocked,omitempty"`

	// One of None, BlockSignIn, or BlockSignInAndDelete.
	ExternalUserLifecycleAction nullable.Type[string] `json:"externalUserLifecycleAction,omitempty"`

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

func (s EntitlementManagementSettings) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EntitlementManagementSettings{}

func (s EntitlementManagementSettings) MarshalJSON() ([]byte, error) {
	type wrapper EntitlementManagementSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EntitlementManagementSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EntitlementManagementSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.entitlementManagementSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EntitlementManagementSettings: %+v", err)
	}

	return encoded, nil
}
