package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AppleUserInitiatedEnrollmentProfile{}

type AppleUserInitiatedEnrollmentProfile struct {
	// The list of assignments for this profile.
	Assignments *[]AppleEnrollmentProfileAssignment `json:"assignments,omitempty"`

	// List of available enrollment type options
	AvailableEnrollmentTypeOptions *[]AppleOwnerTypeEnrollmentType `json:"availableEnrollmentTypeOptions,omitempty"`

	// Profile creation time
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	DefaultEnrollmentType *AppleUserInitiatedEnrollmentType `json:"defaultEnrollmentType,omitempty"`

	// Description of the profile
	Description nullable.Type[string] `json:"description,omitempty"`

	// Name of the profile
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Profile last modified time
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Supported platform types.
	Platform *DevicePlatformType `json:"platform,omitempty"`

	// Priority, 0 is highest
	Priority *int64 `json:"priority,omitempty"`

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

func (s AppleUserInitiatedEnrollmentProfile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AppleUserInitiatedEnrollmentProfile{}

func (s AppleUserInitiatedEnrollmentProfile) MarshalJSON() ([]byte, error) {
	type wrapper AppleUserInitiatedEnrollmentProfile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AppleUserInitiatedEnrollmentProfile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AppleUserInitiatedEnrollmentProfile: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.appleUserInitiatedEnrollmentProfile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AppleUserInitiatedEnrollmentProfile: %+v", err)
	}

	return encoded, nil
}
