package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AndroidManagedStoreAccountEnterpriseSettings{}

type AndroidManagedStoreAccountEnterpriseSettings struct {
	// Company codes for AndroidManagedStoreAccountEnterpriseSettings
	AndroidDeviceOwnerFullyManagedEnrollmentEnabled *bool `json:"androidDeviceOwnerFullyManagedEnrollmentEnabled,omitempty"`

	// Bind status of the tenant with the Google EMM API
	BindStatus *AndroidManagedStoreAccountBindStatus `json:"bindStatus,omitempty"`

	// Company codes for AndroidManagedStoreAccountEnterpriseSettings
	CompanyCodes *[]AndroidEnrollmentCompanyCode `json:"companyCodes,omitempty"`

	// Indicates if this account is flighting for Android Device Owner Management with CloudDPC.
	DeviceOwnerManagementEnabled *bool `json:"deviceOwnerManagementEnabled,omitempty"`

	// Android for Work device management targeting type for the account
	EnrollmentTarget *AndroidManagedStoreAccountEnrollmentTarget `json:"enrollmentTarget,omitempty"`

	// Last completion time for app sync
	LastAppSyncDateTime nullable.Type[string] `json:"lastAppSyncDateTime,omitempty"`

	// Sync status of the tenant with the Google EMM API
	LastAppSyncStatus *AndroidManagedStoreAccountAppSyncStatus `json:"lastAppSyncStatus,omitempty"`

	// Last modification time for Android enterprise settings
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Initial scope tags for MGP apps
	ManagedGooglePlayInitialScopeTagIds *[]string `json:"managedGooglePlayInitialScopeTagIds,omitempty"`

	// Organization name used when onboarding Android Enterprise
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

func (s AndroidManagedStoreAccountEnterpriseSettings) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AndroidManagedStoreAccountEnterpriseSettings{}

func (s AndroidManagedStoreAccountEnterpriseSettings) MarshalJSON() ([]byte, error) {
	type wrapper AndroidManagedStoreAccountEnterpriseSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AndroidManagedStoreAccountEnterpriseSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AndroidManagedStoreAccountEnterpriseSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.androidManagedStoreAccountEnterpriseSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AndroidManagedStoreAccountEnterpriseSettings: %+v", err)
	}

	return encoded, nil
}
