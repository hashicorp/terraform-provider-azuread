package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Connector{}

type Connector struct {
	// The external IP address as detected by the connector server. Read-only.
	ExternalIp *string `json:"externalIp,omitempty"`

	// The name of the computer on which the connector is installed and runs on.
	MachineName *string `json:"machineName,omitempty"`

	// The connectorGroup that the connector is a member of. Read-only.
	MemberOf *[]ConnectorGroup `json:"memberOf,omitempty"`

	Status *ConnectorStatus `json:"status,omitempty"`

	// The version of the connector. Read-only.
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

func (s Connector) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Connector{}

func (s Connector) MarshalJSON() ([]byte, error) {
	type wrapper Connector
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Connector: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Connector: %+v", err)
	}

	delete(decoded, "externalIp")
	delete(decoded, "memberOf")
	delete(decoded, "version")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.connector"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Connector: %+v", err)
	}

	return encoded, nil
}
