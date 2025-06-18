package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DepOnboardingSetting{}

type DepOnboardingSetting struct {
	// The Apple ID used to obtain the current token.
	AppleIdentifier nullable.Type[string] `json:"appleIdentifier,omitempty"`

	// Consent granted for data sharing with Apple Dep Service
	DataSharingConsentGranted *bool `json:"dataSharingConsentGranted,omitempty"`

	// Default iOS Enrollment Profile
	DefaultIosEnrollmentProfile *DepIOSEnrollmentProfile `json:"defaultIosEnrollmentProfile,omitempty"`

	// Default MacOs Enrollment Profile
	DefaultMacOsEnrollmentProfile *DepMacOSEnrollmentProfile `json:"defaultMacOsEnrollmentProfile,omitempty"`

	// Default TvOS Enrollment Profile
	DefaultTvOSEnrollmentProfile *DepTvOSEnrollmentProfile `json:"defaultTvOSEnrollmentProfile,omitempty"`

	// Default VisionOS Enrollment Profile
	DefaultVisionOSEnrollmentProfile *DepVisionOSEnrollmentProfile `json:"defaultVisionOSEnrollmentProfile,omitempty"`

	// The enrollment profiles.
	EnrollmentProfiles *[]EnrollmentProfile `json:"enrollmentProfiles,omitempty"`

	// The imported Apple device identities.
	ImportedAppleDeviceIdentities *[]ImportedAppleDeviceIdentity `json:"importedAppleDeviceIdentities,omitempty"`

	// When the service was onboarded.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// When the service last syned with Intune
	LastSuccessfulSyncDateTime *string `json:"lastSuccessfulSyncDateTime,omitempty"`

	// Error code reported by Apple during last dep sync.
	LastSyncErrorCode *int64 `json:"lastSyncErrorCode,omitempty"`

	// When Intune last requested a sync.
	LastSyncTriggeredDateTime *string `json:"lastSyncTriggeredDateTime,omitempty"`

	// List of Scope Tags for this Entity instance.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// Whether or not the Dep token sharing is enabled with the School Data Sync service.
	ShareTokenWithSchoolDataSyncService *bool `json:"shareTokenWithSchoolDataSyncService,omitempty"`

	// Gets synced device count
	SyncedDeviceCount *int64 `json:"syncedDeviceCount,omitempty"`

	// When the token will expire.
	TokenExpirationDateTime *string `json:"tokenExpirationDateTime,omitempty"`

	// Friendly Name for Dep Token
	TokenName nullable.Type[string] `json:"tokenName,omitempty"`

	TokenType *DepTokenType `json:"tokenType,omitempty"`

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

func (s DepOnboardingSetting) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DepOnboardingSetting{}

func (s DepOnboardingSetting) MarshalJSON() ([]byte, error) {
	type wrapper DepOnboardingSetting
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DepOnboardingSetting: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DepOnboardingSetting: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.depOnboardingSetting"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DepOnboardingSetting: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DepOnboardingSetting{}

func (s *DepOnboardingSetting) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AppleIdentifier                     nullable.Type[string]         `json:"appleIdentifier,omitempty"`
		DataSharingConsentGranted           *bool                         `json:"dataSharingConsentGranted,omitempty"`
		DefaultIosEnrollmentProfile         *DepIOSEnrollmentProfile      `json:"defaultIosEnrollmentProfile,omitempty"`
		DefaultMacOsEnrollmentProfile       *DepMacOSEnrollmentProfile    `json:"defaultMacOsEnrollmentProfile,omitempty"`
		DefaultTvOSEnrollmentProfile        *DepTvOSEnrollmentProfile     `json:"defaultTvOSEnrollmentProfile,omitempty"`
		DefaultVisionOSEnrollmentProfile    *DepVisionOSEnrollmentProfile `json:"defaultVisionOSEnrollmentProfile,omitempty"`
		LastModifiedDateTime                *string                       `json:"lastModifiedDateTime,omitempty"`
		LastSuccessfulSyncDateTime          *string                       `json:"lastSuccessfulSyncDateTime,omitempty"`
		LastSyncErrorCode                   *int64                        `json:"lastSyncErrorCode,omitempty"`
		LastSyncTriggeredDateTime           *string                       `json:"lastSyncTriggeredDateTime,omitempty"`
		RoleScopeTagIds                     *[]string                     `json:"roleScopeTagIds,omitempty"`
		ShareTokenWithSchoolDataSyncService *bool                         `json:"shareTokenWithSchoolDataSyncService,omitempty"`
		SyncedDeviceCount                   *int64                        `json:"syncedDeviceCount,omitempty"`
		TokenExpirationDateTime             *string                       `json:"tokenExpirationDateTime,omitempty"`
		TokenName                           nullable.Type[string]         `json:"tokenName,omitempty"`
		TokenType                           *DepTokenType                 `json:"tokenType,omitempty"`
		Id                                  *string                       `json:"id,omitempty"`
		ODataId                             *string                       `json:"@odata.id,omitempty"`
		ODataType                           *string                       `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AppleIdentifier = decoded.AppleIdentifier
	s.DataSharingConsentGranted = decoded.DataSharingConsentGranted
	s.DefaultIosEnrollmentProfile = decoded.DefaultIosEnrollmentProfile
	s.DefaultMacOsEnrollmentProfile = decoded.DefaultMacOsEnrollmentProfile
	s.DefaultTvOSEnrollmentProfile = decoded.DefaultTvOSEnrollmentProfile
	s.DefaultVisionOSEnrollmentProfile = decoded.DefaultVisionOSEnrollmentProfile
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.LastSuccessfulSyncDateTime = decoded.LastSuccessfulSyncDateTime
	s.LastSyncErrorCode = decoded.LastSyncErrorCode
	s.LastSyncTriggeredDateTime = decoded.LastSyncTriggeredDateTime
	s.RoleScopeTagIds = decoded.RoleScopeTagIds
	s.ShareTokenWithSchoolDataSyncService = decoded.ShareTokenWithSchoolDataSyncService
	s.SyncedDeviceCount = decoded.SyncedDeviceCount
	s.TokenExpirationDateTime = decoded.TokenExpirationDateTime
	s.TokenName = decoded.TokenName
	s.TokenType = decoded.TokenType
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DepOnboardingSetting into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["enrollmentProfiles"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling EnrollmentProfiles into list []json.RawMessage: %+v", err)
		}

		output := make([]EnrollmentProfile, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalEnrollmentProfileImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'EnrollmentProfiles' for 'DepOnboardingSetting': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.EnrollmentProfiles = &output
	}

	if v, ok := temp["importedAppleDeviceIdentities"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ImportedAppleDeviceIdentities into list []json.RawMessage: %+v", err)
		}

		output := make([]ImportedAppleDeviceIdentity, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalImportedAppleDeviceIdentityImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ImportedAppleDeviceIdentities' for 'DepOnboardingSetting': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ImportedAppleDeviceIdentities = &output
	}

	return nil
}
