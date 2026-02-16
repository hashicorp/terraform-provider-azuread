package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = MobileThreatDefenseConnector{}

type MobileThreatDefenseConnector struct {
	// When TRUE, indicates the Mobile Threat Defense partner may collect metadata about installed applications from Intune
	// for IOS devices. When FALSE, indicates the Mobile Threat Defense partner may not collect metadata about installed
	// applications from Intune for IOS devices. Default value is FALSE.
	AllowPartnerToCollectIOSApplicationMetadata *bool `json:"allowPartnerToCollectIOSApplicationMetadata,omitempty"`

	// When TRUE, indicates the Mobile Threat Defense partner may collect metadata about personally installed applications
	// from Intune for IOS devices. When FALSE, indicates the Mobile Threat Defense partner may not collect metadata about
	// personally installed applications from Intune for IOS devices. Default value is FALSE.
	AllowPartnerToCollectIOSPersonalApplicationMetadata *bool `json:"allowPartnerToCollectIOSPersonalApplicationMetadata,omitempty"`

	// For Android, set whether Intune must receive data from the Mobile Threat Defense partner prior to marking a device
	// compliant
	AndroidDeviceBlockedOnMissingPartnerData *bool `json:"androidDeviceBlockedOnMissingPartnerData,omitempty"`

	// For Android, set whether data from the Mobile Threat Defense partner should be used during compliance evaluations
	AndroidEnabled *bool `json:"androidEnabled,omitempty"`

	// When TRUE, inidicates that data from the Mobile Threat Defense partner can be used during Mobile Application
	// Management (MAM) evaluations for Android devices. When FALSE, inidicates that data from the Mobile Threat Defense
	// partner should not be used during Mobile Application Management (MAM) evaluations for Android devices. Only one
	// partner per platform may be enabled for Mobile Application Management (MAM) evaluation. Default value is FALSE.
	AndroidMobileApplicationManagementEnabled *bool `json:"androidMobileApplicationManagementEnabled,omitempty"`

	// For IOS, set whether Intune must receive data from the Mobile Threat Defense partner prior to marking a device
	// compliant
	IosDeviceBlockedOnMissingPartnerData *bool `json:"iosDeviceBlockedOnMissingPartnerData,omitempty"`

	// For IOS, get or set whether data from the Mobile Threat Defense partner should be used during compliance evaluations
	IosEnabled *bool `json:"iosEnabled,omitempty"`

	// When TRUE, inidicates that data from the Mobile Threat Defense partner can be used during Mobile Application
	// Management (MAM) evaluations for IOS devices. When FALSE, inidicates that data from the Mobile Threat Defense partner
	// should not be used during Mobile Application Management (MAM) evaluations for IOS devices. Only one partner per
	// platform may be enabled for Mobile Application Management (MAM) evaluation. Default value is FALSE.
	IosMobileApplicationManagementEnabled *bool `json:"iosMobileApplicationManagementEnabled,omitempty"`

	// DateTime of last Heartbeat recieved from the Mobile Threat Defense partner
	LastHeartbeatDateTime *string `json:"lastHeartbeatDateTime,omitempty"`

	// When TRUE, inidicates that configuration profile management via Microsoft Defender for Endpoint is enabled. When
	// FALSE, inidicates that configuration profile management via Microsoft Defender for Endpoint is disabled. Default
	// value is FALSE.
	MicrosoftDefenderForEndpointAttachEnabled *bool `json:"microsoftDefenderForEndpointAttachEnabled,omitempty"`

	// Partner state of this tenant.
	PartnerState *MobileThreatPartnerTenantState `json:"partnerState,omitempty"`

	// Get or Set days the per tenant tolerance to unresponsiveness for this partner integration
	PartnerUnresponsivenessThresholdInDays *int64 `json:"partnerUnresponsivenessThresholdInDays,omitempty"`

	// Get or set whether to block devices on the enabled platforms that do not meet the minimum version requirements of the
	// Mobile Threat Defense partner
	PartnerUnsupportedOsVersionBlocked *bool `json:"partnerUnsupportedOsVersionBlocked,omitempty"`

	// When TRUE, inidicates that Intune must receive data from the Mobile Threat Defense partner prior to marking a device
	// compliant for Windows. When FALSE, inidicates that Intune may make a device compliant without receiving data from the
	// Mobile Threat Defense partner for Windows. Default value is FALSE.
	WindowsDeviceBlockedOnMissingPartnerData *bool `json:"windowsDeviceBlockedOnMissingPartnerData,omitempty"`

	// When TRUE, inidicates that data from the Mobile Threat Defense partner can be used during compliance evaluations for
	// Windows. When FALSE, inidicates that data from the Mobile Threat Defense partner should not be used during compliance
	// evaluations for Windows. Default value is FALSE.
	WindowsEnabled *bool `json:"windowsEnabled,omitempty"`

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

func (s MobileThreatDefenseConnector) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MobileThreatDefenseConnector{}

func (s MobileThreatDefenseConnector) MarshalJSON() ([]byte, error) {
	type wrapper MobileThreatDefenseConnector
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MobileThreatDefenseConnector: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MobileThreatDefenseConnector: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.mobileThreatDefenseConnector"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MobileThreatDefenseConnector: %+v", err)
	}

	return encoded, nil
}
