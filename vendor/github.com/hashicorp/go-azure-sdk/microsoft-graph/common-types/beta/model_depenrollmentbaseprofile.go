package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DepEnrollmentBaseProfile interface {
	Entity
	EnrollmentProfile
	DepEnrollmentBaseProfile() BaseDepEnrollmentBaseProfileImpl
}

var _ DepEnrollmentBaseProfile = BaseDepEnrollmentBaseProfileImpl{}

type BaseDepEnrollmentBaseProfileImpl struct {
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

func (s BaseDepEnrollmentBaseProfileImpl) DepEnrollmentBaseProfile() BaseDepEnrollmentBaseProfileImpl {
	return s
}

func (s BaseDepEnrollmentBaseProfileImpl) EnrollmentProfile() BaseEnrollmentProfileImpl {
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

func (s BaseDepEnrollmentBaseProfileImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ DepEnrollmentBaseProfile = RawDepEnrollmentBaseProfileImpl{}

// RawDepEnrollmentBaseProfileImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDepEnrollmentBaseProfileImpl struct {
	depEnrollmentBaseProfile BaseDepEnrollmentBaseProfileImpl
	Type                     string
	Values                   map[string]interface{}
}

func (s RawDepEnrollmentBaseProfileImpl) DepEnrollmentBaseProfile() BaseDepEnrollmentBaseProfileImpl {
	return s.depEnrollmentBaseProfile
}

func (s RawDepEnrollmentBaseProfileImpl) EnrollmentProfile() BaseEnrollmentProfileImpl {
	return s.depEnrollmentBaseProfile.EnrollmentProfile()
}

func (s RawDepEnrollmentBaseProfileImpl) Entity() BaseEntityImpl {
	return s.depEnrollmentBaseProfile.Entity()
}

var _ json.Marshaler = BaseDepEnrollmentBaseProfileImpl{}

func (s BaseDepEnrollmentBaseProfileImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseDepEnrollmentBaseProfileImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseDepEnrollmentBaseProfileImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseDepEnrollmentBaseProfileImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.depEnrollmentBaseProfile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseDepEnrollmentBaseProfileImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalDepEnrollmentBaseProfileImplementation(input []byte) (DepEnrollmentBaseProfile, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DepEnrollmentBaseProfile into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.depIOSEnrollmentProfile") {
		var out DepIOSEnrollmentProfile
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DepIOSEnrollmentProfile: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.depMacOSEnrollmentProfile") {
		var out DepMacOSEnrollmentProfile
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DepMacOSEnrollmentProfile: %+v", err)
		}
		return out, nil
	}

	var parent BaseDepEnrollmentBaseProfileImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDepEnrollmentBaseProfileImpl: %+v", err)
	}

	return RawDepEnrollmentBaseProfileImpl{
		depEnrollmentBaseProfile: parent,
		Type:                     value,
		Values:                   temp,
	}, nil

}
