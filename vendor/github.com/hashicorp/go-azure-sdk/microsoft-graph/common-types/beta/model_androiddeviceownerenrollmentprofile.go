package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AndroidDeviceOwnerEnrollmentProfile{}

type AndroidDeviceOwnerEnrollmentProfile struct {
	// Tenant GUID the enrollment profile belongs to.
	AccountId nullable.Type[string] `json:"accountId,omitempty"`

	// Boolean that indicates that the Wi-Fi network should be configured during device provisioning. When set to TRUE,
	// device provisioning will use Wi-Fi related properties to automatically connect to Wi-Fi networks. When set to FALSE
	// or undefined, other Wi-Fi related properties will be ignored. Default value is TRUE. Returned by default.
	ConfigureWifi *bool `json:"configureWifi,omitempty"`

	// Date time the enrollment profile was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Description for the enrollment profile.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Indicates the device name template used for the enrolled Android devices. The maximum length allowed for this
	// property is 63 characters. The template expression contains normal text and tokens, including the serial number of
	// the device, user name, device type, upn prefix, or a randomly generated number. Supported Tokens for device name
	// templates are: (for device naming template expression): {{SERIAL}}, {{SERIALLAST4DIGITS}}, {{ENROLLMENTDATETIME}},
	// {{USERNAME}}, {{DEVICETYPE}}, {{UPNPREFIX}}, {{rand:x}}. Supports: $select, $top, $skip. $Search, $orderBy and
	// $filter are not supported.
	DeviceNameTemplate nullable.Type[string] `json:"deviceNameTemplate,omitempty"`

	// Display name for the enrollment profile.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Total number of Android devices that have enrolled using this enrollment profile.
	EnrolledDeviceCount *int64 `json:"enrolledDeviceCount,omitempty"`

	// The enrollment mode for an enrollment profile.
	EnrollmentMode *AndroidDeviceOwnerEnrollmentMode `json:"enrollmentMode,omitempty"`

	// The enrollment token type for an enrollment profile.
	EnrollmentTokenType *AndroidDeviceOwnerEnrollmentTokenType `json:"enrollmentTokenType,omitempty"`

	// Total number of AOSP devices that have enrolled using the current token. Valid values 0 to 20000
	EnrollmentTokenUsageCount *int64 `json:"enrollmentTokenUsageCount,omitempty"`

	// Boolean indicating if this profile is an Android AOSP for Teams device profile.
	IsTeamsDeviceProfile *bool `json:"isTeamsDeviceProfile,omitempty"`

	// Date time the enrollment profile was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// String used to generate a QR code for the token.
	QrCodeContent nullable.Type[string] `json:"qrCodeContent,omitempty"`

	// String used to generate a QR code for the token.
	QrCodeImage *MimeContent `json:"qrCodeImage,omitempty"`

	// List of Scope Tags for this Entity instance.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// Date time the most recently created token was created.
	TokenCreationDateTime *string `json:"tokenCreationDateTime,omitempty"`

	// Date time the most recently created token will expire.
	TokenExpirationDateTime *string `json:"tokenExpirationDateTime,omitempty"`

	// Value of the most recently created token for this enrollment profile.
	TokenValue nullable.Type[string] `json:"tokenValue,omitempty"`

	// Boolean that indicates if hidden wifi networks are enabled
	WifiHidden *bool `json:"wifiHidden,omitempty"`

	// String that contains the wi-fi login password
	WifiPassword nullable.Type[string] `json:"wifiPassword,omitempty"`

	// This enum represents Wi-Fi Security Types for Android Device Owner AOSP Scenarios.
	WifiSecurityType *AospWifiSecurityType `json:"wifiSecurityType,omitempty"`

	// String that contains the wi-fi login ssid
	WifiSsid nullable.Type[string] `json:"wifiSsid,omitempty"`

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

func (s AndroidDeviceOwnerEnrollmentProfile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AndroidDeviceOwnerEnrollmentProfile{}

func (s AndroidDeviceOwnerEnrollmentProfile) MarshalJSON() ([]byte, error) {
	type wrapper AndroidDeviceOwnerEnrollmentProfile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AndroidDeviceOwnerEnrollmentProfile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AndroidDeviceOwnerEnrollmentProfile: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.androidDeviceOwnerEnrollmentProfile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AndroidDeviceOwnerEnrollmentProfile: %+v", err)
	}

	return encoded, nil
}
