package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceEnrollmentConfiguration = DeviceEnrollmentWindowsHelloForBusinessConfiguration{}

type DeviceEnrollmentWindowsHelloForBusinessConfiguration struct {
	// Possible values of a property
	EnhancedBiometricsState *Enablement `json:"enhancedBiometricsState,omitempty"`

	// Setting to configure Enhanced sign-in security. Default is Not Configured
	EnhancedSignInSecurity *int64 `json:"enhancedSignInSecurity,omitempty"`

	// Controls the period of time (in days) that a PIN can be used before the system requires the user to change it. This
	// must be set between 0 and 730, inclusive. If set to 0, the user's PIN will never expire
	PinExpirationInDays *int64 `json:"pinExpirationInDays,omitempty"`

	// Windows Hello for Business pin usage options
	PinLowercaseCharactersUsage *WindowsHelloForBusinessPinUsage `json:"pinLowercaseCharactersUsage,omitempty"`

	// Controls the maximum number of characters allowed for the Windows Hello for Business PIN. This value must be between
	// 4 and 127, inclusive. This value must be greater than or equal to the value set for the minimum PIN.
	PinMaximumLength *int64 `json:"pinMaximumLength,omitempty"`

	// Controls the minimum number of characters required for the Windows Hello for Business PIN. This value must be between
	// 4 and 127, inclusive, and less than or equal to the value set for the maximum PIN.
	PinMinimumLength *int64 `json:"pinMinimumLength,omitempty"`

	// Controls the ability to prevent users from using past PINs. This must be set between 0 and 50, inclusive, and the
	// current PIN of the user is included in that count. If set to 0, previous PINs are not stored. PIN history is not
	// preserved through a PIN reset.
	PinPreviousBlockCount *int64 `json:"pinPreviousBlockCount,omitempty"`

	// Windows Hello for Business pin usage options
	PinSpecialCharactersUsage *WindowsHelloForBusinessPinUsage `json:"pinSpecialCharactersUsage,omitempty"`

	// Windows Hello for Business pin usage options
	PinUppercaseCharactersUsage *WindowsHelloForBusinessPinUsage `json:"pinUppercaseCharactersUsage,omitempty"`

	// Controls the use of Remote Windows Hello for Business. Remote Windows Hello for Business provides the ability for a
	// portable, registered device to be usable as a companion for desktop authentication. The desktop must be Azure AD
	// joined and the companion device must have a Windows Hello for Business PIN.
	RemotePassportEnabled *bool `json:"remotePassportEnabled,omitempty"`

	// Controls whether to require a Trusted Platform Module (TPM) for provisioning Windows Hello for Business. A TPM
	// provides an additional security benefit in that data stored on it cannot be used on other devices. If set to False,
	// all devices can provision Windows Hello for Business even if there is not a usable TPM.
	SecurityDeviceRequired *bool `json:"securityDeviceRequired,omitempty"`

	// Possible values of a property
	SecurityKeyForSignIn *Enablement `json:"securityKeyForSignIn,omitempty"`

	// Possible values of a property
	State *Enablement `json:"state,omitempty"`

	// Controls the use of biometric gestures, such as face and fingerprint, as an alternative to the Windows Hello for
	// Business PIN. If set to False, biometric gestures are not allowed. Users must still configure a PIN as a backup in
	// case of failures.
	UnlockWithBiometricsEnabled *bool `json:"unlockWithBiometricsEnabled,omitempty"`

	// Fields inherited from DeviceEnrollmentConfiguration

	// The list of group assignments for the device configuration profile
	Assignments *[]EnrollmentConfigurationAssignment `json:"assignments,omitempty"`

	// Created date time in UTC of the device enrollment configuration
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The description of the device enrollment configuration
	Description nullable.Type[string] `json:"description,omitempty"`

	// Describes the TemplateFamily for the Template entity
	DeviceEnrollmentConfigurationType *DeviceEnrollmentConfigurationType `json:"deviceEnrollmentConfigurationType,omitempty"`

	// The display name of the device enrollment configuration
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Last modified date time in UTC of the device enrollment configuration
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Priority is used when a user exists in multiple groups that are assigned enrollment configuration. Users are subject
	// only to the configuration with the lowest priority value.
	Priority *int64 `json:"priority,omitempty"`

	// Optional role scope tags for the enrollment restrictions.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// The version of the device enrollment configuration
	Version *int64 `json:"version,omitempty"`

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

func (s DeviceEnrollmentWindowsHelloForBusinessConfiguration) DeviceEnrollmentConfiguration() BaseDeviceEnrollmentConfigurationImpl {
	return BaseDeviceEnrollmentConfigurationImpl{
		Assignments:                       s.Assignments,
		CreatedDateTime:                   s.CreatedDateTime,
		Description:                       s.Description,
		DeviceEnrollmentConfigurationType: s.DeviceEnrollmentConfigurationType,
		DisplayName:                       s.DisplayName,
		LastModifiedDateTime:              s.LastModifiedDateTime,
		Priority:                          s.Priority,
		RoleScopeTagIds:                   s.RoleScopeTagIds,
		Version:                           s.Version,
		Id:                                s.Id,
		ODataId:                           s.ODataId,
		ODataType:                         s.ODataType,
	}
}

func (s DeviceEnrollmentWindowsHelloForBusinessConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceEnrollmentWindowsHelloForBusinessConfiguration{}

func (s DeviceEnrollmentWindowsHelloForBusinessConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper DeviceEnrollmentWindowsHelloForBusinessConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceEnrollmentWindowsHelloForBusinessConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceEnrollmentWindowsHelloForBusinessConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceEnrollmentWindowsHelloForBusinessConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceEnrollmentWindowsHelloForBusinessConfiguration: %+v", err)
	}

	return encoded, nil
}
