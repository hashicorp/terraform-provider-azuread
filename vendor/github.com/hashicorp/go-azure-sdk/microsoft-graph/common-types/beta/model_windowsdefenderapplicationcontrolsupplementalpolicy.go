package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsDefenderApplicationControlSupplementalPolicy{}

type WindowsDefenderApplicationControlSupplementalPolicy struct {
	// The associated group assignments for the Windows Defender Application Control Supplemental Policy.
	Assignments *[]WindowsDefenderApplicationControlSupplementalPolicyAssignment `json:"assignments,omitempty"`

	// Indicates the content of the Windows Defender Application Control Supplemental Policy in byte array format.
	Content nullable.Type[string] `json:"content,omitempty"`

	// Indicates the file name associated with the content of the Windows Defender Application Control Supplemental Policy.
	ContentFileName nullable.Type[string] `json:"contentFileName,omitempty"`

	// Indicates the created date and time when the Windows Defender Application Control Supplemental Policy was uploaded.
	CreationDateTime *string `json:"creationDateTime,omitempty"`

	// WindowsDefenderApplicationControl supplemental policy deployment summary.
	DeploySummary *WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary `json:"deploySummary,omitempty"`

	// The description of the Windows Defender Application Control Supplemental Policy.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The list of device deployment states for this WindowsDefenderApplicationControl supplemental policy.
	DeviceStatuses *[]WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus `json:"deviceStatuses,omitempty"`

	// The display name of the Windows Defender Application Control Supplemental Policy.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Indicates the last modified date and time of the Windows Defender Application Control Supplemental Policy.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// List of Scope Tags for the Windows Defender Application Control Supplemental Policy entity.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// Indicates the Windows Defender Application Control Supplemental Policy's version.
	Version nullable.Type[string] `json:"version,omitempty"`

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

func (s WindowsDefenderApplicationControlSupplementalPolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsDefenderApplicationControlSupplementalPolicy{}

func (s WindowsDefenderApplicationControlSupplementalPolicy) MarshalJSON() ([]byte, error) {
	type wrapper WindowsDefenderApplicationControlSupplementalPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsDefenderApplicationControlSupplementalPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsDefenderApplicationControlSupplementalPolicy: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsDefenderApplicationControlSupplementalPolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsDefenderApplicationControlSupplementalPolicy: %+v", err)
	}

	return encoded, nil
}
