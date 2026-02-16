package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ConnectorGroup{}

type ConnectorGroup struct {
	Applications       *[]Application      `json:"applications,omitempty"`
	ConnectorGroupType *ConnectorGroupType `json:"connectorGroupType,omitempty"`

	// Indicates if the connectorGroup is the default connectorGroup. Only a single connector group can be the default
	// connectorGroup and this is pre-set by the system. Read-only.
	IsDefault *bool `json:"isDefault,omitempty"`

	Members *[]Connector `json:"members,omitempty"`

	// The name associated with the connectorGroup.
	Name *string `json:"name,omitempty"`

	// The region the connectorGroup is assigned to and will optimize traffic for. This region can only be set if no
	// connectors or applications are assigned to the connectorGroup. The possible values are: nam (for North America), eur
	// (for Europe), aus (for Australia), asia (for Asia), ind (for India), and unknownFutureValue.
	Region *ConnectorGroupRegion `json:"region,omitempty"`

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

func (s ConnectorGroup) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ConnectorGroup{}

func (s ConnectorGroup) MarshalJSON() ([]byte, error) {
	type wrapper ConnectorGroup
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ConnectorGroup: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ConnectorGroup: %+v", err)
	}

	delete(decoded, "isDefault")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.connectorGroup"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ConnectorGroup: %+v", err)
	}

	return encoded, nil
}
