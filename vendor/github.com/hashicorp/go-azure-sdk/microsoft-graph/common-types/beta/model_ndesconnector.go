package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = NdesConnector{}

type NdesConnector struct {
	// The build version of the Ndes Connector.
	ConnectorVersion nullable.Type[string] `json:"connectorVersion,omitempty"`

	// The friendly name of the Ndes Connector.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Timestamp when on-prem certificate connector was enrolled in Intune.
	EnrolledDateTime *string `json:"enrolledDateTime,omitempty"`

	// Last connection time for the Ndes Connector
	LastConnectionDateTime *string `json:"lastConnectionDateTime,omitempty"`

	// Name of the machine running on-prem certificate connector service.
	MachineName nullable.Type[string] `json:"machineName,omitempty"`

	// List of Scope Tags for this Entity instance.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// The current status of the Ndes Connector.
	State *NdesConnectorState `json:"state,omitempty"`

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

func (s NdesConnector) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NdesConnector{}

func (s NdesConnector) MarshalJSON() ([]byte, error) {
	type wrapper NdesConnector
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NdesConnector: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NdesConnector: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.ndesConnector"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NdesConnector: %+v", err)
	}

	return encoded, nil
}
