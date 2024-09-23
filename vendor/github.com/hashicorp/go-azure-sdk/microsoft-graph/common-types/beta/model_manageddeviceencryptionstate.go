package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedDeviceEncryptionState{}

type ManagedDeviceEncryptionState struct {
	// Advanced BitLocker State. Possible values are: success, noUserConsent, osVolumeUnprotected, osVolumeTpmRequired,
	// osVolumeTpmOnlyRequired, osVolumeTpmPinRequired, osVolumeTpmStartupKeyRequired, osVolumeTpmPinStartupKeyRequired,
	// osVolumeEncryptionMethodMismatch, recoveryKeyBackupFailed, fixedDriveNotEncrypted,
	// fixedDriveEncryptionMethodMismatch, loggedOnUserNonAdmin, windowsRecoveryEnvironmentNotConfigured, tpmNotAvailable,
	// tpmNotReady, networkError.
	AdvancedBitLockerStates *AdvancedBitLockerState `json:"advancedBitLockerStates,omitempty"`

	// Device name
	DeviceName nullable.Type[string] `json:"deviceName,omitempty"`

	// Device type.
	DeviceType *DeviceTypes `json:"deviceType,omitempty"`

	EncryptionPolicySettingState *ComplianceStatus `json:"encryptionPolicySettingState,omitempty"`

	// Encryption readiness state
	EncryptionReadinessState *EncryptionReadinessState `json:"encryptionReadinessState,omitempty"`

	// Encryption state
	EncryptionState *EncryptionState `json:"encryptionState,omitempty"`

	// FileVault State. Possible values are: success, driveEncryptedByUser, userDeferredEncryption, escrowNotEnabled.
	FileVaultStates *FileVaultState `json:"fileVaultStates,omitempty"`

	// Operating system version of the device
	OsVersion nullable.Type[string] `json:"osVersion,omitempty"`

	// Policy Details
	PolicyDetails *[]EncryptionReportPolicyDetails `json:"policyDetails,omitempty"`

	// Device TPM Version
	TpmSpecificationVersion nullable.Type[string] `json:"tpmSpecificationVersion,omitempty"`

	// User name
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

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

func (s ManagedDeviceEncryptionState) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedDeviceEncryptionState{}

func (s ManagedDeviceEncryptionState) MarshalJSON() ([]byte, error) {
	type wrapper ManagedDeviceEncryptionState
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedDeviceEncryptionState: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedDeviceEncryptionState: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedDeviceEncryptionState"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedDeviceEncryptionState: %+v", err)
	}

	return encoded, nil
}
