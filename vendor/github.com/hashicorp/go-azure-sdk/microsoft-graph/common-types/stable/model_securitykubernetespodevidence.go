package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityAlertEvidence = SecurityKubernetesPodEvidence{}

type SecurityKubernetesPodEvidence struct {
	// The list of pod containers which are not init or ephemeral containers.
	Containers *[]SecurityContainerEvidence `json:"containers,omitempty"`

	// The pod controller.
	Controller *SecurityKubernetesControllerEvidence `json:"controller,omitempty"`

	// The list of pod ephemeral containers.
	EphemeralContainers *[]SecurityContainerEvidence `json:"ephemeralContainers,omitempty"`

	// The list of pod init containers.
	InitContainers *[]SecurityContainerEvidence `json:"initContainers,omitempty"`

	// The pod labels.
	Labels *SecurityDictionary `json:"labels,omitempty"`

	// The pod name.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The pod namespace.
	Namespace *SecurityKubernetesNamespaceEvidence `json:"namespace,omitempty"`

	// The pod IP.
	PodIp *SecurityIPEvidence `json:"podIp,omitempty"`

	// The pod service account.
	ServiceAccount *SecurityKubernetesServiceAccountEvidence `json:"serviceAccount,omitempty"`

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

func (s SecurityKubernetesPodEvidence) SecurityAlertEvidence() BaseSecurityAlertEvidenceImpl {
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

var _ json.Marshaler = SecurityKubernetesPodEvidence{}

func (s SecurityKubernetesPodEvidence) MarshalJSON() ([]byte, error) {
	type wrapper SecurityKubernetesPodEvidence
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityKubernetesPodEvidence: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityKubernetesPodEvidence: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.kubernetesPodEvidence"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityKubernetesPodEvidence: %+v", err)
	}

	return encoded, nil
}
