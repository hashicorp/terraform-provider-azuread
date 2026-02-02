package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsUpdatesComplianceChange = WindowsUpdatesContentApproval{}

type WindowsUpdatesContentApproval struct {
	Content WindowsUpdatesDeployableContent `json:"content"`

	// Settings for governing how to deploy content.
	DeploymentSettings *WindowsUpdatesDeploymentSettings `json:"deploymentSettings,omitempty"`

	// Deployments created as a result of applying the approval.
	Deployments *[]WindowsUpdatesDeployment `json:"deployments,omitempty"`

	// Fields inherited from WindowsUpdatesComplianceChange

	// The date and time when a compliance change was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// True indicates that a compliance change is revoked, preventing further application. Revoking a compliance change is a
	// final action.
	IsRevoked nullable.Type[bool] `json:"isRevoked,omitempty"`

	// The date and time when the compliance change was revoked.
	RevokedDateTime nullable.Type[string] `json:"revokedDateTime,omitempty"`

	// The policy this compliance change is a member of.
	UpdatePolicy *WindowsUpdatesUpdatePolicy `json:"updatePolicy,omitempty"`

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

func (s WindowsUpdatesContentApproval) WindowsUpdatesComplianceChange() BaseWindowsUpdatesComplianceChangeImpl {
	return BaseWindowsUpdatesComplianceChangeImpl{
		CreatedDateTime: s.CreatedDateTime,
		IsRevoked:       s.IsRevoked,
		RevokedDateTime: s.RevokedDateTime,
		UpdatePolicy:    s.UpdatePolicy,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s WindowsUpdatesContentApproval) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsUpdatesContentApproval{}

func (s WindowsUpdatesContentApproval) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesContentApproval
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesContentApproval: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesContentApproval: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.contentApproval"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesContentApproval: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &WindowsUpdatesContentApproval{}

func (s *WindowsUpdatesContentApproval) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DeploymentSettings *WindowsUpdatesDeploymentSettings `json:"deploymentSettings,omitempty"`
		Deployments        *[]WindowsUpdatesDeployment       `json:"deployments,omitempty"`
		CreatedDateTime    nullable.Type[string]             `json:"createdDateTime,omitempty"`
		IsRevoked          nullable.Type[bool]               `json:"isRevoked,omitempty"`
		RevokedDateTime    nullable.Type[string]             `json:"revokedDateTime,omitempty"`
		UpdatePolicy       *WindowsUpdatesUpdatePolicy       `json:"updatePolicy,omitempty"`
		Id                 *string                           `json:"id,omitempty"`
		ODataId            *string                           `json:"@odata.id,omitempty"`
		ODataType          *string                           `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DeploymentSettings = decoded.DeploymentSettings
	s.Deployments = decoded.Deployments
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Id = decoded.Id
	s.IsRevoked = decoded.IsRevoked
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.RevokedDateTime = decoded.RevokedDateTime
	s.UpdatePolicy = decoded.UpdatePolicy

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling WindowsUpdatesContentApproval into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["content"]; ok {
		impl, err := UnmarshalWindowsUpdatesDeployableContentImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Content' for 'WindowsUpdatesContentApproval': %+v", err)
		}
		s.Content = impl
	}

	return nil
}
