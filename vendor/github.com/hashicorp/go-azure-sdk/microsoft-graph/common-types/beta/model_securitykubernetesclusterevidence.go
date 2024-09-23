package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityAlertEvidence = SecurityKubernetesClusterEvidence{}

type SecurityKubernetesClusterEvidence struct {
	// The cloud identifier of the cluster. Can be either an amazonResourceEvidence, azureResourceEvidence, or
	// googleCloudResourceEvidence object.
	CloudResource SecurityAlertEvidence `json:"cloudResource"`

	// The distribution type of the cluster.
	Distribution nullable.Type[string] `json:"distribution,omitempty"`

	// The cluster name.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The platform the cluster runs on. Possible values are: unknown, aks, eks, gke, arc, unknownFutureValue.
	Platform *SecurityKubernetesPlatform `json:"platform,omitempty"`

	// The kubernetes version of the cluster.
	Version nullable.Type[string] `json:"version,omitempty"`

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

func (s SecurityKubernetesClusterEvidence) SecurityAlertEvidence() BaseSecurityAlertEvidenceImpl {
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

var _ json.Marshaler = SecurityKubernetesClusterEvidence{}

func (s SecurityKubernetesClusterEvidence) MarshalJSON() ([]byte, error) {
	type wrapper SecurityKubernetesClusterEvidence
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityKubernetesClusterEvidence: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityKubernetesClusterEvidence: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.kubernetesClusterEvidence"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityKubernetesClusterEvidence: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecurityKubernetesClusterEvidence{}

func (s *SecurityKubernetesClusterEvidence) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Distribution             nullable.Type[string]              `json:"distribution,omitempty"`
		Name                     nullable.Type[string]              `json:"name,omitempty"`
		Platform                 *SecurityKubernetesPlatform        `json:"platform,omitempty"`
		Version                  nullable.Type[string]              `json:"version,omitempty"`
		CreatedDateTime          *string                            `json:"createdDateTime,omitempty"`
		DetailedRoles            *[]string                          `json:"detailedRoles,omitempty"`
		ODataId                  *string                            `json:"@odata.id,omitempty"`
		ODataType                *string                            `json:"@odata.type,omitempty"`
		RemediationStatus        *SecurityEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
		RemediationStatusDetails nullable.Type[string]              `json:"remediationStatusDetails,omitempty"`
		Roles                    *[]SecurityEvidenceRole            `json:"roles,omitempty"`
		Tags                     *[]string                          `json:"tags,omitempty"`
		Verdict                  *SecurityEvidenceVerdict           `json:"verdict,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Distribution = decoded.Distribution
	s.Name = decoded.Name
	s.Platform = decoded.Platform
	s.Version = decoded.Version
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DetailedRoles = decoded.DetailedRoles
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.RemediationStatus = decoded.RemediationStatus
	s.RemediationStatusDetails = decoded.RemediationStatusDetails
	s.Roles = decoded.Roles
	s.Tags = decoded.Tags
	s.Verdict = decoded.Verdict

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecurityKubernetesClusterEvidence into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["cloudResource"]; ok {
		impl, err := UnmarshalSecurityAlertEvidenceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CloudResource' for 'SecurityKubernetesClusterEvidence': %+v", err)
		}
		s.CloudResource = impl
	}

	return nil
}
