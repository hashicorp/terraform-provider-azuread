package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = GroupPolicyDefinition{}

type GroupPolicyDefinition struct {
	// The group policy category associated with the definition.
	Category *GroupPolicyCategory `json:"category,omitempty"`

	// The localized full category path for the policy.
	CategoryPath nullable.Type[string] `json:"categoryPath,omitempty"`

	// Group Policy Definition Class Type.
	ClassType *GroupPolicyDefinitionClassType `json:"classType,omitempty"`

	// The group policy file associated with the definition.
	DefinitionFile *GroupPolicyDefinitionFile `json:"definitionFile,omitempty"`

	// The localized policy name.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The localized explanation or help text associated with the policy. The default value is empty.
	ExplainText nullable.Type[string] `json:"explainText,omitempty"`

	// The category id of the parent category
	GroupPolicyCategoryId *string `json:"groupPolicyCategoryId,omitempty"`

	// Signifies whether or not there are related definitions to this definition
	HasRelatedDefinitions *bool `json:"hasRelatedDefinitions,omitempty"`

	// The date and time the entity was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Minimum required CSP version for device configuration in this definition
	MinDeviceCspVersion nullable.Type[string] `json:"minDeviceCspVersion,omitempty"`

	// Minimum required CSP version for user configuration in this definition
	MinUserCspVersion nullable.Type[string] `json:"minUserCspVersion,omitempty"`

	// Definition of the next version of this definition
	NextVersionDefinition *GroupPolicyDefinition `json:"nextVersionDefinition,omitempty"`

	// Type of Group Policy File or Definition.
	PolicyType *GroupPolicyType `json:"policyType,omitempty"`

	// The group policy presentations associated with the definition.
	Presentations *[]GroupPolicyPresentation `json:"presentations,omitempty"`

	// Definition of the previous version of this definition
	PreviousVersionDefinition *GroupPolicyDefinition `json:"previousVersionDefinition,omitempty"`

	// Localized string used to specify what operating system or application version is affected by the policy.
	SupportedOn nullable.Type[string] `json:"supportedOn,omitempty"`

	// Setting definition version
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

func (s GroupPolicyDefinition) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = GroupPolicyDefinition{}

func (s GroupPolicyDefinition) MarshalJSON() ([]byte, error) {
	type wrapper GroupPolicyDefinition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GroupPolicyDefinition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GroupPolicyDefinition: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.groupPolicyDefinition"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GroupPolicyDefinition: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &GroupPolicyDefinition{}

func (s *GroupPolicyDefinition) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Category                  *GroupPolicyCategory            `json:"category,omitempty"`
		CategoryPath              nullable.Type[string]           `json:"categoryPath,omitempty"`
		ClassType                 *GroupPolicyDefinitionClassType `json:"classType,omitempty"`
		DisplayName               nullable.Type[string]           `json:"displayName,omitempty"`
		ExplainText               nullable.Type[string]           `json:"explainText,omitempty"`
		GroupPolicyCategoryId     *string                         `json:"groupPolicyCategoryId,omitempty"`
		HasRelatedDefinitions     *bool                           `json:"hasRelatedDefinitions,omitempty"`
		LastModifiedDateTime      *string                         `json:"lastModifiedDateTime,omitempty"`
		MinDeviceCspVersion       nullable.Type[string]           `json:"minDeviceCspVersion,omitempty"`
		MinUserCspVersion         nullable.Type[string]           `json:"minUserCspVersion,omitempty"`
		NextVersionDefinition     *GroupPolicyDefinition          `json:"nextVersionDefinition,omitempty"`
		PolicyType                *GroupPolicyType                `json:"policyType,omitempty"`
		PreviousVersionDefinition *GroupPolicyDefinition          `json:"previousVersionDefinition,omitempty"`
		SupportedOn               nullable.Type[string]           `json:"supportedOn,omitempty"`
		Version                   nullable.Type[string]           `json:"version,omitempty"`
		Id                        *string                         `json:"id,omitempty"`
		ODataId                   *string                         `json:"@odata.id,omitempty"`
		ODataType                 *string                         `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Category = decoded.Category
	s.CategoryPath = decoded.CategoryPath
	s.ClassType = decoded.ClassType
	s.DisplayName = decoded.DisplayName
	s.ExplainText = decoded.ExplainText
	s.GroupPolicyCategoryId = decoded.GroupPolicyCategoryId
	s.HasRelatedDefinitions = decoded.HasRelatedDefinitions
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.MinDeviceCspVersion = decoded.MinDeviceCspVersion
	s.MinUserCspVersion = decoded.MinUserCspVersion
	s.NextVersionDefinition = decoded.NextVersionDefinition
	s.PolicyType = decoded.PolicyType
	s.PreviousVersionDefinition = decoded.PreviousVersionDefinition
	s.SupportedOn = decoded.SupportedOn
	s.Version = decoded.Version
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling GroupPolicyDefinition into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["definitionFile"]; ok {
		impl, err := UnmarshalGroupPolicyDefinitionFileImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'DefinitionFile' for 'GroupPolicyDefinition': %+v", err)
		}
		s.DefinitionFile = &impl
	}

	if v, ok := temp["presentations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Presentations into list []json.RawMessage: %+v", err)
		}

		output := make([]GroupPolicyPresentation, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalGroupPolicyPresentationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Presentations' for 'GroupPolicyDefinition': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Presentations = &output
	}

	return nil
}
