package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SoftwareUpdateStatusSummary{}

type SoftwareUpdateStatusSummary struct {
	// Number of compliant devices.
	CompliantDeviceCount *int64 `json:"compliantDeviceCount,omitempty"`

	// Number of compliant users.
	CompliantUserCount *int64 `json:"compliantUserCount,omitempty"`

	// Number of conflict devices.
	ConflictDeviceCount *int64 `json:"conflictDeviceCount,omitempty"`

	// Number of conflict users.
	ConflictUserCount *int64 `json:"conflictUserCount,omitempty"`

	// The name of the policy.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Number of devices had error.
	ErrorDeviceCount *int64 `json:"errorDeviceCount,omitempty"`

	// Number of users had error.
	ErrorUserCount *int64 `json:"errorUserCount,omitempty"`

	// Number of non compliant devices.
	NonCompliantDeviceCount *int64 `json:"nonCompliantDeviceCount,omitempty"`

	// Number of non compliant users.
	NonCompliantUserCount *int64 `json:"nonCompliantUserCount,omitempty"`

	// Number of not applicable devices.
	NotApplicableDeviceCount *int64 `json:"notApplicableDeviceCount,omitempty"`

	// Number of not applicable users.
	NotApplicableUserCount *int64 `json:"notApplicableUserCount,omitempty"`

	// Number of remediated devices.
	RemediatedDeviceCount *int64 `json:"remediatedDeviceCount,omitempty"`

	// Number of remediated users.
	RemediatedUserCount *int64 `json:"remediatedUserCount,omitempty"`

	// Number of unknown devices.
	UnknownDeviceCount *int64 `json:"unknownDeviceCount,omitempty"`

	// Number of unknown users.
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

func (s SoftwareUpdateStatusSummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SoftwareUpdateStatusSummary{}

func (s SoftwareUpdateStatusSummary) MarshalJSON() ([]byte, error) {
	type wrapper SoftwareUpdateStatusSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SoftwareUpdateStatusSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SoftwareUpdateStatusSummary: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.softwareUpdateStatusSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SoftwareUpdateStatusSummary: %+v", err)
	}

	return encoded, nil
}
