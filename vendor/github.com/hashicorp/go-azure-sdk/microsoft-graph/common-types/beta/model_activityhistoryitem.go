package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ActivityHistoryItem{}

type ActivityHistoryItem struct {
	// Optional. The duration of active user engagement. if not supplied, this is calculated from the startedDateTime and
	// lastActiveDateTime.
	ActiveDurationSeconds nullable.Type[int64] `json:"activeDurationSeconds,omitempty"`

	Activity *UserActivity `json:"activity,omitempty"`

	// Set by the server. DateTime in UTC when the object was created on the server.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Optional. UTC DateTime when the activityHistoryItem will undergo hard-delete. Can be set by the client.
	ExpirationDateTime nullable.Type[string] `json:"expirationDateTime,omitempty"`

	// Optional. UTC DateTime when the activityHistoryItem (activity session) was last understood as active or finished - if
	// null, activityHistoryItem status should be Ongoing.
	LastActiveDateTime nullable.Type[string] `json:"lastActiveDateTime,omitempty"`

	// Set by the server. DateTime in UTC when the object was modified on the server.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Required. UTC DateTime when the activityHistoryItem (activity session) was started. Required for timeline history.
	StartedDateTime string `json:"startedDateTime"`

	// Set by the server. A status code used to identify valid objects. Values: active, updated, deleted, ignored.
	Status *Status `json:"status,omitempty"`

	// Optional. The timezone in which the user's device used to generate the activity was located at activity creation
	// time. Values supplied as Olson IDs in order to support cross-platform representation.
	UserTimezone nullable.Type[string] `json:"userTimezone,omitempty"`

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

func (s ActivityHistoryItem) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ActivityHistoryItem{}

func (s ActivityHistoryItem) MarshalJSON() ([]byte, error) {
	type wrapper ActivityHistoryItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ActivityHistoryItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ActivityHistoryItem: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.activityHistoryItem"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ActivityHistoryItem: %+v", err)
	}

	return encoded, nil
}
