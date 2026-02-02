package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DelegatedAdminAccessAssignment{}

type DelegatedAdminAccessAssignment struct {
	AccessContainer *DelegatedAdminAccessContainer `json:"accessContainer,omitempty"`
	AccessDetails   *DelegatedAdminAccessDetails   `json:"accessDetails,omitempty"`

	// The date and time in ISO 8601 format and in UTC time when the access assignment was created. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The date and time in ISO 8601 and in UTC time when this access assignment was last modified. Read-only.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The status of the access assignment. Read-only. The possible values are: pending, active, deleting, deleted, error,
	// unknownFutureValue.
	Status *DelegatedAdminAccessAssignmentStatus `json:"status,omitempty"`

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

func (s DelegatedAdminAccessAssignment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DelegatedAdminAccessAssignment{}

func (s DelegatedAdminAccessAssignment) MarshalJSON() ([]byte, error) {
	type wrapper DelegatedAdminAccessAssignment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DelegatedAdminAccessAssignment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DelegatedAdminAccessAssignment: %+v", err)
	}

	delete(decoded, "createdDateTime")
	delete(decoded, "lastModifiedDateTime")
	delete(decoded, "status")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.delegatedAdminAccessAssignment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DelegatedAdminAccessAssignment: %+v", err)
	}

	return encoded, nil
}
