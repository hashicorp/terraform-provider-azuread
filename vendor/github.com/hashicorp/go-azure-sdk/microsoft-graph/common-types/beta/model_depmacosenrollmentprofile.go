package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DepEnrollmentBaseProfile = DepMacOSEnrollmentProfile{}

type DepMacOSEnrollmentProfile struct {
	// Indicates if Accessibility screen is disabled
	AccessibilityScreenDisabled *bool `json:"accessibilityScreenDisabled,omitempty"`

	// Indicates what the full name for the admin account is
	AdminAccountFullName nullable.Type[string] `json:"adminAccountFullName,omitempty"`

	// Indicates what the password for the admin account is
	AdminAccountPassword nullable.Type[string] `json:"adminAccountPassword,omitempty"`

	// Indicates what the user name for the admin account is
	AdminAccountUserName nullable.Type[string] `json:"adminAccountUserName,omitempty"`

	// Indicates if Setup Assistant will automatically advance through its screen
	AutoAdvanceSetupEnabled *bool `json:"autoAdvanceSetupEnabled,omitempty"`

	// Indicates if UnlockWithWatch screen is disabled
	AutoUnlockWithWatchDisabled *bool `json:"autoUnlockWithWatchDisabled,omitempty"`

	// Indicates if iCloud Documents and Desktop screen is disabled
	ChooseYourLockScreenDisabled *bool `json:"chooseYourLockScreenDisabled,omitempty"`

	// Indicates whether Setup Assistant will auto populate the primary account information
	DontAutoPopulatePrimaryAccountInfo *bool `json:"dontAutoPopulatePrimaryAccountInfo,omitempty"`

	// Indicates whether the user will enable blockediting
	EnableRestrictEditing *bool `json:"enableRestrictEditing,omitempty"`

	// Indicates if file vault is disabled
	FileVaultDisabled *bool `json:"fileVaultDisabled,omitempty"`

	// Indicates whether the admin account should be hidded or not
	HideAdminAccount *bool `json:"hideAdminAccount,omitempty"`

	// Indicates if iCloud Analytics screen is disabled
	ICloudDiagnosticsDisabled *bool `json:"iCloudDiagnosticsDisabled,omitempty"`

	// Indicates if iCloud Documents and Desktop screen is disabled
	ICloudStorageDisabled *bool `json:"iCloudStorageDisabled,omitempty"`

	// Indicates if Passcode setup pane is disabled
	PassCodeDisabled *bool `json:"passCodeDisabled,omitempty"`

	// Indicates what the full name for the primary account is
	PrimaryAccountFullName nullable.Type[string] `json:"primaryAccountFullName,omitempty"`

	// Indicates what the account name for the primary account is
	PrimaryAccountUserName nullable.Type[string] `json:"primaryAccountUserName,omitempty"`

	// Indicates if registration is disabled
	RegistrationDisabled *bool `json:"registrationDisabled,omitempty"`

	// Indicates if the device is network-tethered to run the command
	RequestRequiresNetworkTether *bool `json:"requestRequiresNetworkTether,omitempty"`

	// Indicates whether Setup Assistant will set the account as a regular user
	SetPrimarySetupAccountAsRegularUser *bool `json:"setPrimarySetupAccountAsRegularUser,omitempty"`

	// Indicates whether Setup Assistant will skip the user interface for primary account setup
	SkipPrimarySetupAccountCreation *bool `json:"skipPrimarySetupAccountCreation,omitempty"`

	// Indicates if zoom setup pane is disabled
	ZoomDisabled *bool `json:"zoomDisabled,omitempty"`

	// Fields inherited from DepEnrollmentBaseProfile

	// Indicates if Apple id setup pane is disabled
	AppleIdDisabled *bool `json:"appleIdDisabled,omitempty"`

	// Indicates if Apple pay setup pane is disabled
	ApplePayDisabled *bool `json:"applePayDisabled,omitempty"`

	// URL for setup assistant login
	ConfigurationWebUrl *bool `json:"configurationWebUrl,omitempty"`

	// Sets a literal or name pattern.
	DeviceNameTemplate nullable.Type[string] `json:"deviceNameTemplate,omitempty"`

	// Indicates if diagnostics setup pane is disabled
	DiagnosticsDisabled *bool `json:"diagnosticsDisabled,omitempty"`

	// Indicates if displaytone setup screen is disabled
	DisplayToneSetupDisabled *bool `json:"displayToneSetupDisabled,omitempty"`

	// enabledSkipKeys contains all the enabled skip keys as strings
	EnabledSkipKeys *[]string `json:"enabledSkipKeys,omitempty"`

	// EnrollmentTimeAzureAdGroupIds contains list of enrollment time Azure Group Ids to be associated with profile
	EnrollmentTimeAzureAdGroupIds *[]string `json:"enrollmentTimeAzureAdGroupIds,omitempty"`

	// Indicates if this is the default profile
	IsDefault *bool `json:"isDefault,omitempty"`

	// Indicates if the profile is mandatory
	IsMandatory *bool `json:"isMandatory,omitempty"`

	// Indicates if Location service setup pane is disabled
	LocationDisabled *bool `json:"locationDisabled,omitempty"`

	// Indicates if privacy screen is disabled
	PrivacyPaneDisabled *bool `json:"privacyPaneDisabled,omitempty"`

	// Indicates if the profile removal option is disabled
	ProfileRemovalDisabled *bool `json:"profileRemovalDisabled,omitempty"`

	// Indicates if Restore setup pane is blocked
	RestoreBlocked *bool `json:"restoreBlocked,omitempty"`

	// Indicates if screen timeout setup is disabled
	ScreenTimeScreenDisabled *bool `json:"screenTimeScreenDisabled,omitempty"`

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

	// Indicates if the device will need to wait for configured confirmation
	WaitForDeviceConfiguredConfirmation *bool `json:"waitForDeviceConfiguredConfirmation,omitempty"`

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

func (s DepMacOSEnrollmentProfile) DepEnrollmentBaseProfile() BaseDepEnrollmentBaseProfileImpl {
	return BaseDepEnrollmentBaseProfileImpl{
		AppleIdDisabled:                      s.AppleIdDisabled,
		ApplePayDisabled:                     s.ApplePayDisabled,
		ConfigurationWebUrl:                  s.ConfigurationWebUrl,
		DeviceNameTemplate:                   s.DeviceNameTemplate,
		DiagnosticsDisabled:                  s.DiagnosticsDisabled,
		DisplayToneSetupDisabled:             s.DisplayToneSetupDisabled,
		EnabledSkipKeys:                      s.EnabledSkipKeys,
		EnrollmentTimeAzureAdGroupIds:        s.EnrollmentTimeAzureAdGroupIds,
		IsDefault:                            s.IsDefault,
		IsMandatory:                          s.IsMandatory,
		LocationDisabled:                     s.LocationDisabled,
		PrivacyPaneDisabled:                  s.PrivacyPaneDisabled,
		ProfileRemovalDisabled:               s.ProfileRemovalDisabled,
		RestoreBlocked:                       s.RestoreBlocked,
		ScreenTimeScreenDisabled:             s.ScreenTimeScreenDisabled,
		SiriDisabled:                         s.SiriDisabled,
		SupervisedModeEnabled:                s.SupervisedModeEnabled,
		SupportDepartment:                    s.SupportDepartment,
		SupportPhoneNumber:                   s.SupportPhoneNumber,
		TermsAndConditionsDisabled:           s.TermsAndConditionsDisabled,
		TouchIdDisabled:                      s.TouchIdDisabled,
		WaitForDeviceConfiguredConfirmation:  s.WaitForDeviceConfiguredConfirmation,
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

func (s DepMacOSEnrollmentProfile) EnrollmentProfile() BaseEnrollmentProfileImpl {
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

func (s DepMacOSEnrollmentProfile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DepMacOSEnrollmentProfile{}

func (s DepMacOSEnrollmentProfile) MarshalJSON() ([]byte, error) {
	type wrapper DepMacOSEnrollmentProfile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DepMacOSEnrollmentProfile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DepMacOSEnrollmentProfile: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.depMacOSEnrollmentProfile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DepMacOSEnrollmentProfile: %+v", err)
	}

	return encoded, nil
}
