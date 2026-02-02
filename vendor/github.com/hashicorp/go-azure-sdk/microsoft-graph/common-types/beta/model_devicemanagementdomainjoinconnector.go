package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceManagementDomainJoinConnector{}

type DeviceManagementDomainJoinConnector struct {
	// The connector display name.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Last time connector contacted Intune.
	LastConnectionDateTime *string `json:"lastConnectionDateTime,omitempty"`

	// The ODJ request states.
	State *DeviceManagementDomainJoinConnectorState `json:"state,omitempty"`

	// The version of the connector.
	Version *string `json:"version,omitempty"`

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

func (s DeviceManagementDomainJoinConnector) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementDomainJoinConnector{}

func (s DeviceManagementDomainJoinConnector) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementDomainJoinConnector
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementDomainJoinConnector: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementDomainJoinConnector: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementDomainJoinConnector"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementDomainJoinConnector: %+v", err)
	}

	return encoded, nil
}
