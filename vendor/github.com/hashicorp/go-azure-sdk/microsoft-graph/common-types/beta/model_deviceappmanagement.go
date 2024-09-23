package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceAppManagement{}

type DeviceAppManagement struct {
	// Android managed app policies.
	AndroidManagedAppProtections *[]AndroidManagedAppProtection `json:"androidManagedAppProtections,omitempty"`

	// Default managed app policies.
	DefaultManagedAppProtections *[]DefaultManagedAppProtection `json:"defaultManagedAppProtections,omitempty"`

	// Device app management tasks.
	DeviceAppManagementTasks *[]DeviceAppManagementTask `json:"deviceAppManagementTasks,omitempty"`

	// The Windows Enterprise Code Signing Certificate.
	EnterpriseCodeSigningCertificates *[]EnterpriseCodeSigningCertificate `json:"enterpriseCodeSigningCertificates,omitempty"`

	// The IOS Lob App Provisioning Configurations.
	IosLobAppProvisioningConfigurations *[]IosLobAppProvisioningConfiguration `json:"iosLobAppProvisioningConfigurations,omitempty"`

	// iOS managed app policies.
	IosManagedAppProtections *[]IosManagedAppProtection `json:"iosManagedAppProtections,omitempty"`

	// Whether the account is enabled for syncing applications from the Microsoft Store for Business.
	IsEnabledForMicrosoftStoreForBusiness *bool `json:"isEnabledForMicrosoftStoreForBusiness,omitempty"`

	// Managed app policies.
	ManagedAppPolicies *[]ManagedAppPolicy `json:"managedAppPolicies,omitempty"`

	// The managed app registrations.
	ManagedAppRegistrations *[]ManagedAppRegistration `json:"managedAppRegistrations,omitempty"`

	// The managed app statuses.
	ManagedAppStatuses *[]ManagedAppStatus `json:"managedAppStatuses,omitempty"`

	// The mobile eBook categories.
	ManagedEBookCategories *[]ManagedEBookCategory `json:"managedEBookCategories,omitempty"`

	// The Managed eBook.
	ManagedEBooks *[]ManagedEBook `json:"managedEBooks,omitempty"`

	// Windows information protection for apps running on devices which are MDM enrolled.
	MdmWindowsInformationProtectionPolicies *[]MdmWindowsInformationProtectionPolicy `json:"mdmWindowsInformationProtectionPolicies,omitempty"`

	// The locale information used to sync applications from the Microsoft Store for Business. Cultures that are specific to
	// a country/region. The names of these cultures follow RFC 4646 (Windows Vista and later). The format is
	// -<country/regioncode2>, where is a lowercase two-letter code derived from ISO 639-1 and <country/regioncode2> is an
	// uppercase two-letter code derived from ISO 3166. For example, en-US for English (United States) is a specific
	// culture.
	MicrosoftStoreForBusinessLanguage nullable.Type[string] `json:"microsoftStoreForBusinessLanguage,omitempty"`

	// The last time an application sync from the Microsoft Store for Business was completed.
	MicrosoftStoreForBusinessLastCompletedApplicationSyncTime *string `json:"microsoftStoreForBusinessLastCompletedApplicationSyncTime,omitempty"`

	// The last time the apps from the Microsoft Store for Business were synced successfully for the account.
	MicrosoftStoreForBusinessLastSuccessfulSyncDateTime *string `json:"microsoftStoreForBusinessLastSuccessfulSyncDateTime,omitempty"`

	// Portal to which admin syncs available Microsoft Store for Business apps. This is available in the Intune Admin
	// Console.
	MicrosoftStoreForBusinessPortalSelection *MicrosoftStoreForBusinessPortalSelectionOptions `json:"microsoftStoreForBusinessPortalSelection,omitempty"`

	// MobileAppCatalogPackage entities.
	MobileAppCatalogPackages *[]MobileAppCatalogPackage `json:"mobileAppCatalogPackages,omitempty"`

	// The mobile app categories.
	MobileAppCategories *[]MobileAppCategory `json:"mobileAppCategories,omitempty"`

	// The Managed Device Mobile Application Configurations.
	MobileAppConfigurations *[]ManagedDeviceMobileAppConfiguration `json:"mobileAppConfigurations,omitempty"`

	// The mobile apps.
	MobileApps *[]MobileApp `json:"mobileApps,omitempty"`

	// The PolicySet of Policies and Applications
	PolicySets *[]PolicySet `json:"policySets,omitempty"`

	// The WinPhone Symantec Code Signing Certificate.
	SymantecCodeSigningCertificate *SymantecCodeSigningCertificate `json:"symantecCodeSigningCertificate,omitempty"`

	// Targeted managed app configurations.
	TargetedManagedAppConfigurations *[]TargetedManagedAppConfiguration `json:"targetedManagedAppConfigurations,omitempty"`

	// List of Vpp tokens for this organization.
	VppTokens *[]VppToken `json:"vppTokens,omitempty"`

	// The collection of Windows Defender Application Control Supplemental Policies.
	WdacSupplementalPolicies *[]WindowsDefenderApplicationControlSupplementalPolicy `json:"wdacSupplementalPolicies,omitempty"`

	// Windows information protection device registrations that are not MDM enrolled.
	WindowsInformationProtectionDeviceRegistrations *[]WindowsInformationProtectionDeviceRegistration `json:"windowsInformationProtectionDeviceRegistrations,omitempty"`

	// Windows information protection for apps running on devices which are not MDM enrolled.
	WindowsInformationProtectionPolicies *[]WindowsInformationProtectionPolicy `json:"windowsInformationProtectionPolicies,omitempty"`

	// Windows information protection wipe actions.
	WindowsInformationProtectionWipeActions *[]WindowsInformationProtectionWipeAction `json:"windowsInformationProtectionWipeActions,omitempty"`

	// Windows managed app policies.
	WindowsManagedAppProtections *[]WindowsManagedAppProtection `json:"windowsManagedAppProtections,omitempty"`

	// Windows management app.
	WindowsManagementApp *WindowsManagementApp `json:"windowsManagementApp,omitempty"`

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

func (s DeviceAppManagement) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceAppManagement{}

func (s DeviceAppManagement) MarshalJSON() ([]byte, error) {
	type wrapper DeviceAppManagement
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceAppManagement: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceAppManagement: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceAppManagement"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceAppManagement: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DeviceAppManagement{}

func (s *DeviceAppManagement) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AndroidManagedAppProtections                              *[]AndroidManagedAppProtection                         `json:"androidManagedAppProtections,omitempty"`
		DefaultManagedAppProtections                              *[]DefaultManagedAppProtection                         `json:"defaultManagedAppProtections,omitempty"`
		EnterpriseCodeSigningCertificates                         *[]EnterpriseCodeSigningCertificate                    `json:"enterpriseCodeSigningCertificates,omitempty"`
		IosLobAppProvisioningConfigurations                       *[]IosLobAppProvisioningConfiguration                  `json:"iosLobAppProvisioningConfigurations,omitempty"`
		IosManagedAppProtections                                  *[]IosManagedAppProtection                             `json:"iosManagedAppProtections,omitempty"`
		IsEnabledForMicrosoftStoreForBusiness                     *bool                                                  `json:"isEnabledForMicrosoftStoreForBusiness,omitempty"`
		ManagedEBookCategories                                    *[]ManagedEBookCategory                                `json:"managedEBookCategories,omitempty"`
		MdmWindowsInformationProtectionPolicies                   *[]MdmWindowsInformationProtectionPolicy               `json:"mdmWindowsInformationProtectionPolicies,omitempty"`
		MicrosoftStoreForBusinessLanguage                         nullable.Type[string]                                  `json:"microsoftStoreForBusinessLanguage,omitempty"`
		MicrosoftStoreForBusinessLastCompletedApplicationSyncTime *string                                                `json:"microsoftStoreForBusinessLastCompletedApplicationSyncTime,omitempty"`
		MicrosoftStoreForBusinessLastSuccessfulSyncDateTime       *string                                                `json:"microsoftStoreForBusinessLastSuccessfulSyncDateTime,omitempty"`
		MicrosoftStoreForBusinessPortalSelection                  *MicrosoftStoreForBusinessPortalSelectionOptions       `json:"microsoftStoreForBusinessPortalSelection,omitempty"`
		MobileAppCategories                                       *[]MobileAppCategory                                   `json:"mobileAppCategories,omitempty"`
		PolicySets                                                *[]PolicySet                                           `json:"policySets,omitempty"`
		SymantecCodeSigningCertificate                            *SymantecCodeSigningCertificate                        `json:"symantecCodeSigningCertificate,omitempty"`
		TargetedManagedAppConfigurations                          *[]TargetedManagedAppConfiguration                     `json:"targetedManagedAppConfigurations,omitempty"`
		VppTokens                                                 *[]VppToken                                            `json:"vppTokens,omitempty"`
		WdacSupplementalPolicies                                  *[]WindowsDefenderApplicationControlSupplementalPolicy `json:"wdacSupplementalPolicies,omitempty"`
		WindowsInformationProtectionDeviceRegistrations           *[]WindowsInformationProtectionDeviceRegistration      `json:"windowsInformationProtectionDeviceRegistrations,omitempty"`
		WindowsInformationProtectionPolicies                      *[]WindowsInformationProtectionPolicy                  `json:"windowsInformationProtectionPolicies,omitempty"`
		WindowsInformationProtectionWipeActions                   *[]WindowsInformationProtectionWipeAction              `json:"windowsInformationProtectionWipeActions,omitempty"`
		WindowsManagedAppProtections                              *[]WindowsManagedAppProtection                         `json:"windowsManagedAppProtections,omitempty"`
		WindowsManagementApp                                      *WindowsManagementApp                                  `json:"windowsManagementApp,omitempty"`
		Id                                                        *string                                                `json:"id,omitempty"`
		ODataId                                                   *string                                                `json:"@odata.id,omitempty"`
		ODataType                                                 *string                                                `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AndroidManagedAppProtections = decoded.AndroidManagedAppProtections
	s.DefaultManagedAppProtections = decoded.DefaultManagedAppProtections
	s.EnterpriseCodeSigningCertificates = decoded.EnterpriseCodeSigningCertificates
	s.IosLobAppProvisioningConfigurations = decoded.IosLobAppProvisioningConfigurations
	s.IosManagedAppProtections = decoded.IosManagedAppProtections
	s.IsEnabledForMicrosoftStoreForBusiness = decoded.IsEnabledForMicrosoftStoreForBusiness
	s.ManagedEBookCategories = decoded.ManagedEBookCategories
	s.MdmWindowsInformationProtectionPolicies = decoded.MdmWindowsInformationProtectionPolicies
	s.MicrosoftStoreForBusinessLanguage = decoded.MicrosoftStoreForBusinessLanguage
	s.MicrosoftStoreForBusinessLastCompletedApplicationSyncTime = decoded.MicrosoftStoreForBusinessLastCompletedApplicationSyncTime
	s.MicrosoftStoreForBusinessLastSuccessfulSyncDateTime = decoded.MicrosoftStoreForBusinessLastSuccessfulSyncDateTime
	s.MicrosoftStoreForBusinessPortalSelection = decoded.MicrosoftStoreForBusinessPortalSelection
	s.MobileAppCategories = decoded.MobileAppCategories
	s.PolicySets = decoded.PolicySets
	s.SymantecCodeSigningCertificate = decoded.SymantecCodeSigningCertificate
	s.TargetedManagedAppConfigurations = decoded.TargetedManagedAppConfigurations
	s.VppTokens = decoded.VppTokens
	s.WdacSupplementalPolicies = decoded.WdacSupplementalPolicies
	s.WindowsInformationProtectionDeviceRegistrations = decoded.WindowsInformationProtectionDeviceRegistrations
	s.WindowsInformationProtectionPolicies = decoded.WindowsInformationProtectionPolicies
	s.WindowsInformationProtectionWipeActions = decoded.WindowsInformationProtectionWipeActions
	s.WindowsManagedAppProtections = decoded.WindowsManagedAppProtections
	s.WindowsManagementApp = decoded.WindowsManagementApp
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceAppManagement into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["deviceAppManagementTasks"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling DeviceAppManagementTasks into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceAppManagementTask, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceAppManagementTaskImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'DeviceAppManagementTasks' for 'DeviceAppManagement': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.DeviceAppManagementTasks = &output
	}

	if v, ok := temp["managedAppPolicies"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ManagedAppPolicies into list []json.RawMessage: %+v", err)
		}

		output := make([]ManagedAppPolicy, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalManagedAppPolicyImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ManagedAppPolicies' for 'DeviceAppManagement': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ManagedAppPolicies = &output
	}

	if v, ok := temp["managedAppRegistrations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ManagedAppRegistrations into list []json.RawMessage: %+v", err)
		}

		output := make([]ManagedAppRegistration, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalManagedAppRegistrationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ManagedAppRegistrations' for 'DeviceAppManagement': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ManagedAppRegistrations = &output
	}

	if v, ok := temp["managedAppStatuses"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ManagedAppStatuses into list []json.RawMessage: %+v", err)
		}

		output := make([]ManagedAppStatus, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalManagedAppStatusImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ManagedAppStatuses' for 'DeviceAppManagement': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ManagedAppStatuses = &output
	}

	if v, ok := temp["managedEBooks"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ManagedEBooks into list []json.RawMessage: %+v", err)
		}

		output := make([]ManagedEBook, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalManagedEBookImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ManagedEBooks' for 'DeviceAppManagement': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ManagedEBooks = &output
	}

	if v, ok := temp["mobileAppCatalogPackages"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling MobileAppCatalogPackages into list []json.RawMessage: %+v", err)
		}

		output := make([]MobileAppCatalogPackage, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalMobileAppCatalogPackageImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'MobileAppCatalogPackages' for 'DeviceAppManagement': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.MobileAppCatalogPackages = &output
	}

	if v, ok := temp["mobileAppConfigurations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling MobileAppConfigurations into list []json.RawMessage: %+v", err)
		}

		output := make([]ManagedDeviceMobileAppConfiguration, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalManagedDeviceMobileAppConfigurationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'MobileAppConfigurations' for 'DeviceAppManagement': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.MobileAppConfigurations = &output
	}

	if v, ok := temp["mobileApps"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling MobileApps into list []json.RawMessage: %+v", err)
		}

		output := make([]MobileApp, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalMobileAppImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'MobileApps' for 'DeviceAppManagement': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.MobileApps = &output
	}

	return nil
}
