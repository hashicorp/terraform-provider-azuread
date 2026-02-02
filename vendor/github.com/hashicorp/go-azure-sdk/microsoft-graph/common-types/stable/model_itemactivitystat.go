package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ItemActivityStat{}

type ItemActivityStat struct {
	// Statistics about the access actions in this interval. Read-only.
	Access *ItemActionStat `json:"access,omitempty"`

	// Exposes the itemActivities represented in this itemActivityStat resource.
	Activities *[]ItemActivity `json:"activities,omitempty"`

	// Statistics about the create actions in this interval. Read-only.
	Create *ItemActionStat `json:"create,omitempty"`

	// Statistics about the delete actions in this interval. Read-only.
	Delete *ItemActionStat `json:"delete,omitempty"`

	// Statistics about the edit actions in this interval. Read-only.
	Edit *ItemActionStat `json:"edit,omitempty"`

	// When the interval ends. Read-only.
	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`

	// Indicates that the statistics in this interval are based on incomplete data. Read-only.
	IncompleteData *IncompleteData `json:"incompleteData,omitempty"`

	// Indicates whether the item is 'trending.' Read-only.
	IsTrending nullable.Type[bool] `json:"isTrending,omitempty"`

	// Statistics about the move actions in this interval. Read-only.
	Move *ItemActionStat `json:"move,omitempty"`

	// When the interval starts. Read-only.
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`

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

func (s ItemActivityStat) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ItemActivityStat{}

func (s ItemActivityStat) MarshalJSON() ([]byte, error) {
	type wrapper ItemActivityStat
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ItemActivityStat: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ItemActivityStat: %+v", err)
	}

	delete(decoded, "access")
	delete(decoded, "create")
	delete(decoded, "delete")
	delete(decoded, "edit")
	delete(decoded, "endDateTime")
	delete(decoded, "incompleteData")
	delete(decoded, "isTrending")
	delete(decoded, "move")
	delete(decoded, "startDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.itemActivityStat"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ItemActivityStat: %+v", err)
	}

	return encoded, nil
}
