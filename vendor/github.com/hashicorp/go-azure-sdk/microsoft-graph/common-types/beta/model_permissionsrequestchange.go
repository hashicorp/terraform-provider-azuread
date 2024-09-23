package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PermissionsRequestChange{}

type PermissionsRequestChange struct {
	// The status of the active occurence of the schedule if one exists. The possible values are: grantingFailed, granted,
	// granting, revoked, revoking, revokingFailed, unknownFutureValue.
	ActiveOccurrenceStatus *PermissionsRequestOccurrenceStatus `json:"activeOccurrenceStatus,omitempty"`

	// Time when the change occurred.
	ModificationDateTime *string `json:"modificationDateTime,omitempty"`

	// The ID of the scheduledPermissionsRequest object.
	PermissionsRequestId *string `json:"permissionsRequestId,omitempty"`

	StatusDetail *StatusDetail `json:"statusDetail,omitempty"`

	// Represents the ticketing system identifier.
	TicketId nullable.Type[string] `json:"ticketId,omitempty"`

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

func (s PermissionsRequestChange) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PermissionsRequestChange{}

func (s PermissionsRequestChange) MarshalJSON() ([]byte, error) {
	type wrapper PermissionsRequestChange
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PermissionsRequestChange: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PermissionsRequestChange: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.permissionsRequestChange"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PermissionsRequestChange: %+v", err)
	}

	return encoded, nil
}
