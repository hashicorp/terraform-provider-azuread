package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = GcpRole{}

type GcpRole struct {
	// The name of the GCP role. Supports $filter and (eq,contains).
	DisplayName *string `json:"displayName,omitempty"`

	// The ID of the GCP role as defined by GCP. Alternate key.
	ExternalId *string `json:"externalId,omitempty"`

	GcpRoleType *GcpRoleType `json:"gcpRoleType,omitempty"`

	// Resources that an identity assigned this GCP role can perform actions on. Supports $filter and (eq).
	Scopes *[]GcpScope `json:"scopes,omitempty"`

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

func (s GcpRole) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = GcpRole{}

func (s GcpRole) MarshalJSON() ([]byte, error) {
	type wrapper GcpRole
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GcpRole: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GcpRole: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.gcpRole"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GcpRole: %+v", err)
	}

	return encoded, nil
}
