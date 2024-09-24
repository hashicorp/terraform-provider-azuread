package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AndroidForWorkSettings{}

type AndroidForWorkSettings struct {
	// Bind status of the tenant with the Google EMM API
	BindStatus *AndroidForWorkBindStatus `json:"bindStatus,omitempty"`

	// Indicates if this account is flighting for Android Device Owner Management with CloudDPC.
	DeviceOwnerManagementEnabled *bool `json:"deviceOwnerManagementEnabled,omitempty"`

	// Android for Work device management targeting type for the account
	EnrollmentTarget *AndroidForWorkEnrollmentTarget `json:"enrollmentTarget,omitempty"`

	// Last completion time for app sync
	LastAppSyncDateTime nullable.Type[string] `json:"lastAppSyncDateTime,omitempty"`

	// Sync status of the tenant with the Google EMM API
	LastAppSyncStatus *AndroidForWorkSyncStatus `json:"lastAppSyncStatus,omitempty"`

	// Last modification time for Android for Work settings
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Organization name used when onboarding Android for Work
	OwnerOrganizationName nullable.Type[string] `json:"ownerOrganizationName,omitempty"`

	// Owner UPN that created the enterprise
	OwnerUserPrincipalName nullable.Type[string] `json:"ownerUserPrincipalName,omitempty"`

	// Specifies which AAD groups can enroll devices in Android for Work device management if enrollmentTarget is set to
	// 'Targeted'
	TargetGroupIds *[]string `json:"targetGroupIds,omitempty"`

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

func (s AndroidForWorkSettings) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AndroidForWorkSettings{}

func (s AndroidForWorkSettings) MarshalJSON() ([]byte, error) {
	type wrapper AndroidForWorkSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AndroidForWorkSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AndroidForWorkSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.androidForWorkSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AndroidForWorkSettings: %+v", err)
	}

	return encoded, nil
}
