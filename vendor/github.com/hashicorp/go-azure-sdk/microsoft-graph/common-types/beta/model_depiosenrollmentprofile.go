package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DepEnrollmentBaseProfile = DepIOSEnrollmentProfile{}

type DepIOSEnrollmentProfile struct {
	// Indicates if Apperance screen is disabled
	AppearanceScreenDisabled *bool `json:"appearanceScreenDisabled,omitempty"`

	// Indicates if the device will need to wait for configured confirmation
	AwaitDeviceConfiguredConfirmation *bool `json:"awaitDeviceConfiguredConfirmation,omitempty"`

	// Carrier URL for activating device eSIM.
	CarrierActivationUrl nullable.Type[string] `json:"carrierActivationUrl,omitempty"`

	// If set, indicates which Vpp token should be used to deploy the Company Portal w/ device licensing.
	// 'enableAuthenticationViaCompanyPortal' must be set in order for this property to be set.
	CompanyPortalVppTokenId nullable.Type[string] `json:"companyPortalVppTokenId,omitempty"`

	// Indicates if Device To Device Migration is disabled
	DeviceToDeviceMigrationDisabled *bool `json:"deviceToDeviceMigrationDisabled,omitempty"`

	// This indicates whether the device is to be enrolled in a mode which enables multi user scenarios. Only applicable in
	// shared iPads.
	EnableSharedIPad *bool `json:"enableSharedIPad,omitempty"`

	// Tells the device to enable single app mode and apply app-lock during enrollment. Default is false.
	// 'enableAuthenticationViaCompanyPortal' and 'companyPortalVppTokenId' must be set for this property to be set.
	EnableSingleAppEnrollmentMode *bool `json:"enableSingleAppEnrollmentMode,omitempty"`

	// Indicates if Express Language screen is disabled
	ExpressLanguageScreenDisabled *bool `json:"expressLanguageScreenDisabled,omitempty"`

	// Indicates if temporary sessions is enabled
	ForceTemporarySession *bool `json:"forceTemporarySession,omitempty"`

	// Indicates if home button sensitivity screen is disabled
	HomeButtonScreenDisabled *bool `json:"homeButtonScreenDisabled,omitempty"`

	// Indicates if iMessage and FaceTime screen is disabled
	IMessageAndFaceTimeScreenDisabled *bool `json:"iMessageAndFaceTimeScreenDisabled,omitempty"`

	ITunesPairingMode *ITunesPairingMode `json:"iTunesPairingMode,omitempty"`

	// Management certificates for Apple Configurator
	ManagementCertificates *[]ManagementCertificateWithThumbprint `json:"managementCertificates,omitempty"`

	// Indicates if onboarding setup screen is disabled
	OnBoardingScreenDisabled *bool `json:"onBoardingScreenDisabled,omitempty"`

	// Indicates if Passcode setup pane is disabled
	PassCodeDisabled *bool `json:"passCodeDisabled,omitempty"`

	// Indicates timeout before locked screen requires the user to enter the device passocde to unlock it
	PasscodeLockGracePeriodInSeconds nullable.Type[int64] `json:"passcodeLockGracePeriodInSeconds,omitempty"`

	// Indicates if Preferred language screen is disabled
	PreferredLanguageScreenDisabled *bool `json:"preferredLanguageScreenDisabled,omitempty"`

	// Indicates if Weclome screen is disabled
	RestoreCompletedScreenDisabled *bool `json:"restoreCompletedScreenDisabled,omitempty"`

	// Indicates if Restore from Android is disabled
	RestoreFromAndroidDisabled *bool `json:"restoreFromAndroidDisabled,omitempty"`

	// This specifies the maximum number of users that can use a shared iPad. Only applicable in shared iPad mode.
	SharedIPadMaximumUserCount *int64 `json:"sharedIPadMaximumUserCount,omitempty"`

	// Indicates if the SIMSetup screen is disabled
	SimSetupScreenDisabled *bool `json:"simSetupScreenDisabled,omitempty"`

	// Indicates if the mandatory sofware update screen is disabled
	SoftwareUpdateScreenDisabled *bool `json:"softwareUpdateScreenDisabled,omitempty"`

	// Indicates timeout of temporary session
	TemporarySessionTimeoutInSeconds *int64 `json:"temporarySessionTimeoutInSeconds,omitempty"`

	// Indicates if Weclome screen is disabled
	UpdateCompleteScreenDisabled *bool `json:"updateCompleteScreenDisabled,omitempty"`

	// Indicates timeout of temporary session
	UserSessionTimeoutInSeconds *int64 `json:"userSessionTimeoutInSeconds,omitempty"`

	// Indicates that this apple device is designated to support 'shared device mode' scenarios. This is distinct from the
	// 'shared iPad' scenario. See https://learn.microsoft.com/mem/intune/enrollment/device-enrollment-shared-ios
	UserlessSharedAadModeEnabled *bool `json:"userlessSharedAadModeEnabled,omitempty"`

	// Indicates if the watch migration screen is disabled
	WatchMigrationScreenDisabled *bool `json:"watchMigrationScreenDisabled,omitempty"`

	// Indicates if Weclome screen is disabled
	WelcomeScreenDisabled *bool `json:"welcomeScreenDisabled,omitempty"`

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

func (s DepIOSEnrollmentProfile) DepEnrollmentBaseProfile() BaseDepEnrollmentBaseProfileImpl {
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

func (s DepIOSEnrollmentProfile) EnrollmentProfile() BaseEnrollmentProfileImpl {
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

func (s DepIOSEnrollmentProfile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DepIOSEnrollmentProfile{}

func (s DepIOSEnrollmentProfile) MarshalJSON() ([]byte, error) {
	type wrapper DepIOSEnrollmentProfile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DepIOSEnrollmentProfile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DepIOSEnrollmentProfile: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.depIOSEnrollmentProfile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DepIOSEnrollmentProfile: %+v", err)
	}

	return encoded, nil
}
