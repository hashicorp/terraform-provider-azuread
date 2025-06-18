package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityAlertEvidence = SecurityIoTDeviceEvidence{}

type SecurityIoTDeviceEvidence struct {
	// The device ID.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// The friendly name of the device.
	DeviceName nullable.Type[string] `json:"deviceName,omitempty"`

	// The URL to the device page in the IoT Defender portal.
	DevicePageLink nullable.Type[string] `json:"devicePageLink,omitempty"`

	// The device subtype.
	DeviceSubType nullable.Type[string] `json:"deviceSubType,omitempty"`

	// The type of the device. For example, 'temperature sensor,' 'freezer,' 'wind turbine,' and so on.
	DeviceType nullable.Type[string] `json:"deviceType,omitempty"`

	// The current IP address of the device.
	IPAddress *SecurityIPEvidence `json:"ipAddress,omitempty"`

	// The importance level for the IoT device. Possible values are low, normal, high, and unknownFutureValue.
	Importance *SecurityIoTDeviceImportanceType `json:"importance,omitempty"`

	// The azureResourceEvidence entity that represents the IoT Hub that the device belongs to.
	IoTHub *SecurityAzureResourceEvidence `json:"ioTHub,omitempty"`

	// The ID of the Azure Security Center for the IoT agent that is running on the device.
	IoTSecurityAgentId nullable.Type[string] `json:"ioTSecurityAgentId,omitempty"`

	// Indicates whether the device classified as an authorized device.
	IsAuthorized nullable.Type[bool] `json:"isAuthorized,omitempty"`

	// Indicates whether the device classified as a programming device.
	IsProgramming nullable.Type[bool] `json:"isProgramming,omitempty"`

	// Indicates whether the device classified as a scanner.
	IsScanner nullable.Type[bool] `json:"isScanner,omitempty"`

	// The MAC address of the device.
	MacAddress nullable.Type[string] `json:"macAddress,omitempty"`

	// The manufacturer of the device.
	Manufacturer nullable.Type[string] `json:"manufacturer,omitempty"`

	// The model of the device.
	Model nullable.Type[string] `json:"model,omitempty"`

	// The current network interface controllers on the device.
	Nics *[]SecurityNicEvidence `json:"nics,omitempty"`

	// The operating system the device is running.
	OperatingSystem nullable.Type[string] `json:"operatingSystem,omitempty"`

	// The owners for the device.
	Owners *[]string `json:"owners,omitempty"`

	// The list of protocols that the device supports.
	Protocols *[]string `json:"protocols,omitempty"`

	// The Purdue Layer of the device.
	PurdueLayer nullable.Type[string] `json:"purdueLayer,omitempty"`

	// The sensor that monitors the device.
	Sensor nullable.Type[string] `json:"sensor,omitempty"`

	// The serial number of the device.
	SerialNumber nullable.Type[string] `json:"serialNumber,omitempty"`

	// The site location of the device.
	Site nullable.Type[string] `json:"site,omitempty"`

	// The source (microsoft/vendor) of the device entity.
	Source nullable.Type[string] `json:"source,omitempty"`

	// A URL reference to the source item where the device is managed.
	SourceRef *SecurityUrlEvidence `json:"sourceRef,omitempty"`

	// The zone location of the device within a site.
	Zone nullable.Type[string] `json:"zone,omitempty"`

	// Fields inherited from SecurityAlertEvidence

	// The date and time when the evidence was created and added to the alert. The Timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Detailed description of the entity role/s in an alert. Values are free-form.
	DetailedRoles *[]string `json:"detailedRoles,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	RemediationStatus *SecurityEvidenceRemediationStatus `json:"remediationStatus,omitempty"`

	// Details about the remediation status.
	RemediationStatusDetails nullable.Type[string] `json:"remediationStatusDetails,omitempty"`

	// The role/s that an evidence entity represents in an alert, for example, an IP address that is associated with an
	// attacker has the evidence role Attacker.
	Roles *[]SecurityEvidenceRole `json:"roles,omitempty"`

	// Array of custom tags associated with an evidence instance, for example, to denote a group of devices, high-value
	// assets, etc.
	Tags *[]string `json:"tags,omitempty"`

	Verdict *SecurityEvidenceVerdict `json:"verdict,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s SecurityIoTDeviceEvidence) SecurityAlertEvidence() BaseSecurityAlertEvidenceImpl {
	return BaseSecurityAlertEvidenceImpl{
		CreatedDateTime:          s.CreatedDateTime,
		DetailedRoles:            s.DetailedRoles,
		ODataId:                  s.ODataId,
		ODataType:                s.ODataType,
		RemediationStatus:        s.RemediationStatus,
		RemediationStatusDetails: s.RemediationStatusDetails,
		Roles:                    s.Roles,
		Tags:                     s.Tags,
		Verdict:                  s.Verdict,
	}
}

var _ json.Marshaler = SecurityIoTDeviceEvidence{}

func (s SecurityIoTDeviceEvidence) MarshalJSON() ([]byte, error) {
	type wrapper SecurityIoTDeviceEvidence
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityIoTDeviceEvidence: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityIoTDeviceEvidence: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.ioTDeviceEvidence"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityIoTDeviceEvidence: %+v", err)
	}

	return encoded, nil
}
