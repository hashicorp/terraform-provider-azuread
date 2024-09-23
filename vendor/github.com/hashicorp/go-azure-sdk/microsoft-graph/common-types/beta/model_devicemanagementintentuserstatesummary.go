package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceManagementIntentUserStateSummary{}

type DeviceManagementIntentUserStateSummary struct {
	// Number of users in conflict
	ConflictCount *int64 `json:"conflictCount,omitempty"`

	// Number of error users
	ErrorCount *int64 `json:"errorCount,omitempty"`

	// Number of failed users
	FailedCount *int64 `json:"failedCount,omitempty"`

	// Number of not applicable users
	NotApplicableCount *int64 `json:"notApplicableCount,omitempty"`

	// Number of succeeded users
	SuccessCount *int64 `json:"successCount,omitempty"`

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

func (s DeviceManagementIntentUserStateSummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementIntentUserStateSummary{}

func (s DeviceManagementIntentUserStateSummary) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementIntentUserStateSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementIntentUserStateSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementIntentUserStateSummary: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementIntentUserStateSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementIntentUserStateSummary: %+v", err)
	}

	return encoded, nil
}
