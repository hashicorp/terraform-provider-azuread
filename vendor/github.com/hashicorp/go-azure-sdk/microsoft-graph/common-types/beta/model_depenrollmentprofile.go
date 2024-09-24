package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EnrollmentProfile = DepEnrollmentProfile{}

type DepEnrollmentProfile struct {
	// Indicates if Apple id setup pane is disabled
	AppleIdDisabled *bool `json:"appleIdDisabled,omitempty"`

	// Indicates if Apple pay setup pane is disabled
	ApplePayDisabled *bool `json:"applePayDisabled,omitempty"`

	// Indicates if the device will need to wait for configured confirmation
	AwaitDeviceConfiguredConfirmation *bool `json:"awaitDeviceConfiguredConfirmation,omitempty"`

	// Indicates if diagnostics setup pane is disabled
	DiagnosticsDisabled *bool `json:"diagnosticsDisabled,omitempty"`

	// This indicates whether the device is to be enrolled in a mode which enables multi user scenarios. Only applicable in
	// shared iPads.
	EnableSharedIPad *bool `json:"enableSharedIPad,omitempty"`

	ITunesPairingMode *ITunesPairingMode `json:"iTunesPairingMode,omitempty"`

	// Indicates if this is the default profile
	IsDefault *bool `json:"isDefault,omitempty"`

	// Indicates if the profile is mandatory
	IsMandatory *bool `json:"isMandatory,omitempty"`

	// Indicates if Location service setup pane is disabled
	LocationDisabled *bool `json:"locationDisabled,omitempty"`

	// Indicates if Mac OS file vault is disabled
	MacOSFileVaultDisabled *bool `json:"macOSFileVaultDisabled,omitempty"`

	// Indicates if Mac OS registration is disabled
	MacOSRegistrationDisabled *bool `json:"macOSRegistrationDisabled,omitempty"`

	// Management certificates for Apple Configurator
	ManagementCertificates *[]ManagementCertificateWithThumbprint `json:"managementCertificates,omitempty"`

	// Indicates if Passcode setup pane is disabled
	PassCodeDisabled *bool `json:"passCodeDisabled,omitempty"`

	// Indicates if the profile removal option is disabled
	ProfileRemovalDisabled *bool `json:"profileRemovalDisabled,omitempty"`

	// Indicates if Restore setup pane is blocked
	RestoreBlocked *bool `json:"restoreBlocked,omitempty"`

	// Indicates if Restore from Android is disabled
	RestoreFromAndroidDisabled *bool `json:"restoreFromAndroidDisabled,omitempty"`

	// This specifies the maximum number of users that can use a shared iPad. Only applicable in shared iPad mode.
	SharedIPadMaximumUserCount *int64 `json:"sharedIPadMaximumUserCount,omitempty"`

	// Indicates if siri setup pane is disabled
	SiriDisabled *bool `json:"siriDisabled,omitempty"`

	// Supervised mode, True to enable, false otherwise. See
	// https://learn.microsoft.com/intune/deploy-use/enroll-devices-in-microsoft-intune for additional information.
	SupervisedModeEnabled *bool `json:"supervisedModeEnabled,omitempty"`

	// Support department information
	SupportDepartment nullable.Type[string] `json:"supportDepartment,omitempty"`

	// Support phone number
	SupportPhoneNumber nullable.Type[string] `json:"supportPhoneNumber,omitempty"`

	// Indicates if 'Terms and Conditions' setup pane is disabled
	TermsAndConditionsDisabled *bool `json:"termsAndConditionsDisabled,omitempty"`

	// Indicates if touch id setup pane is disabled
	TouchIdDisabled *bool `json:"touchIdDisabled,omitempty"`

	// Indicates if zoom setup pane is disabled
	ZoomDisabled *bool `json:"zoomDisabled,omitempty"`

	// Fields inherited from EnrollmentProfile

	// Configuration endpoint url to use for Enrollment
	ConfigurationEndpointUrl nullable.Type[string] `json:"configurationEndpointUrl,omitempty"`

	// Description of the profile
	Description nullable.Type[string] `json:"description,omitempty"`

	// Name of the profile
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Indicates to authenticate with Apple Setup Assistant instead of Company Portal.
	EnableAuthenticationViaCompanyPortal *bool `json:"enableAuthenticationViaCompanyPortal,omitempty"`

	// Indicates that Company Portal is required on setup assistant enrolled devices
	RequireCompanyPortalOnSetupAssistantEnrolledDevices *bool `json:"requireCompanyPortalOnSetupAssistantEnrolledDevices,omitempty"`

	// Indicates if the profile requires user authentication
	RequiresUserAuthentication *bool `json:"requiresUserAuthentication,omitempty"`

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

func (s DepEnrollmentProfile) EnrollmentProfile() BaseEnrollmentProfileImpl {
	return BaseEnrollmentProfileImpl{
		ConfigurationEndpointUrl:             s.ConfigurationEndpointUrl,
		Description:                          s.Description,
		DisplayName:                          s.DisplayName,
		EnableAuthenticationViaCompanyPortal: s.EnableAuthenticationViaCompanyPortal,
		RequireCompanyPortalOnSetupAssistantEnrolledDevices: s.RequireCompanyPortalOnSetupAssistantEnrolledDevices,
		RequiresUserAuthentication:                          s.RequiresUserAuthentication,
		Id:                                                  s.Id,
		ODataId:                                             s.ODataId,
		ODataType:                                           s.ODataType,
	}
}

func (s DepEnrollmentProfile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DepEnrollmentProfile{}

func (s DepEnrollmentProfile) MarshalJSON() ([]byte, error) {
	type wrapper DepEnrollmentProfile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DepEnrollmentProfile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DepEnrollmentProfile: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.depEnrollmentProfile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DepEnrollmentProfile: %+v", err)
	}

	return encoded, nil
}
