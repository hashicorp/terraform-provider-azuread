package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityAlertEvidence = SecurityCloudLogonSessionEvidence{}

type SecurityCloudLogonSessionEvidence struct {
	// The account associated with the sign-in session.
	Account *SecurityUserEvidence `json:"account,omitempty"`

	// The browser that is used for the sign-in, if known.
	Browser nullable.Type[string] `json:"browser,omitempty"`

	// The friendly name of the device, if known.
	DeviceName nullable.Type[string] `json:"deviceName,omitempty"`

	// The operating system that the device is running, if known.
	OperatingSystem nullable.Type[string] `json:"operatingSystem,omitempty"`

	// The previous sign-in time for this account, if known.
	PreviousLogonDateTime nullable.Type[string] `json:"previousLogonDateTime,omitempty"`

	// The authentication protocol that is used in this session, if known.
	Protocol nullable.Type[string] `json:"protocol,omitempty"`

	// The session ID for the account reported in the alert.
	SessionId nullable.Type[string] `json:"sessionId,omitempty"`

	// The session start time, if known.
	StartUtcDateTime nullable.Type[string] `json:"startUtcDateTime,omitempty"`

	// The user agent that is used for the sign-in, if known.
	UserAgent nullable.Type[string] `json:"userAgent,omitempty"`

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

func (s SecurityCloudLogonSessionEvidence) SecurityAlertEvidence() BaseSecurityAlertEvidenceImpl {
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

var _ json.Marshaler = SecurityCloudLogonSessionEvidence{}

func (s SecurityCloudLogonSessionEvidence) MarshalJSON() ([]byte, error) {
	type wrapper SecurityCloudLogonSessionEvidence
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityCloudLogonSessionEvidence: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityCloudLogonSessionEvidence: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.cloudLogonSessionEvidence"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityCloudLogonSessionEvidence: %+v", err)
	}

	return encoded, nil
}
