package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceConfigurationUserStateSummary{}

type DeviceConfigurationUserStateSummary struct {
	// Number of compliant users
	CompliantUserCount *int64 `json:"compliantUserCount,omitempty"`

	// Number of conflict users
	ConflictUserCount *int64 `json:"conflictUserCount,omitempty"`

	// Number of error users
	ErrorUserCount *int64 `json:"errorUserCount,omitempty"`

	// Number of NonCompliant users
	NonCompliantUserCount *int64 `json:"nonCompliantUserCount,omitempty"`

	// Number of not applicable users
	NotApplicableUserCount *int64 `json:"notApplicableUserCount,omitempty"`

	// Number of remediated users
	RemediatedUserCount *int64 `json:"remediatedUserCount,omitempty"`

	// Number of unknown users
	UnknownUserCount *int64 `json:"unknownUserCount,omitempty"`

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

func (s DeviceConfigurationUserStateSummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceConfigurationUserStateSummary{}

func (s DeviceConfigurationUserStateSummary) MarshalJSON() ([]byte, error) {
	type wrapper DeviceConfigurationUserStateSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceConfigurationUserStateSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceConfigurationUserStateSummary: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceConfigurationUserStateSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceConfigurationUserStateSummary: %+v", err)
	}

	return encoded, nil
}
