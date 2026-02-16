package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CopilotAdminLimitedMode{}

type CopilotAdminLimitedMode struct {
	// The ID of a Microsoft Entra group, for which the value of isEnabledForGroup is applied. The default value is null. If
	// isEnabledForGroup is set to true, the groupId value must be provided for the Copilot limited mode in Teams meetings
	// to be enabled for the members of the group. Optional.
	GroupId nullable.Type[string] `json:"groupId,omitempty"`

	// Enables the user to be in limited mode for Copilot in Teams meetings. When copilotAdminLimitedMode=true, users in
	// this mode can ask any questions, but Copilot doesn't respond to certain questions related to inferring emotions,
	// behavior, or judgments. When copilotAdminLimitedMode=false, it responds to all types of questions grounded to the
	// meeting conversation. The default value is false.
	IsEnabledForGroup nullable.Type[bool] `json:"isEnabledForGroup,omitempty"`

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

func (s CopilotAdminLimitedMode) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CopilotAdminLimitedMode{}

func (s CopilotAdminLimitedMode) MarshalJSON() ([]byte, error) {
	type wrapper CopilotAdminLimitedMode
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CopilotAdminLimitedMode: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CopilotAdminLimitedMode: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.copilotAdminLimitedMode"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CopilotAdminLimitedMode: %+v", err)
	}

	return encoded, nil
}
