package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsUpdatesUpdatePolicy{}

type WindowsUpdatesUpdatePolicy struct {
	// Specifies the audience to target.
	Audience *WindowsUpdatesDeploymentAudience `json:"audience,omitempty"`

	// Rules for governing the automatic creation of compliance changes.
	ComplianceChangeRules *[]WindowsUpdatesComplianceChangeRule `json:"complianceChangeRules,omitempty"`

	// Compliance changes like content approvals which result in the automatic creation of deployments using the audience
	// and deploymentSettings of the policy.
	ComplianceChanges *[]WindowsUpdatesComplianceChange `json:"complianceChanges,omitempty"`

	// The date and time when the update policy was created. The Timestamp type represents date and time information using
	// ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Settings for governing how to deploy content.
	DeploymentSettings *WindowsUpdatesDeploymentSettings `json:"deploymentSettings,omitempty"`

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

func (s WindowsUpdatesUpdatePolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsUpdatesUpdatePolicy{}

func (s WindowsUpdatesUpdatePolicy) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesUpdatePolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesUpdatePolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesUpdatePolicy: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.updatePolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesUpdatePolicy: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &WindowsUpdatesUpdatePolicy{}

func (s *WindowsUpdatesUpdatePolicy) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Audience           *WindowsUpdatesDeploymentAudience `json:"audience,omitempty"`
		CreatedDateTime    nullable.Type[string]             `json:"createdDateTime,omitempty"`
		DeploymentSettings *WindowsUpdatesDeploymentSettings `json:"deploymentSettings,omitempty"`
		Id                 *string                           `json:"id,omitempty"`
		ODataId            *string                           `json:"@odata.id,omitempty"`
		ODataType          *string                           `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Audience = decoded.Audience
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DeploymentSettings = decoded.DeploymentSettings
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling WindowsUpdatesUpdatePolicy into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["complianceChangeRules"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ComplianceChangeRules into list []json.RawMessage: %+v", err)
		}

		output := make([]WindowsUpdatesComplianceChangeRule, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalWindowsUpdatesComplianceChangeRuleImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ComplianceChangeRules' for 'WindowsUpdatesUpdatePolicy': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ComplianceChangeRules = &output
	}

	if v, ok := temp["complianceChanges"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ComplianceChanges into list []json.RawMessage: %+v", err)
		}

		output := make([]WindowsUpdatesComplianceChange, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalWindowsUpdatesComplianceChangeImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ComplianceChanges' for 'WindowsUpdatesUpdatePolicy': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ComplianceChanges = &output
	}

	return nil
}
