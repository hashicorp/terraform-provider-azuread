package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyDefinitionFile interface {
	Entity
	GroupPolicyDefinitionFile() BaseGroupPolicyDefinitionFileImpl
}

var _ GroupPolicyDefinitionFile = BaseGroupPolicyDefinitionFileImpl{}

type BaseGroupPolicyDefinitionFileImpl struct {
	// The group policy definitions associated with the file.
	Definitions *[]GroupPolicyDefinition `json:"definitions,omitempty"`

	// The localized description of the policy settings in the ADMX file. The default value is empty.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The localized friendly name of the ADMX file.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The file name of the ADMX file without the path. For example: edge.admx
	FileName nullable.Type[string] `json:"fileName,omitempty"`

	// The supported language codes for the ADMX file.
	LanguageCodes *[]string `json:"languageCodes,omitempty"`

	// The date and time the entity was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Type of Group Policy File or Definition.
	PolicyType *GroupPolicyType `json:"policyType,omitempty"`

	// The revision version associated with the file.
	Revision nullable.Type[string] `json:"revision,omitempty"`

	// Specifies the URI used to identify the namespace within the ADMX file.
	TargetNamespace nullable.Type[string] `json:"targetNamespace,omitempty"`

	// Specifies the logical name that refers to the namespace within the ADMX file.
	TargetPrefix nullable.Type[string] `json:"targetPrefix,omitempty"`

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

func (s BaseGroupPolicyDefinitionFileImpl) GroupPolicyDefinitionFile() BaseGroupPolicyDefinitionFileImpl {
	return s
}

func (s BaseGroupPolicyDefinitionFileImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ GroupPolicyDefinitionFile = RawGroupPolicyDefinitionFileImpl{}

// RawGroupPolicyDefinitionFileImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawGroupPolicyDefinitionFileImpl struct {
	groupPolicyDefinitionFile BaseGroupPolicyDefinitionFileImpl
	Type                      string
	Values                    map[string]interface{}
}

func (s RawGroupPolicyDefinitionFileImpl) GroupPolicyDefinitionFile() BaseGroupPolicyDefinitionFileImpl {
	return s.groupPolicyDefinitionFile
}

func (s RawGroupPolicyDefinitionFileImpl) Entity() BaseEntityImpl {
	return s.groupPolicyDefinitionFile.Entity()
}

var _ json.Marshaler = BaseGroupPolicyDefinitionFileImpl{}

func (s BaseGroupPolicyDefinitionFileImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseGroupPolicyDefinitionFileImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseGroupPolicyDefinitionFileImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseGroupPolicyDefinitionFileImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.groupPolicyDefinitionFile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseGroupPolicyDefinitionFileImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalGroupPolicyDefinitionFileImplementation(input []byte) (GroupPolicyDefinitionFile, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling GroupPolicyDefinitionFile into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyUploadedDefinitionFile") {
		var out GroupPolicyUploadedDefinitionFile
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyUploadedDefinitionFile: %+v", err)
		}
		return out, nil
	}

	var parent BaseGroupPolicyDefinitionFileImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseGroupPolicyDefinitionFileImpl: %+v", err)
	}

	return RawGroupPolicyDefinitionFileImpl{
		groupPolicyDefinitionFile: parent,
		Type:                      value,
		Values:                    temp,
	}, nil

}
