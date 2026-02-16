package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityAlertEvidence = SecurityDeviceEvidence{}

type SecurityDeviceEvidence struct {
	// A unique identifier assigned to a device by Microsoft Entra ID when device is Microsoft Entra joined.
	AzureAdDeviceId nullable.Type[string] `json:"azureAdDeviceId,omitempty"`

	// State of the Defender AntiMalware engine. The possible values are: notReporting, disabled, notUpdated, updated,
	// unknown, notSupported, unknownFutureValue.
	DefenderAvStatus *SecurityDefenderAvStatus `json:"defenderAvStatus,omitempty"`

	// The fully qualified domain name (FQDN) for the device.
	DeviceDnsName nullable.Type[string] `json:"deviceDnsName,omitempty"`

	// The DNS domain that this computer belongs to. A sequence of labels separated by dots.
	DnsDomain nullable.Type[string] `json:"dnsDomain,omitempty"`

	// The date and time when the device was first seen.
	FirstSeenDateTime nullable.Type[string] `json:"firstSeenDateTime,omitempty"`

	// The health state of the device. The possible values are: active, inactive, impairedCommunication, noSensorData,
	// noSensorDataImpairedCommunication, unknown, unknownFutureValue.
	HealthStatus *SecurityDeviceHealthStatus `json:"healthStatus,omitempty"`

	// The hostname without the domain suffix.
	HostName nullable.Type[string] `json:"hostName,omitempty"`

	// Ip interfaces of the device during the time of the alert.
	IPInterfaces *[]string `json:"ipInterfaces,omitempty"`

	LastExternalIPAddress nullable.Type[string] `json:"lastExternalIpAddress,omitempty"`
	LastIPAddress         nullable.Type[string] `json:"lastIpAddress,omitempty"`

	// Users that were logged on the machine during the time of the alert.
	LoggedOnUsers *[]SecurityLoggedOnUser `json:"loggedOnUsers,omitempty"`

	// A unique identifier assigned to a device by Microsoft Defender for Endpoint.
	MdeDeviceId nullable.Type[string] `json:"mdeDeviceId,omitempty"`

	// A logical grouping of computers within a Microsoft Windows network.
	NtDomain nullable.Type[string] `json:"ntDomain,omitempty"`

	// The status of the machine onboarding to Microsoft Defender for Endpoint. The possible values are: insufficientInfo,
	// onboarded, canBeOnboarded, unsupported, unknownFutureValue.
	OnboardingStatus *SecurityOnboardingStatus `json:"onboardingStatus,omitempty"`

	// The build version for the operating system the device is running.
	OsBuild nullable.Type[int64] `json:"osBuild,omitempty"`

	// The operating system platform the device is running.
	OsPlatform nullable.Type[string] `json:"osPlatform,omitempty"`

	// The ID of the role-based access control device group.
	RbacGroupId nullable.Type[int64] `json:"rbacGroupId,omitempty"`

	// The name of the role-based access control device group.
	RbacGroupName nullable.Type[string] `json:"rbacGroupName,omitempty"`

	// Risk score as evaluated by Microsoft Defender for Endpoint. The possible values are: none, informational, low,
	// medium, high, unknownFutureValue.
	RiskScore *SecurityDeviceRiskScore `json:"riskScore,omitempty"`

	// The version of the operating system platform.
	Version nullable.Type[string] `json:"version,omitempty"`

	// Metadata of the virtual machine (VM) on which Microsoft Defender for Endpoint is running.
	VmMetadata *SecurityVmMetadata `json:"vmMetadata,omitempty"`

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

func (s SecurityDeviceEvidence) SecurityAlertEvidence() BaseSecurityAlertEvidenceImpl {
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

var _ json.Marshaler = SecurityDeviceEvidence{}

func (s SecurityDeviceEvidence) MarshalJSON() ([]byte, error) {
	type wrapper SecurityDeviceEvidence
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityDeviceEvidence: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityDeviceEvidence: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.deviceEvidence"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityDeviceEvidence: %+v", err)
	}

	return encoded, nil
}
