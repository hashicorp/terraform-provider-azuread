package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecurityIdentityContainer{}

type SecurityIdentityContainer struct {
	// Represents potential issues within a customer's Microsoft Defender for Identity configuration that Microsoft Defender
	// for Identity identified.
	HealthIssues *[]SecurityHealthIssue `json:"healthIssues,omitempty"`

	// Represents a customer's Microsoft Defender for Identity sensors.
	Sensors *[]SecuritySensor `json:"sensors,omitempty"`

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

func (s SecurityIdentityContainer) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityIdentityContainer{}

func (s SecurityIdentityContainer) MarshalJSON() ([]byte, error) {
	type wrapper SecurityIdentityContainer
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityIdentityContainer: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityIdentityContainer: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.identityContainer"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityIdentityContainer: %+v", err)
	}

	return encoded, nil
}
