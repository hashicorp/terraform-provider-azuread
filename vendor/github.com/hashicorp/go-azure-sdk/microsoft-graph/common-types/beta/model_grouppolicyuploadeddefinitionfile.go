package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ GroupPolicyDefinitionFile = GroupPolicyUploadedDefinitionFile{}

type GroupPolicyUploadedDefinitionFile struct {
	// The contents of the uploaded ADMX file.
	Content nullable.Type[string] `json:"content,omitempty"`

	// The default language of the uploaded ADMX file.
	DefaultLanguageCode nullable.Type[string] `json:"defaultLanguageCode,omitempty"`

	// The list of operations on the uploaded ADMX file.
	GroupPolicyOperations *[]GroupPolicyOperation `json:"groupPolicyOperations,omitempty"`

	// The list of ADML files associated with the uploaded ADMX file.
	GroupPolicyUploadedLanguageFiles *[]GroupPolicyUploadedLanguageFile `json:"groupPolicyUploadedLanguageFiles,omitempty"`

	// Type of Group Policy uploaded definition file status.
	Status *GroupPolicyUploadedDefinitionFileStatus `json:"status,omitempty"`

	// The uploaded time of the uploaded ADMX file.
	UploadDateTime *string `json:"uploadDateTime,omitempty"`

	// Fields inherited from GroupPolicyDefinitionFile

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

func (s GroupPolicyUploadedDefinitionFile) GroupPolicyDefinitionFile() BaseGroupPolicyDefinitionFileImpl {
	return BaseGroupPolicyDefinitionFileImpl{
		Definitions:          s.Definitions,
		Description:          s.Description,
		DisplayName:          s.DisplayName,
		FileName:             s.FileName,
		LanguageCodes:        s.LanguageCodes,
		LastModifiedDateTime: s.LastModifiedDateTime,
		PolicyType:           s.PolicyType,
		Revision:             s.Revision,
		TargetNamespace:      s.TargetNamespace,
		TargetPrefix:         s.TargetPrefix,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s GroupPolicyUploadedDefinitionFile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = GroupPolicyUploadedDefinitionFile{}

func (s GroupPolicyUploadedDefinitionFile) MarshalJSON() ([]byte, error) {
	type wrapper GroupPolicyUploadedDefinitionFile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GroupPolicyUploadedDefinitionFile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GroupPolicyUploadedDefinitionFile: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.groupPolicyUploadedDefinitionFile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GroupPolicyUploadedDefinitionFile: %+v", err)
	}

	return encoded, nil
}
