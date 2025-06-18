package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = MobileThreatDefenseConnector{}

type MobileThreatDefenseConnector struct {
	// When TRUE, indicates the Mobile Threat Defense partner may collect metadata about installed applications from Intune
	// for iOS devices. When FALSE, indicates the Mobile Threat Defense partner may not collect metadata about installed
	// applications from Intune for iOS devices. Default value is FALSE.
	AllowPartnerToCollectIOSApplicationMetadata *bool `json:"allowPartnerToCollectIOSApplicationMetadata,omitempty"`

	// When TRUE, indicates the Mobile Threat Defense partner may collect metadata about personally installed applications
	// from Intune for iOS devices. When FALSE, indicates the Mobile Threat Defense partner may not collect metadata about
	// personally installed applications from Intune for iOS devices. Default value is FALSE.
	AllowPartnerToCollectIOSPersonalApplicationMetadata *bool `json:"allowPartnerToCollectIOSPersonalApplicationMetadata,omitempty"`

	// When TRUE, allows the Mobile Threat Defense partner to request a list of installed certificates on iOS/iPadOS devices
	// from Intune to use for threat analysis. This list of installed certificates will be sent from enrolled iOS/iPadOS
	// devices and will include unmanaged certificates (certificates not deployed through Intune). When FALSE, indicates
	// that metadata about installed certificates will not be collected. Default value is FALSE.
	AllowPartnerToCollectIosCertificateMetadata *bool `json:"allowPartnerToCollectIosCertificateMetadata,omitempty"`

	// When TRUE, allows the Mobile Threat Defense partner to request a list of installed certificates on personally owned
	// iOS/iPadOS devices from Intune to use for threat analysis. This list of installed certificates will be sent from
	// enrolled personally owned iOS/iPadOS devices and will include unmanaged certificates (certificates not deployed
	// through Intune). When FALSE, no metadata for installed certificates is sent for personally owned iOS/iPadOS devices.
	// Default value is FALSE.
	AllowPartnerToCollectIosPersonalCertificateMetadata *bool `json:"allowPartnerToCollectIosPersonalCertificateMetadata,omitempty"`

	// When TRUE, indicates that Intune must receive data from the Mobile Threat Defense partner prior to marking an Android
	// device compliant. When FALSE, indicates that Intune may mark an Android device compliant before receiving data from
	// the Mobile Threat Defense partner.
	AndroidDeviceBlockedOnMissingPartnerData *bool `json:"androidDeviceBlockedOnMissingPartnerData,omitempty"`

	// When TRUE, indicates that data from the Mobile Threat Defense partner will be used during compliance evaluations for
	// Android devices. When FALSE, indicates that data from the Mobile Threat Defense partner will not be used during
	// compliance evaluations for Android devices. Default value is FALSE.
	AndroidEnabled *bool `json:"androidEnabled,omitempty"`

	// When TRUE, inidicates that data from the Mobile Threat Defense partner can be used during Mobile Application
	// Management (MAM) evaluations for Android devices. When FALSE, inidicates that data from the Mobile Threat Defense
	// partner should not be used during Mobile Application Management (MAM) evaluations for Android devices. Only one
	// partner per platform may be enabled for Mobile Application Management (MAM) evaluation. Default value is FALSE.
	AndroidMobileApplicationManagementEnabled *bool `json:"androidMobileApplicationManagementEnabled,omitempty"`

	// When TRUE, indicates that Intune must receive data from the Mobile Threat Defense partner prior to marking a device
	// compliant. When FALSE, indicates that Intune may not recieve data from Mobile Threat Defense partner prior to making
	// device compliant. Default value is FALSE.
	IosDeviceBlockedOnMissingPartnerData *bool `json:"iosDeviceBlockedOnMissingPartnerData,omitempty"`

	// When TRUE, indicates that data from the Mobile Threat Defense partner will be used during compliance evaluations for
	// iOS devices. When FALSE, indicates that data from the Mobile Threat Defense partner will not be used during
	// compliance evaluations for iOS devices. Default value is FALSE.
	IosEnabled *bool `json:"iosEnabled,omitempty"`

	// When TRUE, inidicates that data from the Mobile Threat Defense partner can be used during Mobile Application
	// Management (MAM) evaluations for iOS devices. When FALSE, inidicates that data from the Mobile Threat Defense partner
	// should not be used during Mobile Application Management (MAM) evaluations for iOS devices. Only one partner per
	// platform may be enabled for Mobile Application Management (MAM) evaluation. Default value is FALSE.
	IosMobileApplicationManagementEnabled *bool `json:"iosMobileApplicationManagementEnabled,omitempty"`

	// DateTime of last Heartbeat recieved from the Mobile Threat Defense partner
	LastHeartbeatDateTime *string `json:"lastHeartbeatDateTime,omitempty"`

	// When TRUE, indicates that Intune must receive data from the Mobile Threat Defense partner prior to marking a Mac
	// device compliant. When FALSE, indicates that Intune may mark a Mac device compliant prior to receiving data from the
	// Mobile Threat Defense partner. Default value is FALSE.
	MacDeviceBlockedOnMissingPartnerData *bool `json:"macDeviceBlockedOnMissingPartnerData,omitempty"`

	// When TRUE, indicates that data from the Mobile Threat Defense partner will be used during compliance evaluations for
	// Mac devices. When FALSE, indicates that data from the Mobile Threat Defense partner will not be used during
	// compliance evaluations for Mac devices. Default value is FALSE.
	MacEnabled *bool `json:"macEnabled,omitempty"`

	// When TRUE, inidicates that configuration profile management via Microsoft Defender for Endpoint is enabled. When
	// FALSE, inidicates that configuration profile management via Microsoft Defender for Endpoint is disabled. Default
	// value is FALSE.
	MicrosoftDefenderForEndpointAttachEnabled *bool `json:"microsoftDefenderForEndpointAttachEnabled,omitempty"`

	// Partner state of this tenant.
	PartnerState *MobileThreatPartnerTenantState `json:"partnerState,omitempty"`

	// Indicates the number of days without receiving a heartbeat from a Mobile Threat Defense partner before the partner is
	// marked as unresponsive. Intune will the ignore the data from this Mobile Threat Defense Partner for next compliance
	// calculation.
	PartnerUnresponsivenessThresholdInDays *int64 `json:"partnerUnresponsivenessThresholdInDays,omitempty"`

	// When TRUE, indicates that Intune will mark devices noncompliant on enabled platforms that do not meet the minimum
	// version requirements of the Mobile Threat Defense partner. When FALSE, indicates that Intune will not mark devices
	// noncompliant on enabled platforms that do not meet the minimum version requirements of the Mobile Threat Defense
	// partner. Default value is FALSE.
	PartnerUnsupportedOsVersionBlocked *bool `json:"partnerUnsupportedOsVersionBlocked,omitempty"`

	// When TRUE, indicates that Intune must receive data from the data sync partner prior to marking a device compliant for
	// Windows. When FALSE, indicates that Intune may mark a device compliant without receiving data from the data sync
	// partner for Windows. Default value is FALSE.
	WindowsDeviceBlockedOnMissingPartnerData *bool `json:"windowsDeviceBlockedOnMissingPartnerData,omitempty"`

	// When TRUE, indicates that data from the Mobile Threat Defense partner will be used during compliance evaluations for
	// Windows. When FALSE, indicates that data from the Mobile Threat Defense partner will not be used during compliance
	// evaluations for Windows. Default value is FALSE.
	WindowsEnabled *bool `json:"windowsEnabled,omitempty"`

	// When TRUE, inidicates that data from the Mobile Threat Defense partner can be used during Mobile Application
	// Management (MAM) evaluations for iOS devices. When FALSE, inidicates that data from the Mobile Threat Defense partner
	// should not be used during Mobile Application Management (MAM) evaluations for iOS devices. Only one partner per
	// platform may be enabled for Mobile Application Management (MAM) evaluation. Default value is FALSE.
	WindowsMobileApplicationManagementEnabled *bool `json:"windowsMobileApplicationManagementEnabled,omitempty"`

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
