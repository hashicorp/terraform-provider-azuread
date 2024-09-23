package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityAlertEvidence = SecurityProcessEvidence{}

type SecurityProcessEvidence struct {
	// The status of the detection.The possible values are: detected, blocked, prevented, unknownFutureValue.
	DetectionStatus *SecurityDetectionStatus `json:"detectionStatus,omitempty"`

	// Image file details.
	ImageFile *SecurityFileDetails `json:"imageFile,omitempty"`

	// A unique identifier assigned to a device by Microsoft Defender for Endpoint.
	MdeDeviceId nullable.Type[string] `json:"mdeDeviceId,omitempty"`

	// Date and time when the parent of the process was created. The DateTimeOffset type represents date and time
	// information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	ParentProcessCreationDateTime nullable.Type[string] `json:"parentProcessCreationDateTime,omitempty"`

	// Process ID (PID) of the parent process that spawned the process.
	ParentProcessId nullable.Type[int64] `json:"parentProcessId,omitempty"`

	// Parent process image file details.
	ParentProcessImageFile *SecurityFileDetails `json:"parentProcessImageFile,omitempty"`

	// Command line used to create the new process.
	ProcessCommandLine nullable.Type[string] `json:"processCommandLine,omitempty"`

	// Date and time when the process was created. The DateTimeOffset type represents date and time information using ISO
	// 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	ProcessCreationDateTime nullable.Type[string] `json:"processCreationDateTime,omitempty"`

	// Process ID (PID) of the newly created process.
	ProcessId nullable.Type[int64] `json:"processId,omitempty"`

	// User details of the user that ran the process.
	UserAccount *SecurityUserAccount `json:"userAccount,omitempty"`

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

func (s SecurityProcessEvidence) SecurityAlertEvidence() BaseSecurityAlertEvidenceImpl {
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

var _ json.Marshaler = SecurityProcessEvidence{}

func (s SecurityProcessEvidence) MarshalJSON() ([]byte, error) {
	type wrapper SecurityProcessEvidence
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityProcessEvidence: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityProcessEvidence: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.processEvidence"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityProcessEvidence: %+v", err)
	}

	return encoded, nil
}
